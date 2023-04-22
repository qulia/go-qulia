package slidingwindowlog

import (
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
)

type slidingWindowLogImpl struct {
	threshold int
	lookback  time.Duration
	qAccessor *access.Unique[queue.Queue[time.Time]]
}

// Close implements SlidingWindowLog
func (slw *slidingWindowLogImpl) Close() {
	slw.qAccessor.Close()
}

// Put implements SlidingWindowLog
func (swl *slidingWindowLogImpl) Put() bool {
	q, ok := swl.qAccessor.Acquire()
	if !ok {
		return false
	}

	defer swl.qAccessor.Release()
	t := time.Now()
	cleanup(q, t, swl.lookback)
	if q.Length() == swl.threshold {
		return false
	}

	q.Enqueue(t)
	return true
}

func cleanup(q queue.Queue[time.Time], timeNow time.Time, lookback time.Duration) {
	for q.Length() > 0 {
		if timeNow.Sub(q.Peek()) > lookback {
			// old entry, remove
			q.Dequeue()
		} else {
			break
		}
	}
}

func newSlidingWindowLogImpl(threshold int, lookback time.Duration) *slidingWindowLogImpl {
	return &slidingWindowLogImpl{
		threshold: threshold,
		lookback:  lookback,
		qAccessor: access.NewUnique(queue.NewQueue[time.Time]()),
	}
}
