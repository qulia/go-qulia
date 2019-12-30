package set

import (
	"github.com/qulia/go-qulia/lib"
	log "github.com/sirupsen/logrus"
)

type Interface interface {
	Add(interface{})
	Remove(interface{})
	CopyTo(other Interface)
	Union(other Interface) Interface
	Intersection(other Interface) Interface
	IsSubsetOf(other Interface) bool
	IsSupersetOf(other Interface) bool
	Contains(interface{}) (int, bool)
	Size() int
	ToSlice() []interface{}
	FromSlice([]interface{})
}

type Set struct {
	entries map[int]interface{}
	keyFunc lib.KeyFunc
}

func (s *Set) FromSlice(input []interface{}) {
	for _, elem := range input {
		s.Add(elem)
	}
}

func (s *Set) Union(other Interface) Interface {
	unionSet := NewSet(s.keyFunc)
	s.CopyTo(unionSet)
	other.CopyTo(unionSet)

	return unionSet
}

func (s *Set) Intersection(other Interface) Interface {
	lenS := s.Size()
	lenOther := other.Size()

	var small Interface
	var large Interface
	if lenS < lenOther {
		small = s
		large = other
	} else {
		small = other
		large = s
	}

	intersectionSet := NewSet(s.keyFunc)
	for key, elem := range small.(*Set).entries {
		if keyOther, ok := large.Contains(elem); ok {
			if key != keyOther {
				log.Warnf("keys do not match for %v %d %d", elem, key, keyOther)
			}
			intersectionSet.Add(elem)
		}
	}

	return intersectionSet
}

func (s *Set) IsSubsetOf(other Interface) bool {
	return s.Intersection(other).Size() == s.Size()
}

func (s *Set) IsSupersetOf(other Interface) bool {
	return other.IsSubsetOf(s)
}

func (s *Set) Contains(elem interface{}) (int, bool) {
	key := s.keyFunc(elem)
	_, ok := s.entries[key]
	return key, ok
}

func (s *Set) ToSlice() []interface{} {
	var res []interface{}
	for _, elem := range s.entries {
		res = append(res, elem)
	}

	return res
}

func NewSet(keyFunc lib.KeyFunc) *Set {
	set := Set{
		entries: make(map[int]interface{}),
		keyFunc: keyFunc,
	}

	return &set
}

func (s *Set) Size() int {
	return len(s.entries)
}

func (s *Set) Add(elem interface{}) {
	if key, ok := s.Contains(elem); !ok {
		s.entries[key] = elem
	}
}

func (s *Set) Remove(elem interface{}) {
	if key, ok := s.Contains(elem); ok {
		delete(s.entries, key)
	}
}

func (s *Set) CopyTo(other Interface) {
	for _, elem := range s.entries {
		other.Add(elem)
	}
}
