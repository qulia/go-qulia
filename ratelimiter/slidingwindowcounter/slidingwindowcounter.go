package slidingwindowcounter

import (
	"math"
	"time"

	access "github.com/qulia/go-qulia/concurrency/unique"
	"github.com/qulia/go-qulia/lib/queue"
	"github.com/qulia/go-qulia/ratelimiter"
)

// Window duration is determined by the lookback
// Allowed values for window 1hour, 1min, 1sec
func NewSlidingWindowCounter(threshold int, window time.Duration) ratelimiter.RateLimiter {
	if window != time.Second && window != time.Minute && window != time.Hour {
		panic("window not allowed")
	}

	return &slidingiWndowCounter{
		threshold: threshold,
		window:    window,
		wm:        make(map[int]int),
		qAccessor: access.NewUnique(queue.NewQueue[time.Time]()),
	}
}

type slidingiWndowCounter struct {
	threshold int
	window    time.Duration
	wm        map[int]int

	qAccessor *access.Unique[queue.Queue[time.Time]]
}

func (swc *slidingiWndowCounter) Close() {
	swc.qAccessor.Close()
}

// Calculates the estimated request count based on the previous window
// Removes older entries
func (swc *slidingiWndowCounter) Allow() bool {
	q, ok := swc.qAccessor.Acquire()
	if !ok {
		return false
	}

	defer swc.qAccessor.Release()
	t := time.Now()
	cleanup(q, t, swc.window, swc.wm)

	ps, cs, pratio := getPosition(swc.window, t)
	pCount := swc.wm[ps]
	cCount := swc.wm[cs]
	calculated := float64(cCount) + pratio*float64(pCount)
	if calculated <= float64(swc.threshold) {
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
