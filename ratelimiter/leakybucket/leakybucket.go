package leakybucket

import "time"

type LeakyBucket[T any] interface {
	Put(T) bool
	Take() (T, bool)
	Close()
}

func NewLeakyBucket[T any](capacity int, leakAmount int, leakPeriod time.Duration) LeakyBucket[T] {
	return newLeakyBucketImpl[T](capacity, leakAmount, leakPeriod)
}
