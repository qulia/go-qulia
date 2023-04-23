package leakybucket

import (
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
	"github.com/qulia/go-qulia/ratelimiter"
	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
)

func NewLeakyBucket[T any](capacity int, leakAmount int, leakPeriod time.Duration) ratelimiter.RateLimiterBuffered {
	q := queue.NewQueue[chan<- interface{}]()
	qAccessor := access.NewUnique(q)
	return &leakBucket{
		capacity:    capacity,
		qAccessor:   qAccessor,
		leakPeriod:  leakPeriod,
		tokenBucket: tokenbucket.NewTokenBucket(leakAmount, leakAmount, leakPeriod),
	}
}

type leakBucket struct {
	capacity    int
	leakPeriod  time.Duration
	qAccessor   *access.Unique[queue.Queue[chan<- interface{}]]
	tokenBucket ratelimiter.RateLimiter
}

func (lb *leakBucket) Allow() (<-chan interface{}, bool) {
	q, ok := lb.qAccessor.Acquire()
	if !ok {
		return nil, false
	}

	defer lb.qAccessor.Release()
	go lb.drain()

	if q.Length() == int(lb.capacity) {
		return nil, false
	}

	ch := make(chan interface{})
	q.Enqueue(chan<- interface{}(ch))
	return (<-chan interface{})(ch), true
}

func (lb *leakBucket) drain() {
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

// Close implements TokenBucket
func (lb *leakBucket) Close() {
	lb.tokenBucket.Close()
	lb.qAccessor.Close()
}
