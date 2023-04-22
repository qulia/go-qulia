package fixedwindowcounter

import (
	"time"

	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
)

type fixedWindowCounterImpl struct {
	tb tokenbucket.TokenBucket
}

// Close implements FixedWindowCounter
func (fwc *fixedWindowCounterImpl) Close() {
	fwc.tb.Close()
}

// Put implements FixedWindowCounter
func (fwc *fixedWindowCounterImpl) Put() bool {
	return fwc.tb.Take()
}

// TokenBucket with the provided window is the window "divider"
// Allowed threshold defined number of tokens per slot
// Windows not necessarily start at :00
func newFixedWindowCounterImpl(threshold int, window time.Duration) *fixedWindowCounterImpl {
	return &fixedWindowCounterImpl{
		tb: tokenbucket.NewTokenBucket(threshold, threshold, window),
	}
}
