package stack

type stackImpl[T any] struct {
	elements []T
}

func newStackImpl[T any]() *stackImpl[T] {
	return &stackImpl[T]{}
}

func (st *stackImpl[T]) Push(elem T) {
	st.elements = append(st.elements, elem)
}

func (st *stackImpl[T]) Pop() T {
	elem := st.elements[len(st.elements)-1]
	st.elements = st.elements[:len(st.elements)-1]
	return elem
}

func (st *stackImpl[T]) Peek() T {
	return st.elements[len(st.elements)-1]
}

func (st *stackImpl[T]) IsEmpty() bool {
	return st.Size() == 0
}

func (st *stackImpl[T]) Size() int {
	return len(st.elements)
}
