package fixedwindowcounter

import (
	"time"

	"github.com/qulia/go-qulia/ratelimiter"
	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
)

func NewFixedWindowCounter(threshold int, window time.Duration) ratelimiter.RateLimiter {
	return &fixedWindowCounterImpl{
		tb: tokenbucket.NewTokenBucket(threshold, threshold, window),
	}
}

// TokenBucket with the provided window is the window "divider"
// Allowed threshold defined number of tokens per slot
// Windows not necessarily start at :00
type fixedWindowCounterImpl struct {
	tb ratelimiter.RateLimiter
}

// Close implements FixedWindowCounter
func (fwc *fixedWindowCounterImpl) Close() {
	fwc.tb.Close()
}

// Put implements FixedWindowCounter
func (fwc *fixedWindowCounterImpl) Allow() bool {
	return fwc.tb.Allow()
}
