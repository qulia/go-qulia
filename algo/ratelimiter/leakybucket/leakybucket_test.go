package leakybucket_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/v2/algo/ratelimiter/leakybucket"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/testhelper"
	"github.com/qulia/go-qulia/v2/mock"
	"github.com/qulia/go-qulia/v2/mock/mock_time"
	"github.com/stretchr/testify/assert"
)

func TestLeakyBucketBasic(t *testing.T) {
	leakAmount := 4

	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	lb := leakybucket.NewLeakyBucket(5, leakAmount, time.Minute, mtp)
	defer lb.Close()
	for i := 0; i < leakAmount; i++ {
		ch, ok := lb.Allow()
		assert.True(t, ok)
		<-ch
	}
}

func TestLeakyBucketCapOne(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	lb := leakybucket.NewLeakyBucket(1, 1, time.Second*10, mtp)
	defer lb.Close()
	ch, ok := lb.Allow()
	assert.True(t, ok)
	<-ch
	_, ok = lb.Allow()
	assert.True(t, ok)
}

func TestCallAfterClose(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	lb := leakybucket.NewLeakyBucket(1, 1, time.Second*10, mtp)
	lb.Close()
	_, ok := lb.Allow()
	assert.False(t, ok)
}

func TestLeakyBucketParallelProducersAndConsumers(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	lb := leakybucket.NewLeakyBucket(10, 5, time.Second, mtp)
	defer lb.Close()
	testhelper.RunWorkersBuffered(t, lb, mtp)
}
