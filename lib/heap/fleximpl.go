package heap

import (
	"github.com/qulia/go-qulia/lib/common"
)

// Heap that allows custom comparison for the entries while maintaining heap properties
// The contained type T should implment <lib.Leed
type fleximpl[T common.Comparer[T]] struct {
	maxOnTop bool
	buffer   []T
}

func newFlexImpl[T common.Comparer[T]](input []T, maxOnTop bool) *fleximpl[T] {
	buffer := make([]T, len(input))
	copy(buffer, input)
	return initHeap(buffer, maxOnTop)
}

func initHeap[T common.Comparer[T]](buffer []T, maxOnTop bool) *fleximpl[T] {
	h := fleximpl[T]{buffer: buffer, maxOnTop: maxOnTop}
	h.heapify()
	return &h
}

func (h *fleximpl[T]) Insert(elem T) {
	// Insert at the end, sift up
	h.buffer = append(h.buffer, elem)
	h.siftUp(h.Size() - 1)
}

func (h *fleximpl[T]) Peek() T {
	return h.buffer[0]
}

func (h *fleximpl[T]) Extract() T {
	// Capture first, swap with last, shrink 1, sift down from top
	first := h.buffer[0]
	h.swap(0, h.Size()-1)
	h.buffer = h.buffer[:h.Size()-1]
	h.siftDown(0)

	return first
}

func (h fleximpl[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h fleximpl[T]) Size() int {
	return len(h.buffer)
}

func (h *fleximpl[T]) siftUp(index int) {
	for {
		// If we are already at the root, nothing to do
		if index == 0 {
			return
		}

		current := index
		parent := (current - 1) / 2

		top, equal := h.findTop(current, parent)
		if equal || top == parent {
			return
		}

		h.swap(top, parent)
		index = parent
	}
}

func (h *fleximpl[T]) siftDown(index int) {
	for { // If at the leaf, done
		if index >= h.Size()/2 {
			return
		}

		parent := index
		left := 2*index + 1
		right := 2*index + 2

		top, equal := h.findTop(left, right)
		if equal {
			top, equal = h.findTop(parent, left)
			if equal {
				return
			}
		} else {
			top, equal = h.findTop(parent, top)
			if equal {
				return
			}
		}

		if top == parent {
			return
		}
		h.swap(top, parent)
		index = top
	}
}

func (h *fleximpl[T]) findTop(first int, second int) (int, bool) {
	if first >= h.Size() {
		return second, false
	}

	if second >= h.Size() {
		return first, false
	}

	var top int
	multiplier := 1
	if !h.maxOnTop {
		multiplier = -1
	}

	comp := h.buffer[first].Compare(h.buffer[second]) * multiplier
	if comp > 0 {
		top = first
	} else if comp < 0 {
		top = second
	} else {
		// Equal
		return -1, true
	}
	return top, false
}

func (h *fleximpl[T]) swap(i, j int) {
	tmp := h.buffer[i]
	h.buffer[i] = h.buffer[j]
	h.buffer[j] = tmp
}

func (h *fleximpl[T]) heapify() {
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
