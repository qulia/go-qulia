package tree

import "golang.org/x/exp/constraints"

type bstImpl[T constraints.Ordered] struct {
	root *Node[T]
}

func newBSTImpl[T constraints.Ordered]() *bstImpl[T] {
	return &bstImpl[T]{}
}

func (bst *bstImpl[T]) Search(elem T) *Node[T] {
	return bst.search(bst.root, elem)
}

func (bst *bstImpl[T]) search(root *Node[T], elem T) *Node[T] {
	if root == nil {
		return nil
	}

	if elem == root.Data {
		return root
	}

	if elem < root.Data {
		return bst.search(root.Left, elem)
	}

	return bst.search(root.Right, elem)
}

func (bst *bstImpl[T]) ToSlice() []T {
	var res []T
	bst.Traverse(func(elem T) {
		res = append(res, elem)
	})

	return res
}

func (bst *bstImpl[T]) Traverse(call func(T)) {
	VisitInOrder(bst.root, call)
}

func (bst *bstImpl[T]) Insert(elem T) *Node[T] {
	return bst.insert(&bst.root, elem)
}

func (bst *bstImpl[T]) insert(root **Node[T], elem T) *Node[T] {
	if *root == nil {
		*root = NewNode(elem)
	} else if elem < (*root).Data {
		(*root).Left = bst.insert(&(*root).Left, elem)
	} else {
		(*root).Right = bst.insert(&(*root).Right, elem)
	}

	return *root
}

func (bst *bstImpl[T]) Floor(elem T) *Node[T] {
	return bst.floor(elem, bst.root)
}

func (bst *bstImpl[T]) floor(elem T, root *Node[T]) *Node[T] {
	// if we reach to the leaf and value does not match, not found
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	} else if root.Data > elem {
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

func (bst *bstImpl[T]) Ceiling(elem T) *Node[T] {
	return bst.ceiling(elem, bst.root)
}

func (bst *bstImpl[T]) ceiling(elem T, root *Node[T]) *Node[T] {
	// if we reach to the leaf and value does not match, not found
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	} else if root.Data < elem {
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
