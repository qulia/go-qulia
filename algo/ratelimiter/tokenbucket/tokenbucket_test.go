package tokenbucket_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/v2/algo/ratelimiter/testhelper"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/tokenbucket"
	"github.com/qulia/go-qulia/v2/mock"
	"github.com/qulia/go-qulia/v2/mock/mock_time"
	"github.com/stretchr/testify/assert"
)

func TestTokenBucketBasic(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	fillAmount := 3
	tb := tokenbucket.NewTokenBucket(10, fillAmount, time.Minute, mtp)
	defer tb.Close()
	for i := 0; i < fillAmount; i++ {
		assert.True(t, tb.Allow())
	}
	assert.False(t, tb.Allow())
}

func TestCallAfterClose(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	fillAmount := 3
	tb := tokenbucket.NewTokenBucket(10, fillAmount, time.Minute, mtp)
	assert.True(t, tb.Allow())
	tb.Close()
	assert.False(t, tb.Allow())
}

func TestTokenOverflow(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	tb := tokenbucket.NewTokenBucket(10, 8, time.Second, mtp)
	assert.True(t, tb.Allow())
	mtp.Sleep(time.Second * 2)
	assert.True(t, tb.Allow())
}

func TestTokenBucketParallelRequestors(t *testing.T) {
	mtp := mock.GetMockTimeProviderDefault()
	defer mtp.(*mock_time.MockTimeProvider).Close()
	tb := tokenbucket.NewTokenBucket(10, 5, time.Second, mtp)
	defer tb.Close()

	testhelper.RunWorkers(t, tb, mtp)
}
