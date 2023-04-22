package tokenbucket

import "time"

type TokenBucket interface {
	Take() bool
	Close()
}

func NewTokenBucket(capacity int, fillAmount int, fillPeriod time.Duration) TokenBucket {
	return newTokenBucketImpl(capacity, fillAmount, fillPeriod)
}
