package tokenbucket_test

import (
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/testhelper"
	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
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

func TestTokenBucketParallelRequestors(t *testing.T) {
	tb := tokenbucket.NewTokenBucket(10, 5, time.Second)
	defer tb.Close()

	testhelper.RunWokers(t, tb)
}
