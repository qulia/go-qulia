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
func Visit(node *Node, call func(interface{})) {
	if node == nil {
		call(nil)
		return
	}

	Visit(node.Left, call)
	call(node.Data)
	Visit(node.Right, call)
}
