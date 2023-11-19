package tokenbucket_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter/testhelper"
	"github.com/qulia/go-qulia/algo/ratelimiter/tokenbucket"
	"github.com/stretchr/testify/assert"
)

func TestTokenBucketBasic(t *testing.T) {
	fillAmount := 3
	tb := tokenbucket.NewTokenBucket(10, fillAmount, time.Minute)
	defer tb.Close()
	for i := 0; i < fillAmount; i++ {
		assert.True(t, tb.Allow())
	}
	assert.False(t, tb.Allow())
}

func TestCallAfterClose(t *testing.T) {
	fillAmount := 3
	tb := tokenbucket.NewTokenBucket(10, fillAmount, time.Minute)
	assert.True(t, tb.Allow())
	tb.Close()
	assert.False(t, tb.Allow())
}

func TestTokenOverflow(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 8, time.Second)
	assert.True(t, tb.Allow())
	time.Sleep(time.Second * 2)
	assert.True(t, tb.Allow())
}

func TestTokenBucketParallelRequestors(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 5, time.Second)
	defer tb.Close()

	testhelper.RunWorkers(t, tb)
}
