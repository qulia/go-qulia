package window

import (
	"time"

	"github.com/qulia/go-qulia/v2/lib/common"
)

type Window[T common.Timed, R any, A common.Aggregator[T, R]] interface {
	Add(T) bool
	Results() <-chan PaneOutput[R]
}

type PaneOutput[R any] struct {
	Start time.Time
	End   time.Time
	Value R
}

func NewFixedWindow[T common.Timed, R any, A common.Aggregator[T, R]](
	size time.Duration, agg A, mtp common.TimeProvider,
) Window[T, R, A] {
	return newSlidingWindowImpl[T, R, A](size, size, agg, mtp)
}

func NewSlidingWindow[T common.Timed, R any, A common.Aggregator[T, R]](
	size time.Duration, sliding time.Duration, agg A, mtp common.TimeProvider,
) Window[T, R, A] {
	return newSlidingWindowImpl[T, R, A](size, sliding, agg, mtp)
}
