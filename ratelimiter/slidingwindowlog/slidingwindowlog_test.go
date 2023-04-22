package slidingwindowlog_test

import (
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/slidingwindowlog"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowLogBasic(t *testing.T) {
	swl := slidingwindowlog.NewSlidingWindowLog(4, time.Second)
	assert.True(t, swl.Put())
	assert.True(t, swl.Put())

	swl = slidingwindowlog.NewSlidingWindowLog(0, time.Second)
	assert.False(t, swl.Put())
}

func TestSlidingWindowLogParallelRequestors(t *testing.T) {
	swl := slidingwindowlog.NewSlidingWindowLog(3, time.Second)

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			tc := time.NewTicker(time.Millisecond * 200)
			done := time.After(time.Second * 5)
			defer tc.Stop()
			defer t.Logf("Exiting worker %d", i)
			defer wg.Done()
			for {
				select {
				case <-tc.C:
					if swl.Put() {
						t.Logf("allowed %d", i)
					} else {
						t.Logf("not allowed %d", i)
					}
				case <-done:
					t.Logf("Done %d", i)
					return
				}
			}
		}(i)
	}

	wg.Wait()
	swl.Close()
}
