package lib

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Metadata to append properties,tags to Graph, Node, Edge, etc
type Metadata map[string]interface{}

// OrderFunc definition used to decide heap configuration;
// function takes two elements and returns positive value if first > second,
// negative value if first < second, 0 otherwise
type OrderFunc func(first, second interface{}) int

type Lesser[T any] interface {
	Less(T) bool
}

type DefaultLesser[T constraints.Ordered] struct {
	Val T
}

func (dc DefaultLesser[T]) Less(other DefaultLesser[T]) bool {
	return dc.Val < other.Val
}

type Keyable[K comparable] interface {
	Key() K
}

type DefaultKeyable[T comparable] struct {
	Val T
}

func (dk DefaultKeyable[T]) Key() T {
	return dk.Val
}

var (
	IntCompFunc = func(first, second interface{}) int {
		firstInt := first.(int)
		secondInt := second.(int)
		if firstInt < secondInt {
			return -1
		} else if firstInt > secondInt {
			return 1
		} else {
			return 0
		}
	}
)

type QueryEvalFunc func(a, b interface{}) interface{}
type UpdateFunc func(current interface{}) interface{}
type DisjointValFunc func() interface{}

func IntQueryEvalMinFunc(a, b interface{}) interface{} {
	var aInt, bInt int
	if a == nil {
		aInt = math.MaxInt32
	} else {
		aInt = a.(int)
	}
	if b == nil {
		bInt = math.MaxInt32
	} else {
		bInt = b.(int)
	}
	if aInt < bInt {
		return aInt
	}
	return bInt
}

func IntQueryEvalSumFunc(a, b interface{}) interface{} {
	var aInt, bInt int
	if a == nil {
		aInt = 0
	} else {
		aInt = a.(int)
	}
	if b == nil {
		bInt = 0
	} else {
		bInt = b.(int)
	}
	return aInt + bInt
}
