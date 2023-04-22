package fixedwindowcounter_test

import (
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/fixedwindowcounter"
	"github.com/stretchr/testify/assert"
)

func TestFixedWindowCounter(t *testing.T) {
	fwc := fixedwindowcounter.NewFixedWindowCounter(3, time.Second*2)
	time.Sleep(time.Second * 5)
	assert.True(t, fwc.Put())
	assert.True(t, fwc.Put())
	fwc.Close()
}

func TestFixedWindowCounterParallelRequestors(t *testing.T) {
	fwc := fixedwindowcounter.NewFixedWindowCounter(3, time.Second)

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
					if fwc.Put() {
						t.Logf("got token %d", i)
					} else {
						t.Logf(" %d", i)
					}
				case <-done:
					t.Logf("Done %d", i)
					return
				}
			}
		}(i)
	}

	wg.Wait()
	fwc.Close()
}
