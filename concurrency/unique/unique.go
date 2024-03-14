package unique

// Control access to obj with unique access
type Unique[T any] struct {
	obj T
	ch  chan T
}

func NewUnique[T any](obj T) *Unique[T] {
	u := Unique[T]{
		obj: obj,
		ch:  make(chan T, 1),
	}
	// initially make it available
	u.Release()
	return &u
}

func (u *Unique[T]) Acquire() (T, bool) {
	o, ok := <-u.ch
	return o, ok
}

func (u *Unique[T]) Release() {
	u.ch <- u.obj
}

// ok to call multiple times
func (u *Unique[T]) Close() {
	_, ok := u.Acquire()
	if ok {
		close(u.ch)
	}
}
