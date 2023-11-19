package leakybucket

import (
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter"
	"github.com/qulia/go-qulia/algo/ratelimiter/tokenbucket"
	"github.com/qulia/go-qulia/concurrency/unique"
	"github.com/qulia/go-qulia/lib/queue"
)

// Allows flow in as long as not at capacity
// Whatever passes through is buffered, the caller is responsible for waiting on the channel to proceed
// It will be able to proceed based on leakAmount and leakPeriod
func NewLeakyBucket(capacity int, leakAmount int, leakPeriod time.Duration) ratelimiter.RateLimiterBuffered {
	q := queue.NewQueue[chan<- interface{}]()
	qAccessor := unique.NewUnique(q)
	return &leakyBucket{
		capacity:    capacity,
		qAccessor:   qAccessor,
		leakPeriod:  leakPeriod,
		tokenBucket: tokenbucket.NewTokenBucket(leakAmount, leakAmount, leakPeriod),
	}
}

type leakyBucket struct {
	capacity    int
	leakPeriod  time.Duration
	qAccessor   *unique.Unique[queue.Queue[chan<- interface{}]]
	tokenBucket ratelimiter.RateLimiter
}

func (lb *leakyBucket) Allow() (<-chan interface{}, bool) {
	q, ok := lb.qAccessor.Acquire()
	if !ok {
		return nil, false
	}

	defer lb.qAccessor.Release()
	go lb.drain()

	if q.Length() == lb.capacity {
		return nil, false
	}

	ch := make(chan interface{})
	q.Enqueue(chan<- interface{}(ch))
	return (<-chan interface{})(ch), true
}

func (lb *leakyBucket) drain() {
	q, ok := lb.qAccessor.Acquire()
	if !ok {
		return
	}

	defer lb.qAccessor.Release()
	for q.Length() > 0 && lb.tokenBucket.Allow() {
		x := q.Dequeue()
		go func(x chan<- interface{}) {
			x <- struct{}{}
		}(x)
	}
	// Schedule next check, in case there are no incoming calls
	time.AfterFunc(lb.leakPeriod, lb.drain)
}

func (lb *leakyBucket) Close() {
	lb.tokenBucket.Close()
	lb.qAccessor.Close()
}
