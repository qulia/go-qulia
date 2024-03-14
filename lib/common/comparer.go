package common

import "golang.org/x/exp/constraints"

// Return negative if less, positive if greater, 0 if equal
type Comparer[T any] interface {
	Compare(T) int
}

type DefaultComparer[T constraints.Ordered] struct {
	Val T
}

func (dc DefaultComparer[T]) Compare(other DefaultComparer[T]) int {
	if dc.Val < other.Val {
		return -1
	} else if dc.Val > other.Val {
		return 1
	}

	return 0
}
