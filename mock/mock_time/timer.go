package mock_time

import (
	"sync"
	"time"

	"github.com/qulia/go-qulia/lib/common"
)

type mockTimer struct {
	deadline time.Time
	mtp      common.TimeProvider
	ch       chan time.Time
	duration time.Duration
	lock     sync.Mutex
	active   bool
}

func (mt *mockTimer) Stop() {
	mt.lock.Lock()
	defer mt.lock.Unlock()
	if mt.active {
		mt.active = false
		close(mt.ch)
	}
}

func (mt *mockTimer) IsActive() bool {
	mt.lock.Lock()
	defer mt.lock.Unlock()
	return mt.active
}

func (mt *mockTimer) C() <-chan time.Time {
	return mt.ch
}

func (mt *mockTimer) SendC() chan<- time.Time {
	return mt.ch
}

func (mt *mockTimer) Next() (time.Time, bool) {
	return time.Time{}, false
}

func (mt *mockTimer) Cancel() {
	mt.Stop()
}
