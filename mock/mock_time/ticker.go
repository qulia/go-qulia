package mock_time

import (
	"sync"
	"time"
)

type mockTicker struct {
	nextTick time.Time
	ch       chan time.Time
	duration time.Duration
	lock     sync.Mutex
	active   bool
}

func (mt *mockTicker) Stop() {
	mt.lock.Lock()
	defer mt.lock.Unlock()
	if mt.active {
		mt.active = false
		close(mt.ch)
	}
}

func (mt *mockTicker) IsActive() bool {
	mt.lock.Lock()
	defer mt.lock.Unlock()
	return mt.active
}

func (mt *mockTicker) C() <-chan time.Time {
	return mt.ch
}

func (mt *mockTicker) SendC() chan<- time.Time {
	return mt.ch
}

func (mt *mockTicker) Next() (time.Time, bool) {
	mt.lock.Lock()
	defer mt.lock.Unlock()
	if mt.active {
		mt.nextTick = mt.nextTick.Add(mt.duration)
		return mt.nextTick, true
	}

	return time.Time{}, false
}

func (mt *mockTicker) Cancel() {
	mt.Stop()
}
