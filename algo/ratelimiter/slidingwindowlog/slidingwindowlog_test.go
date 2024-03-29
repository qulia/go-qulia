package slidingwindowlog_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter/slidingwindowlog"
	"github.com/qulia/go-qulia/algo/ratelimiter/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowLogBasic(t *testing.T) {
	threshold := 4
	swl := slidingwindowlog.NewSlidingWindowLog(threshold, time.Minute)
	for i := 0; i < threshold; i++ {
		assert.True(t, swl.Allow())
	}
	swl.Close()

	swl = slidingwindowlog.NewSlidingWindowLog(0, time.Second)
	assert.False(t, swl.Allow())
	swl.Close()
}

func TestCallAfterClose(t *testing.T) {
	threshold := 4
	swl := slidingwindowlog.NewSlidingWindowLog(threshold, time.Minute)
	assert.True(t, swl.Allow())
	swl.Close()
	assert.False(t, swl.Allow())
}

func TestSlidingWindowLogParallelRequestors(t *testing.T) {
	swl := slidingwindowlog.NewSlidingWindowLog(3, time.Second)
	defer swl.Close()
	testhelper.RunWorkers(t, swl)
}
