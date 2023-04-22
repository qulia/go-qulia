package slidingwindowcounter_test

import (
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/slidingwindowcounter"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowcounterBasic(t *testing.T) {
	swc := slidingwindowcounter.NewSlidingWindowCounter(7, time.Minute)
	assert.True(t, swc.Put())
	assert.True(t, swc.Put())
	assert.True(t, swc.Put())
	assert.True(t, swc.Put())
}

func TestSlidingWindowCounterParallelRequestors(t *testing.T) {
	swc := slidingwindowcounter.NewSlidingWindowCounter(3, time.Second)

	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			tc := time.NewTicker(time.Millisecond * 500)
			done := time.After(time.Second * 10)
			defer tc.Stop()
			defer t.Logf("Exiting worker %d", i)
			defer wg.Done()
			for {
				select {
				case <-tc.C:
					if swc.Put() {
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
	swc.Close()
}
