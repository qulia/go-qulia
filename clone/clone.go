package clone

import "encoding/gob"

// Clones input
type Cloner[T any] interface {
	Clone(T) (T, error)
}

// Uses gob encoder/decoder to clone
func NewCloner[T any]() Cloner[T] {
	return newGobImpl[T]()
}

func NewInterfaceCloner[I any, C any]() Cloner[I] {
	gob.Register(*new(C))
	return newGobImpl[I]()
}
