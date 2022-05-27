package stack

type anyStack[T any] struct {
	elements []T
}

func (st *anyStack[T]) Push(elem T) {
	st.elements = append(st.elements, elem)
}

func (st *anyStack[T]) Pop() T {
	elem := st.elements[len(st.elements)-1]
	st.elements = st.elements[:len(st.elements)-1]
	return elem
}

func (st *anyStack[T]) Peek() T {
	return st.elements[len(st.elements)-1]
}

func (st *anyStack[T]) IsEmpty() bool {
	return st.Size() == 0
}

func (st *anyStack[T]) Size() int {
	return len(st.elements)
}
