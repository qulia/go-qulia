package heap

import (
	"github.com/qulia/go-qulia/v2/lib/common"
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

	// Peek top element from the heap
	Peek() T

	// Size of the heap
	Size() int

	// IsEmpty returns true for empty heap, false otherwise
	IsEmpty() bool
}

// Heap content of any type should implement lib.Comparer interface
type HeapFlex[T common.Comparer[T]] interface {
	// Insert element to the heap
	Insert(T)

	// Extract top element from the heap
	// If heap is empty, the call will panic
	Extract() T

	// Peek top element from the heap
	Peek() T

	// Size of the heap
	Size() int

	// IsEmpty returns true for empty heap, false otherwise
	IsEmpty() bool
}

// NewMinHeapFlex initializes the heap structure from provided slice
// returned heap implements min heap properties where min value defined by
// lib.Comparer implementation of the type is at the top of the heap to be extracted first
//
// input: The input slice is cloned and will not be modified by this method
// Pass nil as input if you do not have any initial entries
func NewMinHeapFlex[T common.Comparer[T]](input []T) HeapFlex[T] {
	return newFlexImpl(input, false)
}

// NewMaxHeapFlex initializes the heap structure from provided slice.
// returned heap implements max heap properties where max value defined by
// lib.Comparer implementation of the type is at the top of the heap to be extracted first
//
// input: The input slice is cloned and will not be modified by this method.
// Pass nil as input if you do not have any initial entries
func NewMaxHeapFlex[T common.Comparer[T]](input []T) HeapFlex[T] {
	return newFlexImpl(input, true)
}

// NewMinHeap initializes the heap structure from provided slice.
// Returned heap implements heap properties using <, > operators
// of type with constraints.Ordered
//
// input: The input slice is cloned and will not be modified by this method.
// Pass nil as input if you do not have any initial entries
func NewMinHeap[T constraints.Ordered](input []T) Heap[T] {
	return newImpl(input, false)
}

// NewMaxHeap initializes the heap structure from provided slice.
// Returned heap implements heap properties using <, > operators
// of type with constraints.Ordered
//
// input: The input slice is cloned and will not be modified by this method.
// Pass nil as input if you do not have any initial entries
func NewMaxHeap[T constraints.Ordered](input []T) Heap[T] {
	return newImpl(input, true)
}
