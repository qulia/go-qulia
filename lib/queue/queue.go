package queue

// FIFO Queue
type Queue[T any] interface {
	// Enqueues element
	Enqueue(x T)

	// Dequeues element
	// If queue is empty, it panics
	Dequeue() T

	// Returns the value of the front of the queue
	Peek() T

	// IsEmpty returns true if queue is empty
	IsEmpty() bool

	// Length returns number of elements in the queue
	Length() int
}

func NewQueue[T any]() Queue[T] {
	return newQueueImpl[T]()
}
