package queue

func newQueueImpl[T any]() *queueImpl[T] {
	return &queueImpl[T]{}
}

// Implementation of queueImpl interface
type queueImpl[T any] struct {
	elements []T
}

func (q *queueImpl[T]) Enqueue(elem T) {
	q.elements = append(q.elements, elem)
}

func (q *queueImpl[T]) Dequeue() T {
	elem := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	return elem
}

func (q *queueImpl[T]) IsEmpty() bool {
	return q.Length() == 0
}

func (q *queueImpl[T]) Length() int {
	return len(q.elements)
}
