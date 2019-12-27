package stack

// Interface implemented by stack, LIFO
type  Interface interface {
	// Push element to the stack
	Push(x interface{})

	// Pop element from the top of the stack
	// If stack is empty, returns nil
	Pop() interface{}

	// Peek check the value of top element
	Peek() interface{}

	// IsEmpty returns true if stack is empty
	IsEmpty() bool

	// Size returns number of elements in the stack
	Size() int
}

// Implementation of stack.Interface
type Stack struct {
	elements []interface{}
}

func (st *Stack) Push(elem interface{}) {
	st.elements = append(st.elements, elem)
}

func (st *Stack) Pop() interface{} {
	if st.IsEmpty() {
		return nil
	}
	elem := st.elements[len(st.elements) - 1]
	st.elements = st.elements[:len(st.elements) - 1]
	return elem
}

func (st *Stack) Peek() interface{} {
	if st.IsEmpty() {
		return nil
	}
	return st.elements[len(st.elements) - 1]
}

func (st *Stack) IsEmpty() bool {
	return st.Size() == 0
}

func (st *Stack) Size() int {
	return len(st.elements)
}