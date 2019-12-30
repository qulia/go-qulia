package tree

type Node struct {
	Left  *Node
	Right *Node
	Data  interface{}
}

func NewNode(data interface{}) *Node {
	node := Node{
		Data: data,
	}

	return &node
}

// in-order visit of node
func VisitInOrder(node *Node, call func(interface{})) {
	if node == nil {
		call(nil)
		return
	}

	VisitInOrder(node.Left, call)
	call(node.Data)
	VisitInOrder(node.Right, call)
}
