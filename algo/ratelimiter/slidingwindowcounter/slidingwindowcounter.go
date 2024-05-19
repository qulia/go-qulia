package slidingwindowcounter

import (
	"math"
	"time"

	"github.com/qulia/go-qulia/v2/algo/ratelimiter"
	"github.com/qulia/go-qulia/v2/concurrency/unique"
	"github.com/qulia/go-qulia/v2/lib/common"
	"github.com/qulia/go-qulia/v2/lib/queue"
)

// Window duration is determined by the lookback
// Allowed values for window 1hour, 1min, 1sec
func NewSlidingWindowCounter(threshold int, window time.Duration, timeP common.TimeProvider) ratelimiter.RateLimiter {
	if window != time.Second && window != time.Minute && window != time.Hour {
		panic("window not allowed")
	}

	return &slidingWindowCounter{
		threshold: threshold,
		window:    window,
		wm:        make(map[int]int),
		qAccessor: unique.NewUnique(queue.NewQueue[time.Time]()),
		timeP:     timeP,
	}
}

type slidingWindowCounter struct {
	threshold int
	window    time.Duration
	wm        map[int]int
	timeP     common.TimeProvider

	qAccessor *unique.Unique[queue.Queue[time.Time]]
}

func (swc *slidingWindowCounter) Close() {
	swc.qAccessor.Close()
}

func (swc *slidingWindowCounter) Allow() bool {
	q, ok := swc.qAccessor.Acquire()
	if !ok {
		return false
	}

	defer swc.qAccessor.Release()
	t := swc.timeP.Now()
	cleanup(q, t, swc.window, swc.wm)

	// based on current time and window
	// find previous slot, current slot and ratio previous slot in the window
	// calculate the total within the window
	// allow if less than threshold
	ps, cs, pratio := getPosition(swc.window, t)
	pCount := swc.wm[ps]
	cCount := swc.wm[cs]
	calculated := float64(cCount) + pratio*float64(pCount)
	if calculated < float64(swc.threshold) {
		wm := map[int]int{}
		wm[ps] = pCount
		wm[cs] = cCount + 1
		swc.wm = wm
		q.Enqueue(t)
		return true
	}

	return false
}

func getPosition(window time.Duration, t time.Time) (int, int, float64) {
	cs := 0
	ps := 0
	cratio := 0.0
	switch window {
	case time.Hour:
		cs = t.Hour()
		ps = (cs + 23) % 24
		cratio = float64(t.Minute()) / 60.0
	case time.Minute:
		cs = t.Minute()
		ps = (cs + 59) % 69
		cratio = float64(t.Second()) / 60.0
	case time.Second:
		cs = t.Second()
		ps = (cs + 59) % 69
		cratio = float64(t.Nanosecond()) / math.Pow10(9)
	}
	return ps, cs, (1 - cratio)
}

func cleanup(q queue.Queue[time.Time], timeNow time.Time, window time.Duration, wm map[int]int) {
	for q.Length() > 0 {
		if timeNow.Sub(q.Peek()) > window {
			// old entry, remove
			t := q.Dequeue()
			_, cs, _ := getPosition(window, t)
			wm[cs]--
		} else {
			break
		}
	}
}
