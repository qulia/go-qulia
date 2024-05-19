package window

import (
	"time"

	"github.com/qulia/go-qulia/v2/lib/common"
	"github.com/qulia/go-qulia/v2/lib/heap"
)

// Sliding window aggregator
// Assumes continuous invocation of Add call
// Out of window events are not added
//
// Window aggregation result is emitted to the Results channel
// If there are no Results receivers Add will block after an item
// Caller should synchronize calls if needed
type slidingWindowImpl[T common.Timed, R any, A common.Aggregator[T, R]] struct {
	size       time.Duration
	sliding    time.Duration
	agg        A
	wh         heap.HeapFlex[common.TimeComparer[T]]
	currentEnd time.Time
	results    chan PaneOutput[R]
	mtp        common.TimeProvider
}

func (swi *slidingWindowImpl[T, R, A]) Add(item T) bool {
	swi.processWindow()
	if swi.currentEnd.Sub(item.Time()) < 0 || swi.currentEnd.Sub(item.Time()) > swi.size {
		// out of window
		return false
	}

	swi.wh.Insert(common.TimeComparer[T]{Val: item})
	swi.agg.Add(item)

	return true
}

func (swi *slidingWindowImpl[T, R, A]) processWindow() {
	if swi.wh.Size() == 0 {
		return
	}

	if swi.mtp.Since(swi.currentEnd) > 0 {
		windowStart := swi.currentEnd.Add(-swi.size)
		// cleanup older entries
		for swi.wh.Size() > 0 && swi.wh.Peek().Val.Time().Before(windowStart) {
			val := swi.wh.Extract()
			swi.agg.Remove(val.Val)
		}

		// close window and emit result
		swi.results <- PaneOutput[R]{Start: swi.currentEnd.Add(-swi.size), End: swi.currentEnd, Value: swi.agg.Result()}
		swi.currentEnd = swi.currentEnd.Add(swi.sliding)
	}
}

func (swi *slidingWindowImpl[T, R, A]) Results() <-chan PaneOutput[R] {
	return swi.results
}

func newSlidingWindowImpl[T common.Timed, R any, A common.Aggregator[T, R]](
	size, sliding time.Duration, agg A, mtp common.TimeProvider,
) *slidingWindowImpl[T, R, A] {
	swi := slidingWindowImpl[T, R, A]{
		size:       size,
		sliding:    sliding,
		agg:        agg,
		wh:         heap.NewMinHeapFlex[common.TimeComparer[T]](nil),
		currentEnd: mtp.Now().Add(size),
		results:    make(chan PaneOutput[R], 1),
		mtp:        mtp,
	}

	return &swi
}
