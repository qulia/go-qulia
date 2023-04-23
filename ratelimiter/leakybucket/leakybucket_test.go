package leakybucket_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/leakybucket"
	"github.com/qulia/go-qulia/ratelimiter/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestLeakyBucketBasic(t *testing.T) {
	leakAmount := 4
	lb := leakybucket.NewLeakyBucket[int](5, leakAmount, time.Minute)
	defer lb.Close()
	for i := 0; i < leakAmount; i++ {
		ch, ok := lb.Allow()
		assert.True(t, ok)
		<-ch
	}
}

func TestLeakyBucketParallelProducersAndConsumers(t *testing.T) {
	lb := leakybucket.NewLeakyBucket[int](10, 5, time.Second)
	defer lb.Close()
	testhelper.RunWokersBuffered(t, lb)
}
