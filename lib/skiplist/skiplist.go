package skiplist

import (
	"golang.org/x/exp/constraints"
)

type SkipList[T constraints.Ordered] interface {
	Add(T)
	// Removes and returns true if exists, no-op and returns false otherwise
	Remove(T) bool
	Search(T) bool
}

func NewSkipList[T constraints.Ordered](minVal, maxVal T) SkipList[T] {
	return newSkipListImpl(minVal, maxVal)
}
