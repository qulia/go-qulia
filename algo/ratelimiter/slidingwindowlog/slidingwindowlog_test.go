package slidingwindowlog_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter/slidingwindowlog"
	"github.com/qulia/go-qulia/algo/ratelimiter/testhelper"
	"github.com/qulia/go-qulia/mock"
	"github.com/qulia/go-qulia/mock/mock_time"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowLogBasic(t *testing.T) {
	threshold := 4
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	swl := slidingwindowlog.NewSlidingWindowLog(threshold, time.Minute, mtp)
	for i := 0; i < threshold; i++ {
		assert.True(t, swl.Allow())
	}
	swl.Close()

	swl = slidingwindowlog.NewSlidingWindowLog(0, time.Second, mtp)
	assert.False(t, swl.Allow())
	swl.Close()
}

func TestCallAfterClose(t *testing.T) {
	threshold := 4

	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	swl := slidingwindowlog.NewSlidingWindowLog(threshold, time.Minute, mtp)
	assert.True(t, swl.Allow())
	swl.Close()
	assert.False(t, swl.Allow())
}

func TestSlidingWindowLogParallelRequestors(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	swl := slidingwindowlog.NewSlidingWindowLog(3, time.Second, mtp)
	defer swl.Close()
	testhelper.RunWorkers(t, swl, mtp)
}
