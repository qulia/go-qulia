package utils

import (
	"github.com/qulia/go-qulia/lib/set"
)

// Returns true if all elements in slice2 are in slice1
func SliceContains[T comparable](slice1, slice2 []T) bool {
	set1 := set.NewSet[T]()
	set1.FromSlice(slice1)
	set2 := set.NewSet[T]()
	set2.FromSlice(slice2)

	return set1.IsSupersetOf(set2)
}

func SliceContainsElement[T comparable](slice []T, elem T) bool {
	return SliceContains(slice, []T{elem})
}
