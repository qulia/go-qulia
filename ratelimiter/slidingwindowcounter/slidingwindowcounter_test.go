package slidingwindowcounter_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/slidingwindowcounter"
	"github.com/qulia/go-qulia/ratelimiter/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowcounterBasic(t *testing.T) {
	cap := 7
	swc := slidingwindowcounter.NewSlidingWindowCounter(cap, time.Minute)
	defer swc.Close()
	for i := 0; i < cap; i++ {
		assert.True(t, swc.Allow())
	}
}

func TestSlidingWindowCounterParallelRequestors(t *testing.T) {
	swc := slidingwindowcounter.NewSlidingWindowCounter(3, time.Second)
	defer swc.Close()
	testhelper.RunWokers(t, swc)
}
