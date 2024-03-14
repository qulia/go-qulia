package clone

import (
	"bytes"
	"encoding/gob"
)

type goblImpl[T any] struct{}

func (gi *goblImpl[T]) Clone(in T) (T, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(&in)
	if err != nil {
		return *new(T), err
	}
	var out T
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&out)
	if err != nil {
		return *new(T), err
	}
	return out, nil
}

func newGobImpl[T any]() Cloner[T] {
	return &goblImpl[T]{}
}
