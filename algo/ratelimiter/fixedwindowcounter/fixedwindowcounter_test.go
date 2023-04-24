package fixedwindowcounter_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter/fixedwindowcounter"
	"github.com/qulia/go-qulia/algo/ratelimiter/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestFixedWindowCounter(t *testing.T) {
	threshold := 3
	fwc := fixedwindowcounter.NewFixedWindowCounter(3, time.Minute)
	defer fwc.Close()
	for i := 0; i < threshold; i++ {
		assert.True(t, fwc.Allow())
	}
}

func TestFixedWindowCounterParallelRequestors(t *testing.T) {
	fwc := fixedwindowcounter.NewFixedWindowCounter(3, time.Second)
	defer fwc.Close()

	testhelper.RunWokers(t, fwc)
}
