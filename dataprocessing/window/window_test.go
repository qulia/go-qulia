package window

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/v2/lib/common"
	"github.com/qulia/go-qulia/v2/lib/set"
	"github.com/qulia/go-qulia/v2/mock"
	"github.com/qulia/go-qulia/v2/mock/mock_time"
	"github.com/stretchr/testify/assert"
)

type Event struct {
	Name     string
	Data     string
	Occurred time.Time
}

func (e Event) Time() time.Time {
	return e.Occurred
}

type EventAggregator struct {
	s set.Set[string]
}

func (ea EventAggregator) Add(e Event) {
	ea.s.Add(e.Name)
}

func (ea EventAggregator) Remove(e Event) {
	ea.s.Remove(e.Name)
}

func (ea EventAggregator) Result() []string {
	res := ea.s.ToSlice()
	sort.Strings(res)
	return res
}

func TestBasicSlidingWindowedAgg(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	wa := NewSlidingWindow[Event, []string, EventAggregator](
		time.Second*3, time.Second, EventAggregator{s: set.NewSet[string]()}, mtp)
	testWithWindow(wa, mtp, t)
}

func TestFixedWindowedAgg(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	wa := NewFixedWindow[Event, []string, EventAggregator](
		time.Second*2, EventAggregator{s: set.NewSet[string]()}, mtp)
	testWithWindow(wa, mtp, t)
}

func testWithWindow(wa Window[Event, []string, EventAggregator], mtp common.TimeProvider, t *testing.T) {
	startTime := mtp.Now()
	expectedResults := sync.Map{}
	go func() {
		for r := range wa.Results() {
			log.Printf("Result: @%dms %v\n", int(r.Start.Sub(startTime).Milliseconds()), r.Value)

			for _, v := range r.Value {
				_, exists := expectedResults.Load(v)
				assert.True(t, exists)
			}
		}
	}()

	for i := 0; i < 20; i++ {
		occurred := mtp.Now()

		name := fmt.Sprintf("test-%c-%d", 'a'+(i/10), i)
		expectedAdded := true
		if i%5 == 0 {
			// also send late events
			occurred = occurred.Add(-time.Second * 6)
			expectedAdded = false
		} else {
			expectedResults.Store(name, true)
		}
		added := wa.Add(Event{Name: name, Data: "data", Occurred: occurred})
		assert.Equal(t, expectedAdded, added)
		mtp.Sleep(time.Second / 2)
	}
}
