package mock_time

import "time"

type schedule struct {
	scheduleAt time.Time
	sch        schedulable
}

func (sch schedule) Compare(schOther schedule) int {
	return int(sch.scheduleAt.Sub(schOther.scheduleAt))
}
