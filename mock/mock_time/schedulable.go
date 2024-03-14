package mock_time

import "time"

type schedulable interface {
	SendC() chan<- time.Time
	IsActive() bool
	Next() (time.Time, bool)
	Cancel()
}
