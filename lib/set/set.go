package set

import (
	"github.com/qulia/go-qulia/lib"
)

type Interface[T lib.Keyable[K], K comparable] interface {
	// Add element to the set
	Add(T)

	//Remove element from the set
	Remove(T)

	// CopyTo another set
	CopyTo(other Interface[T, K])

	// Returns the union set with other set
	Union(other Interface[T, K]) Interface[T, K]

	// Returns the intersection set with other set
	Intersection(other Interface[T, K]) Interface[T, K]

	// Is current set subset (contained in) of the provided set
	IsSubsetOf(other Interface[T, K]) bool

	// Is current set superset of the provided set
	IsSupersetOf(other Interface[T, K]) bool

	// Returns true if set contains the element, -1, false otherwise
	Contains(T) bool

	// Size of the set
	Len() int

	// Creates a slice from the set
	ToSlice() []T

	// Initializes the set from slice
	FromSlice([]T)

	// Create a slice of keys 
	Keys() []K
}

// Set is implementation of set.Interface
type KeyedSet[T lib.Keyable[K], K comparable] struct {
	entries map[K]T
}

type Set[T comparable] struct {
	KeyedSet[lib.DefaultKeyable[T], T]
}

func NewKeyedSet[T lib.Keyable[K], K comparable]() *KeyedSet[T, K] {
	set := KeyedSet[T,K]{
		entries: make(map[K]T),
	}

	return &set
}

func NewSet[T comparable]() *Set[T] {
	set := Set[T]{}
	set.entries = make(map[T]lib.DefaultKeyable[T])
	return &set
}

func (s *Set[T]) AddKey(elem T) {
	s.Add(lib.DefaultKeyable[T]{Val:elem})
} 

func (s *Set[T]) FromKeys(input []T) {
	for _, elem := range input {
		s.AddKey(elem)
	}
}

func (s *KeyedSet[T, K]) Len() int {
	return len(s.entries)
}

func (s *KeyedSet[T, K]) Add(elem T) {
	s.entries[elem.Key()] = elem
}

func (s *KeyedSet[T, K]) Remove(elem T) {
	s.RemoveKey(elem.Key())
}

func (s *KeyedSet[T, K]) RemoveKey(key K) {
	delete(s.entries, key)
}

func (s *KeyedSet[T, K]) Contains(elem T) bool {
	_, ok := s.entries[elem.Key()]
	return ok
}

func (s *KeyedSet[T, K]) ContainsKey(key K) bool {
	_, ok := s.entries[key]
	return ok
}

func (s *KeyedSet[T, K]) Union(other Interface[T, K]) Interface[T, K] {
	unionSet := NewKeyedSet[T,K]()
	s.CopyTo(unionSet)
	other.CopyTo(unionSet)

	return unionSet
}

func (s *KeyedSet[T, K]) Intersection(other Interface[T, K]) Interface[T, K] {
	intersectionSet := NewKeyedSet[T, K]()
	for _, elem := range s.entries {
		if other.Contains(elem) {
			intersectionSet.Add(elem)
		}
	}

	return intersectionSet
}

func (s *KeyedSet[T, K]) IsSubsetOf(other Interface[T, K]) bool {
	return s.Intersection(other).Len() == s.Len()
}

func (s *KeyedSet[T, K]) IsSupersetOf(other Interface[T, K]) bool {
	return other.IsSubsetOf(s)
}

func (s *KeyedSet[T, K]) Keys() []K {
	var res []K
	for k, _ := range s.entries {
		res = append(res, k)
	}

	return res
}

func (s *KeyedSet[T, K]) FromSlice(input []T) {
	for _, elem := range input {
		s.Add(elem)
	}
}

func (s *KeyedSet[T, K]) ToSlice() []T {
	var res []T
	for _, elem := range s.entries {
		res = append(res, elem)
	}

	return res
}

func (s *KeyedSet[T, K]) CopyTo(other Interface[T, K]) {
	for _, elem := range s.entries {
		other.Add(elem)
	}
}
