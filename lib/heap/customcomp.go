package heap

import (
	"sort"

	"github.com/qulia/go-qulia/lib"
)

// Heap that allows custom comparison for the entries while maintaining heap properties
// The contained type T should implment <lib.Leed
type customComp[T lib.Lesser[T]] struct {
	maxOnTop bool
	buffer   []T
}

func newCustomComp[T lib.Lesser[T]](input []T, maxOnTop bool) *customComp[T] {
	buffer := make([]T, len(input))
	copy(buffer, input)
	return initHeap(buffer, maxOnTop)
}

func initHeap[T lib.Lesser[T]](buffer []T, maxOnTop bool) *customComp[T] {
	h := customComp[T]{buffer: buffer, maxOnTop: maxOnTop}
	h.heapify()
	return &h
}

func (h *customComp[T]) Insert(elem T) {
	// Insert at the end, sift up
	h.buffer = append(h.buffer, elem)
	h.siftUp(h.Size() - 1)
}

func (h *customComp[T]) Extract() T {
	// Capture first, swap with last, shrink 1, sift down from top
	first := h.buffer[0]
	h.swap(0, h.Size()-1)
	h.buffer = h.buffer[:h.Size()-1]
	h.siftDown(0)

	return first
}

func (h customComp[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h customComp[T]) Size() int {
	return len(h.buffer)
}

func (h *customComp[T]) siftUp(index int) {
	// If we are already at the root, nothing to do
	if index == 0 {
		return
	}

	current := index
	parent := (current - 1) / 2

	if top := h.getTop([]int{current, parent}); top != parent {
		h.swap(top, parent)
		h.siftUp(parent)
	}
}

func (h *customComp[T]) siftDown(index int) {
	// If at the leaf, done
	if index >= h.Size()/2 {
		return
	}

	parent := index
	left := 2*index + 1
	right := 2*index + 2

	comps := []int{parent, left}
	if right < h.Size() {
		comps = append(comps, right)
	}

	if top := h.getTop(comps); top != parent {
		h.swap(top, parent)
		h.siftDown(top)
	}
}

func (h *customComp[T]) swap(i, j int) {
	h.buffer[i], h.buffer[j] = h.buffer[j], h.buffer[i]
}

func (h *customComp[T]) heapify() {
	if h.Size() <= 1 {
		return
	}

	// leaf nodes are already heaps
	// Start at first non-leaf node and go up to the root sifting up as needed
	// Leaf nodes start at n/2 goes to n-1
	for i := h.Size()/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *customComp[T]) getTop(indices []int) int {
	sort.Slice(indices, func(i, j int) bool {
		return h.less(h.buffer[indices[i]], h.buffer[indices[j]])
	})
	return indices[len(indices)-1]
}

func (h *customComp[T]) less(one, other T) bool {
	if h.maxOnTop {
		return one.Less(other)
	} else {
		// reverse comp
		return other.Less(one)
	}
}
