package slidingwindowlog_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/slidingwindowlog"
	"github.com/qulia/go-qulia/ratelimiter/testhelper"
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

func TestSlidingWindowLogParallelRequestors(t *testing.T) {
	swl := slidingwindowlog.NewSlidingWindowLog(3, time.Second)
	defer swl.Close()
	testhelper.RunWokers(t, swl)
}
