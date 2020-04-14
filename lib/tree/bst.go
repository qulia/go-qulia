package tree

import (
	"github.com/qulia/go-qulia/lib"
	log "github.com/sirupsen/logrus"
)

// Binary Search Tree implementation https://en.wikipedia.org/wiki/Binary_search_tree
// Element comparison is based on OrderFunc provided to NewBST method
type BSTInterface interface {
	// Adds new elem to the tree, will not check if same val already exists
	Insert(interface{}) *Node

	// Returns the first node that matches
	Search(interface{}) *Node

	// In-order traversal of tree copied to slice
	ToSlice() []interface{}

	// In-order traversal of tree, calling the func for each element visited
	Traverse(func(interface{}))

	// Returns the largest value node that is smaller than or equal to the given value.
	Floor(interface{}) *Node

	// Returns the smallest value node that is larger than or equal to the given value.
	Ceiling(interface{}) *Node
}

type BST struct {
	root      *Node
	orderFunc lib.OrderFunc
}

func NewBST(orderFunc lib.OrderFunc) *BST {
	if orderFunc == nil {
		log.Fatal("Nil orderFunc param")
	}
	bst := BST{orderFunc: orderFunc}
	return &bst
}

func (bst *BST) Search(elem interface{}) *Node {
	return bst.search(bst.root, elem)
}

func (bst *BST) search(root *Node, elem interface{}) *Node {
	if root == nil {
		return nil
	}

	comp := bst.orderFunc(elem, root.Data)
	if comp == 0 {
		return root
	}

	if comp < 0 {
		return bst.search(root.Left, elem)
	}

	return bst.search(root.Right, elem)
}

func (bst *BST) ToSlice() []interface{} {
	var res []interface{}
	bst.Traverse(func(elem interface{}) {
		if elem == nil {
			return
		}
		res = append(res, elem)
	})

	return res
}

func (bst *BST) Traverse(call func(interface{})) {
	VisitInOrder(bst.root, call)
}

func (bst *BST) Insert(elem interface{}) *Node {
	return bst.insert(&bst.root, elem)
}

func (bst *BST) insert(root **Node, elem interface{}) *Node {
	if *root == nil {
		*root = NewNode(elem)
	} else if bst.orderFunc(elem, (*root).Data) < 0 {
		(*root).Left = bst.insert(&(*root).Left, elem)
	} else {
		(*root).Right = bst.insert(&(*root).Right, elem)
	}

	return *root
}

func (bst *BST) Floor(elem interface{}) *Node {
	return bst.floor(elem, bst.root)
}

func (bst *BST) floor(elem interface{}, root *Node) *Node {
	// if we reach to the leaf and value does not match, not found
	if root == nil {
		return nil
	}
	comp := bst.orderFunc(root.Data, elem)
	if comp == 0 {
		return root
	} else if comp > 0 {
		// root is larger, search left
		return bst.floor(elem, root.Left)
	} else {
		// root is smaller, but there could be something larger on right
		right := bst.floor(elem, root.Right)
		if right == nil {
			return root
		} else {
			return right
		}
	}
}

func (bst *BST) Ceiling(elem interface{}) *Node {
	return bst.ceiling(elem, bst.root)
}

func (bst *BST) ceiling(elem interface{}, root *Node) *Node {
	// if we reach to the leaf and value does not match, not found
	if root == nil {
		return nil
	}

	comp := bst.orderFunc(root.Data, elem)
	if comp == 0 {
		return root
	} else if comp < 0 {
		// root is smaller, search right
		return bst.ceiling(elem, root.Right)
	} else {
		// root is larger but there could be something smaller on the left
		left := bst.ceiling(elem, root.Left)
		if left == nil {
			return root
		} else {
			return left
		}
	}
}
