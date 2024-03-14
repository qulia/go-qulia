package window

import (
	"fmt"
	"sort"
	"time"

	"github.com/qulia/go-qulia/lib/common"
)

type PlayEvent struct {
	SongPlayCount
	Occurred time.Time
}

type SongPlayCount struct {
	Song  string
	Count int
}

func (p SongPlayCount) Compare(other SongPlayCount) int {
	if p.Song == other.Song {
		return p.Count - other.Count
	}
	if p.Song < other.Song {
		return -1
	}

	return 1
}

type PlayCountAggregator struct {
	sm map[string]int
}

func NewPlayCountAggregator() PlayCountAggregator {
	return PlayCountAggregator{make(map[string]int)}
}

func (pca PlayCountAggregator) Add(pe PlayEvent) {
	pca.sm[pe.Song] += pe.Count
}

func (tka PlayCountAggregator) Remove(pe PlayEvent) {
	tka.sm[pe.Song] -= pe.Count
}

func (tka PlayCountAggregator) Result() []SongPlayCount {
	var songPlayCounts []SongPlayCount
	for k, v := range tka.sm {
		songPlayCounts = append(songPlayCounts, SongPlayCount{k, v})
	}
	return songPlayCounts
}

func (p PlayEvent) Time() time.Time {
	return p.Occurred
}

func ExampleWindow_topK() {
	sw := NewSlidingWindow[PlayEvent, []SongPlayCount, PlayCountAggregator](
		time.Second*3, time.Second, NewPlayCountAggregator(), common.NewRealTimeProvider())
	startTime := time.Now()
	go func() {
		for r := range sw.Results() {
			v := r.Value
			// Sort for consistent output
			sort.Slice(v, func(i, j int) bool { return v[i].Compare(v[j]) < 0 })
			fmt.Printf("Result: @%dms %v\n", int(r.Start.Sub(startTime).Milliseconds()), v)
		}
	}()

	for i := 0; i < 20; i++ {
		sw.Add(PlayEvent{SongPlayCount{fmt.Sprintf("song%d", i%3), 1}, time.Now()})
		time.Sleep(time.Second / 2)
	}
	// Output:
	// Result: @0ms [{song0 2} {song1 2} {song2 2}]
	// Result: @999ms [{song0 2} {song1 2} {song2 2}]
	// Result: @1999ms [{song0 2} {song1 2} {song2 2}]
	// Result: @2999ms [{song0 2} {song1 2} {song2 2}]
	// Result: @3999ms [{song0 2} {song1 2} {song2 2}]
	// Result: @4999ms [{song0 2} {song1 2} {song2 2}]
	// Result: @5999ms [{song0 2} {song1 2} {song2 2}]
}
