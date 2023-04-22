package tokenbucket_test

import (
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
	"github.com/stretchr/testify/assert"
)

func TestTokenBucketBasic(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 3, time.Second*2)
	time.Sleep(time.Second * 5)
	assert.True(t, tb.Take())
	tb.Close()
}

func TestTokenBucketParallelRequestors(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 5, time.Second)

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
					if tb.Take() {
						t.Logf("got token %d", i)
					} else {
						t.Logf("did not get token %d", i)
					}
				case <-done:
					t.Logf("Done %d", i)
					return
				}
			}
		}(i)
	}

	wg.Wait()
	tb.Close()
}
