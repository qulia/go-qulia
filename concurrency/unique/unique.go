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
	o, more := <-u.ch
	return o, more
}

func (u *Unique[T]) Release() {
	u.ch <- u.obj
}

func (u *Unique[T]) Close() {
	u.Acquire()
	close(u.ch)
}
