package leakybucket

import "time"

type LeakyBucket[T any] interface {
	Put(T) bool
	Take() (T, bool)
	Close()
}

func NewLeakyBucket[T any](capacity uint32, leakAmount uint32, leakPeriod time.Duration) LeakyBucket[T] {
	return newLeakyBucketImpl[T](capacity, leakAmount, leakPeriod)
}
