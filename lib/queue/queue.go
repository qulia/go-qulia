package queue

import (
	"github.com/qulia/go-qulia/lib"
)

// Interface implemented by stack, LIFO
type Interface interface {
	// Enqueue element to the stack
	Enqueue(x interface{})

	// Dequeue element
	// If queue is empty, returns nil
	Dequeue() interface{}

	// IsEmpty returns true if queue is empty
	IsEmpty() bool

	// Length returns number of elements in the queue
	Length() int
}

// Implementation of queue.Interface
type Queue struct {
	elements []interface{}
	Metadata lib.Metadata
}

func NewQueue() *Queue {
	q := Queue{}
	q.Metadata = lib.Metadata{}

	return &q
}

func (q *Queue) Enqueue(elem interface{}) {
	q.elements = append(q.elements, elem)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	elem := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	return elem
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (st *Queue) Length() int {
	return len(st.elements)
}
