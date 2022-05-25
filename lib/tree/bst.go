package tree

import (
	"golang.org/x/exp/constraints"
)

// Binary Search Tree https://en.wikipedia.org/wiki/Binary_search_tree
// Element comparison is based on OrderFunc provided to NewBST method
type BST[T constraints.Ordered] interface {
	// Adds new elem to the tree, will not check if same val already exists
	Insert(T) *Node[T]

	// Returns the first node that matches
	Search(T) *Node[T]

	// In-order traversal of tree copied to slice
	ToSlice() []T

	// In-order traversal of tree, calling the func for each element visited
	Traverse(func(T))

	// Returns the largest value node that is smaller than or equal to the given value.
	Floor(T) *Node[T]

	// Returns the smallest value node that is larger than or equal to the given value.
	Ceiling(T) *Node[T]
}

func NewBST[T constraints.Ordered]() *bstOrdered[T] {
	return &bstOrdered[T]{}
}
