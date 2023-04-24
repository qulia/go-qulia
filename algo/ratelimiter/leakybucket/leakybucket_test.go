package leakybucket_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter/leakybucket"
	"github.com/qulia/go-qulia/algo/ratelimiter/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestLeakyBucketBasic(t *testing.T) {
	leakAmount := 4
	lb := leakybucket.NewLeakyBucket(5, leakAmount, time.Minute)
	defer lb.Close()
	for i := 0; i < leakAmount; i++ {
		ch, ok := lb.Allow()
		assert.True(t, ok)
		<-ch
	}
}

func TestLeakyBucketCapOne(t *testing.T) {
	lb := leakybucket.NewLeakyBucket(1, 1, time.Second*10)
	defer lb.Close()
	ch, ok := lb.Allow()
	assert.True(t, ok)
	<-ch
	_, ok = lb.Allow()
	assert.True(t, ok)
}

func TestCallAfterClose(t *testing.T) {
	lb := leakybucket.NewLeakyBucket(1, 1, time.Second*10)
	lb.Close()
	_, ok := lb.Allow()
	assert.False(t, ok)
}

func TestLeakyBucketParallelProducersAndConsumers(t *testing.T) {
	lb := leakybucket.NewLeakyBucket(10, 5, time.Second)
	defer lb.Close()
	testhelper.RunWokersBuffered(t, lb)
}
