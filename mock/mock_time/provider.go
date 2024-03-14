package mock_time

import (
	"sync"
	"time"

	"github.com/qulia/go-qulia/lib/common"
	"github.com/qulia/go-qulia/lib/heap"
)

// Maintains a heap of schedules
// Go routine checks and schedules one every period
type MockTimeProvider struct {
	currentTime time.Time
	schedules   heap.HeapFlex[schedule]
	checker     *time.Ticker
	period      time.Duration
	lock        sync.Mutex
}

func NewMockTimeProvider(start time.Time, schedulePeriod time.Duration) common.TimeProvider {
	mtp := &MockTimeProvider{
		currentTime: start,
		period:      schedulePeriod,
		schedules:   heap.NewMinHeapFlex[schedule](nil),
	}

	mtp.checker = time.NewTicker(schedulePeriod)
	go func() {
		for range mtp.checker.C {
			mtp.trigger()
		}
	}()
	return mtp
}

func (mtp *MockTimeProvider) Close() {
	mtp.lock.Lock()
	defer mtp.lock.Unlock()
	mtp.checker.Stop()
	cancelSchedules(mtp)
}

func (mtp *MockTimeProvider) Now() time.Time {
	mtp.lock.Lock()
	defer mtp.lock.Unlock()
	return mtp.currentTime
}

func (mtp *MockTimeProvider) Since(t time.Time) time.Duration {
	return mtp.Now().Sub(t)
}

func (mtp *MockTimeProvider) Sleep(d time.Duration) {
	mtp.lock.Lock()
	defer mtp.lock.Unlock()
	mtp.currentTime = mtp.currentTime.Add(d)
}

func (mtp *MockTimeProvider) NewTicker(d time.Duration) common.Ticker {
	mtp.lock.Lock()
	defer mtp.lock.Unlock()

	ticker := &mockTicker{
		ch:       make(chan time.Time, 1),
		duration: d,
		nextTick: mtp.currentTime.Add(d),
		active:   true,
	}
	mtp.schedules.Insert(schedule{scheduleAt: ticker.nextTick, sch: ticker})
	return ticker
}

func (mtp *MockTimeProvider) NewTimer(d time.Duration) common.Timer {
	mtp.lock.Lock()
	defer mtp.lock.Unlock()

	timer := &mockTimer{
		ch:       make(chan time.Time, 1),
		duration: d,
		deadline: mtp.currentTime.Add(d),
		mtp:      mtp,
		active:   true,
	}
	mtp.schedules.Insert(schedule{scheduleAt: timer.deadline, sch: timer})
	return timer
}

func (mtp *MockTimeProvider) After(d time.Duration) <-chan time.Time {
	mtp.lock.Lock()
	defer mtp.lock.Unlock()

	c := make(chan time.Time, 1)
	ma := &mockAfter{ch: c, deadline: mtp.currentTime.Add(d)}
	mtp.schedules.Insert(schedule{scheduleAt: ma.deadline, sch: ma})
	return c
}

func (mtp *MockTimeProvider) AfterFunc(d time.Duration, f func()) common.Timer {
	timer := mtp.NewTimer(d)
	go func() {
		<-timer.C()
		f()
	}()
	return timer
}

func (mtp *MockTimeProvider) trigger() {
	mtp.lock.Lock()
	defer mtp.lock.Unlock()
	if mtp.schedules.IsEmpty() {
		return
	}

	cur := mtp.schedules.Extract()
	// advance the time
	// Sleep can advance the time, check not to go backward
	if mtp.currentTime.Before(cur.scheduleAt) {
		mtp.currentTime = cur.scheduleAt
	}

	if !cur.sch.IsActive() {
		return
	}

	go func(tr schedule) {
		cur.sch.SendC() <- cur.scheduleAt
	}(cur)

	if nt, ok := cur.sch.Next(); ok {
		mtp.schedules.Insert(schedule{nt, cur.sch})
	}
}

func cancelSchedules(mtp *MockTimeProvider) {
	for !mtp.schedules.IsEmpty() {
		mtp.schedules.Extract().sch.Cancel()
	}
}
