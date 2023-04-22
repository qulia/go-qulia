package tokenbucket

import "time"

type TokenBucket interface {
	Take() bool
	Close()
}

func NewTokenBucket(capacity uint32, fillAmount uint32, fillPeriod time.Duration) TokenBucket {
	return newTokenBucketImpl(capacity, fillAmount, fillPeriod)
}
