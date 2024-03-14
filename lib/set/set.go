package set

import (
	"github.com/qulia/go-qulia/lib/common"
)

// Set with elements of any comparable type
type Set[T comparable] interface {
	// Add element to the set
	Add(T)

	// Remove element from the set
	Remove(T)

	// CopyTo another set
	CopyTo(other Set[T])

	// Returns the union set with other set
	Union(other Set[T]) Set[T]

	// Returns the intersection set with other set
	Intersection(other Set[T]) Set[T]

	// Is current set subset (contained in) of the provided set
	IsSubsetOf(other Set[T]) bool

	// Is current set superset of the provided set
	IsSupersetOf(other Set[T]) bool

	// Returns true if set contains the element, -1, false otherwise
	Contains(T) bool

	// Size of the set
	Len() int

	// Creates a slice from the set
	ToSlice() []T

	// Initializes the set from slice
	FromSlice([]T) Set[T]

	// Create a slice of keys
	Keys() []T
}

// Set with elements of any type that implements common.Keyable with Key() method that returns
// K which is comparable
type SetFlex[T common.Keyable[K], K comparable] interface {
	// Add element to the set
	Add(T)

	// Remove element from the set
	Remove(T)

	// CopyTo another set
	CopyTo(other SetFlex[T, K])

	// Returns the union set with other set
	Union(other SetFlex[T, K]) SetFlex[T, K]

	// Returns the intersection set with other set
	Intersection(other SetFlex[T, K]) SetFlex[T, K]

	// Is current set subset (contained in) of the provided set
	IsSubsetOf(other SetFlex[T, K]) bool

	// Is current set superset of the provided set
	IsSupersetOf(other SetFlex[T, K]) bool

	// Returns true if set contains the element, -1, false otherwise
	Contains(T) bool

	// Gets the element with key, call Contains first
	GetWithKey(K) T

	// Size of the set
	Len() int

	// Creates a slice from the set
	ToSlice() []T

	// Initializes the set from slice
	FromSlice([]T) SetFlex[T, K]

	// Create a slice of keys
	Keys() []K
}

func NewSetFlex[T common.Keyable[K], K comparable]() SetFlex[T, K] {
	return newFlexImpl[T, K]()
}

func NewSet[T comparable]() Set[T] {
	return newSetImpl[T]()
}
