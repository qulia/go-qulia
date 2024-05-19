package set

import (
	"github.com/qulia/go-qulia/v2/lib/common"
)

type setImpl[T comparable] struct {
	*flexImpl[common.DefaultKeyable[T], T]
}

func newSetImpl[T comparable]() *setImpl[T] {
	return &setImpl[T]{newFlexImpl[common.DefaultKeyable[T], T]()}
}

func (s *setImpl[T]) Add(elem T) {
	s.flexImpl.Add(common.DefaultKeyable[T]{Val: elem})
}

func (s *setImpl[T]) Remove(elem T) {
	s.flexImpl.Remove(common.DefaultKeyable[T]{Val: elem})
}

func (s *setImpl[T]) Contains(elem T) bool {
	return s.flexImpl.Contains(common.DefaultKeyable[T]{Val: elem})
}

func (s *setImpl[T]) Union(other Set[T]) Set[T] {
	unionSet := newSetImpl[T]()
	s.CopyTo(unionSet)
	other.CopyTo(unionSet)

	return unionSet
}

func (s *setImpl[T]) Intersection(other Set[T]) Set[T] {
	intersectionSet := newSetImpl[T]()
	for _, elem := range s.flexImpl.entries {
		if other.Contains(elem.Key()) {
			intersectionSet.Add(elem.Key())
		}
	}

	return intersectionSet
}

func (s *setImpl[T]) IsSubsetOf(other Set[T]) bool {
	return s.Intersection(other).Len() == s.Len()
}

func (s *setImpl[T]) IsSupersetOf(other Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s *setImpl[T]) Keys() []T {
	return s.flexImpl.Keys()
}

func (s *setImpl[T]) FromSlice(input []T) Set[T] {
	for _, elem := range input {
		s.Add(elem)
	}

	return s
}

func (s *setImpl[T]) ToSlice() []T {
	var res []T
	for _, elem := range s.flexImpl.entries {
		res = append(res, elem.Key())
	}

	return res
}

func (s *setImpl[T]) CopyTo(other Set[T]) {
	for _, elem := range s.flexImpl.entries {
		other.Add(elem.Key())
	}
}
