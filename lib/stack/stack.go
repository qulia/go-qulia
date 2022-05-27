package stack

// Interface implemented by stack, LIFO
type Stack[T any] interface {
	// Push element to the stack
	Push(T)

	// Pop element from the top of the stack
	// If stack is empty, panics
	Pop() T

	// Peek check the value of top element
	// If stack is empty, panics
	Peek() T

	// IsEmpty returns true if stack is empty, call before Pop and Peek
	IsEmpty() bool

	// Size returns number of elements in the stack
	Size() int
}

func NewStack[T any]() Stack[T] {
	return newStackImpl[T]()
}
