package lib

import (
	"golang.org/x/exp/constraints"
)

type Lesser[T any] interface {
	Less(T) bool
}

type DefaultLesser[T constraints.Ordered] struct {
	Val T
}

func (dc DefaultLesser[T]) Less(other DefaultLesser[T]) bool {
	return dc.Val < other.Val
}

type Keyable[K comparable] interface {
	Key() K
}

type DefaultKeyable[T comparable] struct {
	Val T
}

func (dk DefaultKeyable[T]) Key() T {
	return dk.Val
}
