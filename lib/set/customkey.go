package set

import "github.com/qulia/go-qulia/lib"

type keyedSet[T lib.Keyable[K], K comparable] struct {
	entries map[K]T
}

func newKeyedSet[T lib.Keyable[K], K comparable]() *keyedSet[T, K] {
	set := keyedSet[T, K]{
		entries: make(map[K]T),
	}

	return &set
}

func (s *keyedSet[T, K]) Len() int {
	return len(s.entries)
}

func (s *keyedSet[T, K]) Add(elem T) {
	s.entries[elem.Key()] = elem
}

func (s *keyedSet[T, K]) Remove(elem T) {
	delete(s.entries, elem.Key())
}

func (s *keyedSet[T, K]) Contains(elem T) bool {
	_, ok := s.entries[elem.Key()]
	return ok
}

func (s *keyedSet[T, K]) Union(other CustomKeySet[T, K]) CustomKeySet[T, K] {
	unionSet := newKeyedSet[T, K]()
	s.CopyTo(unionSet)
	other.CopyTo(unionSet)

	return unionSet
}

func (s *keyedSet[T, K]) Intersection(other CustomKeySet[T, K]) CustomKeySet[T, K] {
	intersectionSet := newKeyedSet[T, K]()
	for _, elem := range s.entries {
		if other.Contains(elem) {
			intersectionSet.Add(elem)
		}
	}

	return intersectionSet
}

func (s *keyedSet[T, K]) IsSubsetOf(other CustomKeySet[T, K]) bool {
	return s.Intersection(other).Len() == s.Len()
}

func (s *keyedSet[T, K]) IsSupersetOf(other CustomKeySet[T, K]) bool {
	return other.IsSubsetOf(s)
}

func (s *keyedSet[T, K]) Keys() []K {
	var res []K
	for k, _ := range s.entries {
		res = append(res, k)
	}

	return res
}

func (s *keyedSet[T, K]) FromSlice(input []T) CustomKeySet[T, K] {
	for _, elem := range input {
		s.Add(elem)
	}

	return s
}

func (s *keyedSet[T, K]) ToSlice() []T {
	var res []T
	for _, elem := range s.entries {
		res = append(res, elem)
	}

	return res
}

func (s *keyedSet[T, K]) CopyTo(other CustomKeySet[T, K]) {
	for _, elem := range s.entries {
		other.Add(elem)
	}
}
