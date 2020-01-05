package lib

import (
	"strconv"

	"github.com/qulia/go-qulia/lib/hash"
)

// Metadata to append properties,tags to Graph, Node, Edge, etc
type Metadata map[string]interface{}

// OrderFunc definition used to decide heap configuration;
// function takes two elements and returns positive value if first > second,
// negative value if first < second, 0 otherwise
type OrderFunc func(first, second interface{}) int

type KeyFunc func(interface{}) string

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

	IntKeyFunc = func(elem interface{}) string {
		return strconv.Itoa(elem.(int))
	}

	HashKeyFunc = func(elem interface{}) string {
		return hash.Sha1(elem)
	}
)
