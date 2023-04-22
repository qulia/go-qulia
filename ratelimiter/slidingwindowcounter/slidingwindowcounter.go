package slidingwindowcounter

import "time"

type SlidingWindowCounter interface {
	Put() bool
	Close()
}

// Window duration is determined by the lookback
// Allowed values for window 1hour, 1min, 1sec
func NewSlidingWindowCounter(threshold int, window time.Duration) SlidingWindowCounter {
	return newSlidingWindowCounterImpl(threshold, window)
}
