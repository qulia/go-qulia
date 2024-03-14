package common

import "time"

type Timed interface {
	Time() time.Time
}

type TimeComparer[T Timed] struct {
	Val T
}

func (tc TimeComparer[T]) Compare(other TimeComparer[T]) int {
	return tc.Val.Time().Compare(other.Val.Time())
}

type TimeProvider interface {
	Now() time.Time
	NewTicker(d time.Duration) Ticker
	NewTimer(d time.Duration) Timer
	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) Timer
	Sleep(d time.Duration)
	Since(d time.Time) time.Duration
}

type Ticker interface {
	C() <-chan time.Time
	Stop()
}

type Timer interface {
	C() <-chan time.Time
	Stop()
}

type realTimeProvider struct{}

func (*realTimeProvider) AfterFunc(d time.Duration, f func()) Timer {
	return &realTimer{timer: time.AfterFunc(d, f)}
}

func (*realTimeProvider) Sleep(d time.Duration) {
	time.Sleep(d)
}

func NewRealTimeProvider() TimeProvider {
	return &realTimeProvider{}
}

func (rtp *realTimeProvider) Now() time.Time {
	return time.Now()
}

func (rtp *realTimeProvider) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (rtp *realTimeProvider) NewTicker(d time.Duration) Ticker {
	return &realTicker{ticker: time.NewTicker(d)}
}

func (rtp *realTimeProvider) NewTimer(d time.Duration) Timer {
	return &realTimer{timer: time.NewTimer(d)}
}

func (rtp *realTimeProvider) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

type realTicker struct {
	ticker *time.Ticker
}

func (rt *realTicker) C() <-chan time.Time {
	return rt.ticker.C
}

func (rt *realTicker) Stop() {
	rt.ticker.Stop()
}

type realTimer struct {
	timer *time.Timer
}

func (rt *realTimer) C() <-chan time.Time {
	return rt.timer.C
}

func (rt *realTimer) Stop() {
	rt.timer.Stop()
}

func (rt *realTimer) Reset(d time.Duration) bool {
	return rt.timer.Reset(d)
}
