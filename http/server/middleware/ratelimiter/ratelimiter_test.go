package ratelimiter_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/v2/http/server/middleware/ratelimiter"
	"github.com/qulia/go-qulia/v2/lib/common"
	"github.com/qulia/go-qulia/v2/mock"
	"github.com/qulia/go-qulia/v2/mock/mock_time"
	"github.com/stretchr/testify/assert"
)

var testMap = map[string]func(*http.ServeMux, <-chan interface{}, common.TimeProvider) http.Handler{
	"TokenBucket": func(mux *http.ServeMux, doneCh <-chan interface{}, mtp common.TimeProvider) http.Handler {
		return ratelimiter.TokenBucket(1, 1, time.Second*10, mux, doneCh, mtp)
	},
	"LeakyBucket": func(mux *http.ServeMux, doneCh <-chan interface{}, mtp common.TimeProvider) http.Handler {
		return ratelimiter.LeakyBucket(1, 1, time.Second*10, mux, doneCh, mtp)
	},
	"FixedWindowCounter": func(mux *http.ServeMux, doneCh <-chan interface{}, mtp common.TimeProvider) http.Handler {
		return ratelimiter.FixedWindowCounter(1, time.Second*10, mux, doneCh, mtp)
	},
	"SlidingWindowLog": func(mux *http.ServeMux, doneCh <-chan interface{}, mtp common.TimeProvider) http.Handler {
		return ratelimiter.SlidingWindowLog(1, time.Second*10, mux, doneCh, mtp)
	},
	"SlidingWindowCounter": func(mux *http.ServeMux, doneCh <-chan interface{}, mtp common.TimeProvider) http.Handler {
		return ratelimiter.SlidingWindowCounter(1, time.Minute, mux, doneCh, mtp)
	},
}

func TestRateLimiterBasic(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}
	doneCh := make(chan interface{})
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	for tn, rateLimitedHandlerFunc := range testMap {
		mtp := mock.GetMockTimeProviderLateScheduling()
		ts := httptest.NewServer(rateLimitedHandlerFunc(mux, doneCh, mtp))
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(tn string, ts *httptest.Server) {
			defer wg.Done()
			// need to make sure this routine does not continue first
			time.Sleep(time.Millisecond * 10)
			res, err := http.Get(ts.URL)
			if err != nil {
				log.Fatal(err)
			}
			// Request that run later in a different routine should be rejected
			fmt.Printf("call-separate-go-routine:%s:%d\n", tn, res.StatusCode)
			assert.Equal(t, http.StatusTooManyRequests, res.StatusCode)
			res.Body.Close()
		}(tn, ts)
		res, err := http.Get(ts.URL)
		if err != nil || res.StatusCode != http.StatusOK {
			log.Fatal(err)
		}

		res, err = http.Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("call-same-go-routine:%s:%d\n", tn, res.StatusCode)
		if tn == "LeakyBucket" {
			// for leaky bucket,
			// first call is allowed and leaks as token starts full with 1
			// since there is a buffer the second call will be allowed but only proceed after the
			// time progresses
			assert.Equal(t, http.StatusOK, res.StatusCode)
		} else {
			// Request that run later in a the same routine should be rejected
			assert.Equal(t, http.StatusTooManyRequests, res.StatusCode)
		}

		wg.Wait()
		res.Body.Close()
		ts.Close()
		mtp.(*mock_time.MockTimeProvider).Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	doneCh <- struct{}{}
}
