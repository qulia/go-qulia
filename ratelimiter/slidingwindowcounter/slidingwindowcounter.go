package slidingwindowcounter

import (
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
	"github.com/qulia/go-qulia/ratelimiter"
)

// Window duration is determined by the lookback
// Allowed values for window 1hour, 1min, 1sec
func NewSlidingWindowCounter(threshold int, window time.Duration) ratelimiter.RateLimiter {
	if window != time.Second && window != time.Minute && window != time.Hour {
		panic("window not allowed")
	}

	return &slidingIWndowCounter{
		threshold: threshold,
		window:    window,
		wm:        make(map[int]int),
		qAccessor: access.NewUnique(queue.NewQueue[time.Time]()),
	}
}

type slidingIWndowCounter struct {
	threshold int
	window    time.Duration
	wm        map[int]int

	qAccessor *access.Unique[queue.Queue[time.Time]]
}

func (swc *slidingIWndowCounter) Close() {
	swc.qAccessor.Close()
}

func (swc *slidingIWndowCounter) Allow() bool {
	q, ok := swc.qAccessor.Acquire()
	if !ok {
		return false
	}

	defer swc.qAccessor.Release()
	t := time.Now()
	cleanup(q, t, swc.window, swc.wm)

	currentSlot := getSlot(swc.window, t)

	previousSlot := currentSlot - 1
	countPreviousWindow := swc.wm[previousSlot]
	countCurrentWindow := swc.wm[currentSlot]
	ratioPrevious := 0.0
	if countCurrentWindow+countPreviousWindow > 0 {
		ratioPrevious = float64(countPreviousWindow) / float64(countCurrentWindow+countPreviousWindow)
	}
	calculated := float64(countCurrentWindow) + ratioPrevious*float64(countPreviousWindow)
	if calculated <= float64(swc.threshold) {
		wm := map[int]int{}
		wm[previousSlot] = countPreviousWindow
		wm[currentSlot] = countCurrentWindow + 1
		swc.wm = wm
		q.Enqueue(t)
		return true
	}

	return false
}

func getSlot(window time.Duration, t time.Time) int {
	currentSlot := 0
	switch window {
	case time.Hour:
		currentSlot = t.Hour()
	case time.Minute:
		currentSlot = t.Minute()
	case time.Second:
		currentSlot = t.Second()
	}
	return currentSlot
}

func cleanup(q queue.Queue[time.Time], timeNow time.Time, window time.Duration, wm map[int]int) {
	for q.Length() > 0 {
		if timeNow.Sub(q.Peek()) > window {
			// old entry, remove
			t := q.Dequeue()
			wm[getSlot(window, t)]--
		} else {
			break
		}
	}
}
