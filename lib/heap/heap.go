package heap

import (
	"github.com/qulia/go-qulia/lib"
	"golang.org/x/exp/constraints"
)

// Heap content type satisfy constraints.Ordered
// While comparing during heap oparation "<" is used
type Heap[T constraints.Ordered] interface {
	// Insert element to the heap
	Insert(T)

	// Extract top element from the heap
	// If heap is empty, the call will panic
	Extract() T

	// Size of the heap
	Size() int

	// IsEmpty returns true for empty heap, false otherwise
	IsEmpty() bool
}

// Heap content of any type should implement lib.Lesser interface
type HeapCustomComp[T lib.Lesser[T]] interface {
	// Insert element to the heap
	Insert(T)

	// Extract top element from the heap
	// If heap is empty, the call will panic
	Extract() T

	// Size of the heap
	Size() int

	// IsEmpty returns true for empty heap, false otherwise
	IsEmpty() bool
}

// NewMinHeapCustomComp initializes the heap structure from provided slice
// returned heap implements min heap properties where min value defined by
// lib.Lesser implementation of the type is at the top of the heap to be extracted first
//
// input: The input slice is cloned and will not be modified by this method
// Pass nil as input if you do not have any initial entries
func NewMinHeapCustomComp[T lib.Lesser[T]](input []T) HeapCustomComp[T] {
	return newCustomComp(input, false)
}

// NewMaxHeapCustomComp initializes the heap structure from provided slice
// returned heap implements max heap properties where max value defined by
// lib.Lesser implementation of the type is at the top of the heap to be extracted first
//
// input: The input slice is cloned and will not be modified by this method.
// Pass nil as input if you do not have any initial entries
func NewMaxHeapCustomComp[T lib.Lesser[T]](input []T) HeapCustomComp[T] {
	return newCustomComp(input, true)
}

// NewMinHeap initializes the heap structure from provided slice
// returned heap implements min heap properties where min value defined by
// < operator result of the type is at the top of the heap to be extracted first
//
// input: The input slice is cloned and will not be modified by this method.
// Pass nil as input if you do not have any initial entries
func NewMinHeap[T constraints.Ordered](input []T) Heap[T] {
	return newOrdered(input, false)
}

// NewMaxHeap initializes the heap structure from provided slice
// returned heap implements max heap properties where max value defined by
// < operator result of the type is at the top of the heap to be extracted first
//
// input: The input slice is cloned and will not be modified by this method.
// Pass nil as input if you do not have any initial entries
func NewMaxHeap[T constraints.Ordered](input []T) Heap[T] {
	return newOrdered(input, true)
}
