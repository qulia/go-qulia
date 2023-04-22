package leakybucket_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/leakybucket"
	"github.com/stretchr/testify/assert"
)

func TestLeakyBucketBasic(t *testing.T) {
	lb := leakybucket.NewLeakyBucket[int](3, 4, 2*time.Second)
	time.Sleep(3 * time.Second)
	assert.True(t, lb.Put(1))
	assert.True(t, lb.Put(2))
	assert.True(t, lb.Put(3))
	assert.False(t, lb.Put(4))

	it, ok := lb.Take()
	assert.True(t, ok)
	assert.Equal(t, 1, it)

	it, ok = lb.Take()
	assert.True(t, ok)
	assert.Equal(t, 2, it)

	it, ok = lb.Take()
	assert.True(t, ok)
	assert.Equal(t, 3, it)

	_, ok = lb.Take()
	assert.False(t, ok)
}

func TestLeakyBucketParallelProducersAndConsumers(t *testing.T) {
	lb := leakybucket.NewLeakyBucket[int](10, 5, time.Second)
	var counter int32
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			tcp := time.NewTicker(time.Millisecond * 200)
			tcc := time.NewTicker(time.Millisecond * 300)
			done := time.After(time.Second * 5)
			defer tcp.Stop()
			defer tcc.Stop()
			defer t.Logf("Exiting worker %d", i)
			defer wg.Done()
			for {
				select {
				case <-tcc.C:
					if val, ok := lb.Take(); ok {
						t.Logf("got item at %d:%d %d %d",
							time.Now().Minute(), time.Now().Second(), i, val)
					} else {
						t.Logf("did not get item %d", i)
					}
				case <-tcp.C:
					if lb.Put(int(atomic.AddInt32(&counter, 1))) {
						t.Logf("put item %d", i)
					} else {
						t.Logf("did not put item %d", i)
					}
				case <-done:
					t.Logf("Done %d", i)
					return
				}
			}
		}(i)
	}

	wg.Wait()
	lb.Close()
}
