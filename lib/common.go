package lib

import (
	"golang.org/x/exp/constraints"
)

// Metadata to append properties,tags to Graph, Node, Edge, etc
type Metadata map[string]interface{}

// OrderFunc definition used to decide heap configuration;
// function takes two elements and returns positive value if first > second,
// negative value if first < second, 0 otherwise
type OrderFunc func(first, second interface{}) int

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
