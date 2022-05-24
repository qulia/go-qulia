package set

import "github.com/qulia/go-qulia/lib"

type keyableSet[T lib.Keyable[K], K comparable] struct {
	entries map[K]T
}

func newkeyableSet[T lib.Keyable[K], K comparable]() *keyableSet[T, K] {
	set := keyableSet[T, K]{
		entries: make(map[K]T),
	}

	return &set
}

func (s *keyableSet[T, K]) Len() int {
	return len(s.entries)
}

func (s *keyableSet[T, K]) Add(elem T) {
	s.entries[elem.Key()] = elem
}

func (s *keyableSet[T, K]) Remove(elem T) {
	delete(s.entries, elem.Key())
}

func (s *keyableSet[T, K]) Contains(elem T) bool {
	_, ok := s.entries[elem.Key()]
	return ok
}

func (s *keyableSet[T, K]) Union(other SetFlex[T, K]) SetFlex[T, K] {
	unionSet := newkeyableSet[T, K]()
	s.CopyTo(unionSet)
	other.CopyTo(unionSet)

	return unionSet
}

func (s *keyableSet[T, K]) Intersection(other SetFlex[T, K]) SetFlex[T, K] {
	intersectionSet := newkeyableSet[T, K]()
	for _, elem := range s.entries {
		if other.Contains(elem) {
			intersectionSet.Add(elem)
		}
	}

	return intersectionSet
}

func (s *keyableSet[T, K]) IsSubsetOf(other SetFlex[T, K]) bool {
	return s.Intersection(other).Len() == s.Len()
}

func (s *keyableSet[T, K]) IsSupersetOf(other SetFlex[T, K]) bool {
	return other.IsSubsetOf(s)
}

func (s *keyableSet[T, K]) Keys() []K {
	var res []K
	for k, _ := range s.entries {
		res = append(res, k)
	}

	return res
}

func (s *keyableSet[T, K]) FromSlice(input []T) SetFlex[T, K] {
	for _, elem := range input {
		s.Add(elem)
	}

	return s
}

func (s *keyableSet[T, K]) ToSlice() []T {
	var res []T
	for _, elem := range s.entries {
		res = append(res, elem)
	}

	return res
}

func (s *keyableSet[T, K]) CopyTo(other SetFlex[T, K]) {
	for _, elem := range s.entries {
		other.Add(elem)
	}
}
