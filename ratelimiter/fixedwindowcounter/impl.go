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

func newFixedWindowCounterImpl(threshold uint32, window time.Duration) *fixedWindowCounterImpl {
	return &fixedWindowCounterImpl{
		tb: tokenbucket.NewTokenBucket(threshold, threshold, window),
	}
}
