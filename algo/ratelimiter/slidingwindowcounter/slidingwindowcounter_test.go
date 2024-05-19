package slidingwindowcounter_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/v2/algo/ratelimiter/slidingwindowcounter"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/testhelper"
	"github.com/qulia/go-qulia/v2/mock"
	"github.com/qulia/go-qulia/v2/mock/mock_time"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowcounterBasic(t *testing.T) {
	cap := 7
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	swc := slidingwindowcounter.NewSlidingWindowCounter(cap, time.Minute, mtp)
	defer swc.Close()
	for i := 0; i < cap; i++ {
		assert.True(t, swc.Allow())
	}
}

func TestCallAfterClose(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	swc := slidingwindowcounter.NewSlidingWindowCounter(700, time.Hour, mtp)
	assert.True(t, swc.Allow())
	swc.Close()
	assert.False(t, swc.Allow())
}

func TestSlidingWindowCounterParallelRequestors(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	swc := slidingwindowcounter.NewSlidingWindowCounter(3, time.Second, mtp)
	defer swc.Close()
	testhelper.RunWorkers(t, swc, mtp)
}
