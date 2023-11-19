package fixedwindowcounter

import (
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter"
	"github.com/qulia/go-qulia/algo/ratelimiter/tokenbucket"
)

// TokenBucket with the provided window is the "divider"
// Allowed threshold defines the number of tokens per slot
// Windows not necessarily start at exact intervals
func NewFixedWindowCounter(threshold int, window time.Duration) ratelimiter.RateLimiter {
	return &fixedWindowCounter{
		tb: tokenbucket.NewTokenBucket(threshold, threshold, window),
	}
}

type fixedWindowCounter struct {
	tb ratelimiter.RateLimiter
}

// Close implements FixedWindowCounter
func (fwc *fixedWindowCounter) Close() {
	fwc.tb.Close()
}

// Put implements FixedWindowCounter
func (fwc *fixedWindowCounter) Allow() bool {
	return fwc.tb.Allow()
}
