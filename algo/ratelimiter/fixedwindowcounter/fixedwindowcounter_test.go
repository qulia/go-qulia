package fixedwindowcounter_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter/fixedwindowcounter"
	"github.com/qulia/go-qulia/algo/ratelimiter/testhelper"
	"github.com/qulia/go-qulia/mock"
	"github.com/qulia/go-qulia/mock/mock_time"
	"github.com/stretchr/testify/assert"
)

func TestFixedWindowCounter(t *testing.T) {
	threshold := 3

	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	fwc := fixedwindowcounter.NewFixedWindowCounter(3, time.Minute, mtp)
	defer fwc.Close()
	for i := 0; i < threshold; i++ {
		assert.True(t, fwc.Allow())
	}
}

func TestFixedWindowCounterParallelRequestors(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	fwc := fixedwindowcounter.NewFixedWindowCounter(3, time.Second, mtp)
	defer fwc.Close()

	testhelper.RunWorkers(t, fwc, mtp)
}
