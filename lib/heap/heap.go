package heap

import (
	log "github.com/sirupsen/logrus"
)

type Interface interface {
	// Insert element to the heap
	Insert(interface{})

	// Extract top element from the heap
	Extract()interface{}

	// Size of the heap
	Size() int

	// IsEmpty returns true for empty heap, false otherwise
	IsEmpty() bool
}

// CompareFunc definition used to decide heap configuration
type CompareFunc func(first, second interface{}) int

// MinHeap is heap structure with min element is top
type MinHeap struct {
}

// NewMinHeap initializes the heap structure from provided slice

// input: The input slice is cloned and will not be modified by this method
// Pass nil as input if you do not have any initial entries

// compareToFunc: function that takes two entries and returns positive value if first > second,
// negative value if first < second, 0 otherwise
func NewMinHeap(input []interface{}, compareToFunc CompareFunc) Interface {
	if compareToFunc == nil {
		log.Fatal("Nil compareToFunc param")
	}
	buffer := make([]interface{}, len(input))
	copy(buffer, input)
	return initHeap(buffer, compareToFunc, false)
}

// MinHeap is heap structure with max element is top
type MaxHeap struct {
}

// NewMaxHeap initializes the heap structure from provided slice
//
// input: The input slice is cloned and will not be modified by this method
// Pass nil as input if you do not have any initial entries
//
// compareToFunc: function that takes two entries and returns positive value if first > second,
// negative value if first < second, 0 otherwise
func NewMaxHeap(input []interface{}, compareToFunc CompareFunc) Interface{
	if compareToFunc == nil {
		log.Fatal("Nil compareToFunc param")
	}
	buffer := make([]interface{}, len(input))
	copy(buffer, input)
	return initHeap(buffer, compareToFunc, true)
}

type heap struct {
	maxOnTop bool
	buffer []interface{}
	compareFunc CompareFunc
}

func initHeap(buffer []interface{}, compareToFunc CompareFunc, maxOnTop bool) Interface {
	h := heap{buffer:buffer, compareFunc:compareToFunc, maxOnTop:maxOnTop}
	h.heapify()
	return &h
}

func (h *heap) Insert(elem interface{}) {
	// Insert at the end, sift up
	h.buffer = append(h.buffer, elem)
	h.siftUp(h.Size() - 1)
}

func (h *heap) Extract() interface{} {
	if h.IsEmpty() {
		return nil
	}

	// Capture first, swap with last, shrink 1, sift down from top
	first := h.buffer[0]
	h.swap(0, h.Size() - 1)
	h.buffer = h.buffer[:h.Size() - 1]
	h.siftDown(0)

	return first
}

func (h *heap) IsEmpty() bool {
	return h.Size() == 0
}

func (h * heap) Size() int {
	return len(h.buffer)
}

func (h * heap) siftUp(index int) {
	// If we are already at the root, nothing to do
	if index == 0 {
		return
	}

	current := index
	parent := current / 2

	top, equal := h.findTop(current, parent)
	if equal {
		return
	}

	if top != parent {
		h.swap(top, parent)
		h.siftUp(parent)
	}
}

func (h * heap) siftDown(index int) {
	// If at the leaf, done
	if index >= h.Size() / 2 {
		return
	}

	parent := index
	left := 2 * index + 1
	right := 2 * index + 2

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

	if top != parent {
		h.swap(top, parent)
		h.siftDown(top)
	}
}

func (h *heap) findTop(first int, second int) (int, bool) {
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

	comp := h.compareFunc(h.buffer[first], h.buffer[second]) * multiplier
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

func (h * heap) swap(i,j int) {
	tmp := h.buffer[i]
	h.buffer[i] = h.buffer[j]
	h.buffer[j] = tmp
}

func (h * heap) heapify() {
	if h.Size() <= 1 {
		return
	}

	// leaf nodes are already heaps
	// Start at first non-leaf node and go up to the root sifting up as needed
	// Leaf nodes start at n/2 goes to n-1
	for i := h.Size() / 2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}
