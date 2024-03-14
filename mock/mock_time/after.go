package mock_time

import "time"

type mockAfter struct {
	ch       chan time.Time
	deadline time.Time
}

func (ma *mockAfter) IsActive() bool {
	return true
}

func (ma *mockAfter) SendC() chan<- time.Time {
	return ma.ch
}

func (ma *mockAfter) Next() (time.Time, bool) {
	return time.Time{}, false
}

func (ma *mockAfter) Cancel() {
	// no-op
}
