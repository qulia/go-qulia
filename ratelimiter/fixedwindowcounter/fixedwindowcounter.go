package fixedwindowcounter

import (
	"time"
)

type FixedWindowCounter interface {
	Put() bool
	Close()
}

func NewFixedWindowCounter(threshold uint32, window time.Duration) FixedWindowCounter {
	return newFixedWindowCounterImpl(threshold, window)
}
