package set

import (
	"github.com/qulia/go-qulia/lib"
)

// Set with elements of any type that implements lib.Keyable with Key() method that returns
// K which is comparable
type CustomKeySet[T lib.Keyable[K], K comparable] interface {
	// Add element to the set
	Add(T)

	//Remove element from the set
	Remove(T)

	// CopyTo another set
	CopyTo(other CustomKeySet[T, K])

	// Returns the union set with other set
	Union(other CustomKeySet[T, K]) CustomKeySet[T, K]

	// Returns the intersection set with other set
	Intersection(other CustomKeySet[T, K]) CustomKeySet[T, K]

	// Is current set subset (contained in) of the provided set
	IsSubsetOf(other CustomKeySet[T, K]) bool

	// Is current set superset of the provided set
	IsSupersetOf(other CustomKeySet[T, K]) bool

	// Returns true if set contains the element, -1, false otherwise
	Contains(T) bool

	// Size of the set
	Len() int

	// Creates a slice from the set
	ToSlice() []T

	// Initializes the set from slice
	FromSlice([]T) CustomKeySet[T, K]

	// Create a slice of keys
	Keys() []K
}

// Set with elements of any comparable type
type Set[T comparable] interface {
	// Add element to the set
	Add(T)

	//Remove element from the set
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

func NewCustomKeySet[T lib.Keyable[K], K comparable]() CustomKeySet[T, K] {
	return newKeyedSet[T, K]()
}

func NewSet[T comparable]() Set[T] {
	return newComparableSet[T]()
}
