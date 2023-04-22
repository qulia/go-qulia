package slidingwindowlog

import "time"

type SlidingWindowLog interface {
	Put() bool
	Close()
}

func NewSlidingWindowLog(threshold int, lookback time.Duration) SlidingWindowLog {
	return newSlidingWindowLogImpl(threshold, lookback)
}
