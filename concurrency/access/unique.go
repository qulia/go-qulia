package access

type Unique struct {
	obj interface{}
	ch  chan interface{}
}

func NewUnique(obj interface{}) *Unique {
	u := Unique{
		obj: obj,
		ch:  make(chan interface{}, 1),
	}

	return &u
}

func (u *Unique) Acquire() interface{} {
	return <-u.ch
}

func (u *Unique) Release() {
	u.ch <- u.obj
}

func (u *Unique) Done() {
	// Wait for the current user to be done before pushing nil
	u.Acquire()
	u.obj = nil
	u.Release()
}
