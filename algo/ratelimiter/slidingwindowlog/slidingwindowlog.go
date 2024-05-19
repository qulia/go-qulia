package slidingwindowlog

import (
	"time"

	"github.com/qulia/go-qulia/v2/algo/ratelimiter"
	"github.com/qulia/go-qulia/v2/concurrency/unique"
	"github.com/qulia/go-qulia/v2/lib/common"
	"github.com/qulia/go-qulia/v2/lib/queue"
)

func NewSlidingWindowLog(threshold int, lookback time.Duration, timeP common.TimeProvider) ratelimiter.RateLimiter {
	return &slidingWindowLog{
		threshold: threshold,
		lookback:  lookback,
		qAccessor: unique.NewUnique(queue.NewQueue[time.Time]()),
		timeP:     timeP,
	}
}

type slidingWindowLog struct {
	threshold int
	lookback  time.Duration
	qAccessor *unique.Unique[queue.Queue[time.Time]]
	timeP     common.TimeProvider
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
	t := swl.timeP.Now()
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
