package common

type Keyable[K comparable] interface {
	Key() K
}

type DefaultKeyable[T comparable] struct {
	Val T
}

func (dk DefaultKeyable[T]) Key() T {
	return dk.Val
}
