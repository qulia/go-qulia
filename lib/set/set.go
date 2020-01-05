package set

import (
	"github.com/qulia/go-qulia/lib"
	log "github.com/sirupsen/logrus"
)

type Interface interface {
	// Add element to the set
	Add(interface{})

	//Remove element from the set
	Remove(interface{})

	// CopyTo another set
	CopyTo(other Interface)

	// Returns the union set with other set
	Union(other Interface) Interface

	// Returns the intersection set with other set
	Intersection(other Interface) Interface

	// Is current set subset (contained in) of the provided set
	IsSubsetOf(other Interface) bool

	// Is current set superset of the provided set
	IsSupersetOf(other Interface) bool

	// Returns true if set contains the element, -1, false otherwise
	Contains(interface{}) bool

	// Size of the set
	Size() int

	// Creates a slice from the set
	ToSlice() []interface{}

	// Initializes the set from slice
	FromSlice([]interface{})
}

// Set is implementation of set.Interface
// Comparisons to match elements are based on KeyFunc
type Set struct {
	entries map[string]interface{}
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

	var small *Set
	var large *Set
	if lenS < lenOther {
		small = s
		large = other.(*Set)
	} else {
		small = other.(*Set)
		large = s
	}

	intersectionSet := NewSet(s.keyFunc)
	for key, elem := range small.entries {
		if keyOther, ok := large.contains(elem); ok {
			if key != keyOther {
				log.Warnf("keys do not match for %v %s %s", elem, key, keyOther)
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

func (s *Set) Contains(elem interface{}) bool {
	_, ok := s.contains(elem)
	return ok
}

func (s *Set) contains(elem interface{}) (string, bool) {
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
		entries: make(map[string]interface{}),
		keyFunc: keyFunc,
	}

	return &set
}

func (s *Set) Size() int {
	return len(s.entries)
}

func (s *Set) Add(elem interface{}) {
	if key, ok := s.contains(elem); !ok {
		s.entries[key] = elem
	}
}

func (s *Set) Remove(elem interface{}) {
	if key, ok := s.contains(elem); ok {
		delete(s.entries, key)
	}
}

func (s *Set) CopyTo(other Interface) {
	for _, elem := range s.entries {
		other.Add(elem)
	}
}
