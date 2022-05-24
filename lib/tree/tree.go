package tree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Left  *Node[T]
	Right *Node[T]
	Data  T
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	return &Node[T]{Data: data}
}

// in-order visit of node
func VisitInOrder[T constraints.Ordered](node *Node[T], call func(T)) {
	if node == nil {
		return
	}

	VisitInOrder(node.Left, call)
	call(node.Data)
	VisitInOrder(node.Right, call)
}
