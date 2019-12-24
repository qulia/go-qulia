package stack

type  Interface interface {
	Push(x interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	Length() int
}

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
	return st.Length() == 0
}

func (st *Stack) Length() int {
	return len(st.elements)
}