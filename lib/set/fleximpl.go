package set

import "github.com/qulia/go-qulia/lib"

type flexImpl[T lib.Keyable[K], K comparable] struct {
	entries map[K]T
}

func newFlexImpl[T lib.Keyable[K], K comparable]() *flexImpl[T, K] {
	set := flexImpl[T, K]{
		entries: make(map[K]T),
	}

	return &set
}

func (s *flexImpl[T, K]) Len() int {
	return len(s.entries)
}

func (s *flexImpl[T, K]) Add(elem T) {
	s.entries[elem.Key()] = elem
}

func (s *flexImpl[T, K]) Remove(elem T) {
	delete(s.entries, elem.Key())
}

func (s *flexImpl[T, K]) Contains(elem T) bool {
	_, ok := s.entries[elem.Key()]
	return ok
}

func (s *flexImpl[T, K]) GetWithKey(key K) T {
	if el, ok := s.entries[key]; ok {
		return el
	}
	return *new(T)
}

func (s *flexImpl[T, K]) Union(other SetFlex[T, K]) SetFlex[T, K] {
	unionSet := newFlexImpl[T, K]()
	s.CopyTo(unionSet)
	other.CopyTo(unionSet)

	return unionSet
}

func (s *flexImpl[T, K]) Intersection(other SetFlex[T, K]) SetFlex[T, K] {
	intersectionSet := newFlexImpl[T, K]()
	for _, elem := range s.entries {
		if other.Contains(elem) {
			intersectionSet.Add(elem)
		}
	}

	return intersectionSet
}

func (s *flexImpl[T, K]) IsSubsetOf(other SetFlex[T, K]) bool {
	return s.Intersection(other).Len() == s.Len()
}

func (s *flexImpl[T, K]) IsSupersetOf(other SetFlex[T, K]) bool {
	return other.IsSubsetOf(s)
}

func (s *flexImpl[T, K]) Keys() []K {
	var res []K
	for k := range s.entries {
		res = append(res, k)
	}

	return res
}

func (s *flexImpl[T, K]) FromSlice(input []T) SetFlex[T, K] {
	for _, elem := range input {
		s.Add(elem)
	}

	return s
}

func (s *flexImpl[T, K]) ToSlice() []T {
	var res []T
	for _, elem := range s.entries {
		res = append(res, elem)
	}

	return res
}

func (s *flexImpl[T, K]) CopyTo(other SetFlex[T, K]) {
	for _, elem := range s.entries {
		other.Add(elem)
	}
}
