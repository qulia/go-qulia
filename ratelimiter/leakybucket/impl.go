package leakybucket

import (
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
)

type leakyBucketImpl[T any] struct {
	capacity   int
	qAccessor  *access.Unique[queue.Queue[T]]
	leakBucket tokenbucket.TokenBucket
}

func (lb *leakyBucketImpl[T]) Put(x T) bool {
	q, ok := lb.qAccessor.Acquire()
	if !ok {
		return false
	}

	defer lb.qAccessor.Release()
	if q.Length() == int(lb.capacity) {
		return false
	}

	q.Enqueue(x)
	return true
}

func (lb *leakyBucketImpl[T]) Take() (T, bool) {
	q, ok := lb.qAccessor.Acquire()
	if !ok {
		return *new(T), false
	}

	defer lb.qAccessor.Release()
	if q.Length() > 0 && lb.leakBucket.Take() {
		x := q.Dequeue()
		return x, true
	}

	return *new(T), false
}

func newLeakyBucketImpl[T any](capacity int, leakAmount int, leakPeriod time.Duration) *leakyBucketImpl[T] {
	q := queue.NewQueue[T]()
	qAccessor := access.NewUnique(q)
	return &leakyBucketImpl[T]{
		capacity:   capacity,
		qAccessor:  qAccessor,
		leakBucket: tokenbucket.NewTokenBucket(leakAmount, leakAmount, leakPeriod),
	}
}

// Close implements TokenBucket
func (lb *leakyBucketImpl[T]) Close() {
	lb.leakBucket.Close()
	lb.qAccessor.Close()
}
