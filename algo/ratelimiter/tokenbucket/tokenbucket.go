package tokenbucket

import (
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter"
	"github.com/qulia/go-qulia/concurrency/unique"
)

func NewTokenBucket(capacity int, fillAmount int, fillPeriod time.Duration) ratelimiter.RateLimiter {
	if fillAmount > capacity || fillPeriod == 0 {
		panic("invalid arguments")
	}
	tb := &tockenBucket{
		capacity:   capacity,
		fillAmount: fillAmount,
		tokens:     fillAmount,
		fillPeriod: fillPeriod,
		lastFill:   time.Now(),
	}

	tb.tokensAccessor = unique.NewUnique(&tb.tokens)
	return tb
}

type tockenBucket struct {
	capacity   int
	fillAmount int
	fillPeriod time.Duration
	lastFill   time.Time
	tokens     int

	tokensAccessor *unique.Unique[*int]
}

// Close implements TokenBucket
func (tb *tockenBucket) Close() {
	tb.tokensAccessor.Close()
}

// Take implements TokenBucket
func (tb *tockenBucket) Allow() bool {
	_, ok := tb.tokensAccessor.Acquire()
	if !ok {
		return false
	}

	defer tb.tokensAccessor.Release()

	tb.fill()
	if tb.tokens == 0 {
		return false
	}
	tb.tokens--
	return true
}

func (tb *tockenBucket) fill() {
	tn := time.Now()
	periods := tn.Sub(tb.lastFill) / tb.fillPeriod
	refillAmount := int(periods) * int(tb.fillAmount)

	if refillAmount > 0 && refillAmount+tb.tokens < tb.capacity {
		tb.tokens = refillAmount + tb.tokens
		tb.lastFill = tn
	} else if refillAmount < 0 /*overflow*/ || refillAmount+tb.tokens >= tb.capacity {
		tb.tokens = tb.capacity
		tb.lastFill = tn
	}
}
