package utils

import (
	"github.com/qulia/go-qulia/lib"
	"github.com/qulia/go-qulia/lib/set"
)

// Returns true if all elements in slice2 are in slice1
func SliceContains(slice1, slice2 []interface{}, keyFunc lib.KeyFunc) bool {
	set1 := set.NewSet(keyFunc)
	set1.FromSlice(slice1)
	set2 := set.NewSet(keyFunc)
	set2.FromSlice(slice2)

	return set1.IsSupersetOf(set2)
}

func SliceContainsElement(slice []interface{}, elem interface{}, keyFunc lib.KeyFunc) bool {
	return SliceContains(slice, []interface{}{elem}, keyFunc)
}
