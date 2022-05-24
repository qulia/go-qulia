package queue

// FIFO Queue
type Queue[T any] interface {
	// Enqueues element
	Enqueue(x T)

	// Dequeues element
	// If queue is empty, it panics
	Dequeue() T

	// IsEmpty returns true if queue is empty
	IsEmpty() bool

	// Length returns number of elements in the queue
	Length() int
}

// Implementation of queue interface
type queue[T any] struct {
	elements []T
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(elem T) {
	q.elements = append(q.elements, elem)
}

func (q *queue[T]) Dequeue() T {
	elem := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	return elem
}

func (q *queue[T]) IsEmpty() bool {
	return q.Length() == 0
}

func (q *queue[T]) Length() int {
	return len(q.elements)
}
