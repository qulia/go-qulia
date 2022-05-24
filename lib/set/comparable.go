package set

import "github.com/qulia/go-qulia/lib"

type comparableSet[T comparable] struct {
	ks *keyableSet[lib.DefaultKeyable[T], T]
}

func newComparableSet[T comparable]() *comparableSet[T] {
	return &comparableSet[T]{ks: newkeyableSet[lib.DefaultKeyable[T], T]()}
}

func (s *comparableSet[T]) Len() int {
	return s.ks.Len()
}

func (s *comparableSet[T]) Add(elem T) {
	s.ks.Add(lib.DefaultKeyable[T]{Val: elem})
}

func (s *comparableSet[T]) Remove(elem T) {
	s.ks.Remove(lib.DefaultKeyable[T]{Val: elem})
}

func (s *comparableSet[T]) Contains(elem T) bool {
	return s.ks.Contains(lib.DefaultKeyable[T]{Val: elem})
}

func (s *comparableSet[T]) Union(other Set[T]) Set[T] {
	unionSet := newComparableSet[T]()
	s.CopyTo(unionSet)
	other.CopyTo(unionSet)

	return unionSet
}

func (s *comparableSet[T]) Intersection(other Set[T]) Set[T] {
	intersectionSet := newComparableSet[T]()
	for _, elem := range s.ks.entries {
		if other.Contains(elem.Key()) {
			intersectionSet.Add(elem.Key())
		}
	}

	return intersectionSet
}

func (s *comparableSet[T]) IsSubsetOf(other Set[T]) bool {
	return s.Intersection(other).Len() == s.Len()
}

func (s *comparableSet[T]) IsSupersetOf(other Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s *comparableSet[T]) Keys() []T {
	return s.ks.Keys()
}

func (s *comparableSet[T]) FromSlice(input []T) Set[T] {
	for _, elem := range input {
		s.Add(elem)
	}

	return s
}

func (s *comparableSet[T]) ToSlice() []T {
	var res []T
	for _, elem := range s.ks.entries {
		res = append(res, elem.Key())
	}

	return res
}

func (s *comparableSet[T]) CopyTo(other Set[T]) {
	for _, elem := range s.ks.entries {
		other.Add(elem.Key())
	}
}
