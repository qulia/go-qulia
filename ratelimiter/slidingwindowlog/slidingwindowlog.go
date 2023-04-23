package slidingwindowlog

import (
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
	"github.com/qulia/go-qulia/ratelimiter"
)

func NewSlidingWindowLog(threshold int, lookback time.Duration) ratelimiter.RateLimiter {
	return &slidingWindowLog{
		threshold: threshold,
		lookback:  lookback,
		qAccessor: access.NewUnique(queue.NewQueue[time.Time]()),
	}
}

type slidingWindowLog struct {
	threshold int
	lookback  time.Duration
	qAccessor *access.Unique[queue.Queue[time.Time]]
}

func (slw *slidingWindowLog) Close() {
	slw.qAccessor.Close()
}

func (swl *slidingWindowLog) Allow() bool {
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
