package tokenbucket

import (
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
)

type tokenBucketImpl struct {
	capacity   uint32
	fillAmount uint32

	tokens uint32
	ticker *time.Ticker

	tokensAccessor *access.Unique[*uint32]
}

// Close implements TokenBucket
func (tb *tokenBucketImpl) Close() {
	tb.tokensAccessor.Close()
}

// Take implements TokenBucket
func (tb *tokenBucketImpl) Take() bool {
	cur, ok := tb.tokensAccessor.Acquire()
	defer tb.tokensAccessor.Release()
	if !ok || *cur == 0 {
		return false
	}

	*cur--
	return true
}

func (tb *tokenBucketImpl) filler() {
	for range tb.ticker.C {
		cur, ok := tb.tokensAccessor.Acquire()
		if !ok {
			return
		}

		if *cur+tb.fillAmount <= tb.capacity {
			*cur += tb.fillAmount
		} else {
			*cur = tb.capacity
		}

		tb.tokensAccessor.Release()
	}
}

func newTokenBucketImpl(capacity uint32, fillAmount uint32, fillPeriod time.Duration) *tokenBucketImpl {
	if fillAmount > capacity {
		panic("fillAmount is larger than the capacity")
	}
	tb := &tokenBucketImpl{
		capacity:   capacity,
		fillAmount: fillAmount,
		tokens:     0,
		ticker:     time.NewTicker(fillPeriod),
	}

	go tb.filler()
	tb.tokensAccessor = access.NewUnique(&tb.tokens)
	tb.tokensAccessor.Release()
	return tb
}
