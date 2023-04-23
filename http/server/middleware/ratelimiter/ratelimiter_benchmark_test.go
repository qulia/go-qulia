package ratelimiter_test

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/qulia/go-qulia/http/server/middleware/ratelimiter"
)

const (
	ratePerSecond = 6
)

var benchMap = map[string]func(*http.ServeMux, <-chan interface{}) http.Handler{
	"TokenBucket": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.TokenBucket(ratePerSecond*2, ratePerSecond, time.Second, mux, doneCh)
	},
	"LeakyBucket": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.LeakyBucket(ratePerSecond*2, ratePerSecond, time.Second, mux, doneCh)
	},
	"FixedWindowCounter": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.FixedWindowCounter(ratePerSecond, time.Second, mux, doneCh)
	},
	"SlidingWindowLog": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.SlidingWindowLog(ratePerSecond, time.Second, mux, doneCh)
	},
	"SlidingWindowCounter": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.SlidingWindowCounter(ratePerSecond, time.Second, mux, doneCh)
	},
}

func BenchmarkRateLimit(b *testing.B) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}
	doneCh := make(chan interface{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	for limiter, handler := range benchMap {
		b.Run(limiter, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runBenchmark(handler(mux, doneCh), b)
			}
		})
	}
}

func runBenchmark(rateLimitedHandler http.Handler, b *testing.B) {
	ts := httptest.NewServer(rateLimitedHandler)
	defer ts.Close()

	time.Sleep(time.Second)
	var allowedCount, droppedCount int32
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			time.Sleep(time.Duration(rand.Uint32()%5) * time.Second)
			res, err := http.Get(ts.URL)
			if err != nil {
				log.Fatal(err)
			}

			if res.StatusCode == http.StatusOK {
				atomic.AddInt32(&allowedCount, 1)
			} else {
				atomic.AddInt32(&droppedCount, 1)
			}
			wg.Done()
		}(wg)
	}

	wg.Wait()
	b.ReportMetric(float64(ratePerSecond), "expectedRequestsAllowed/sec")
	b.ReportMetric(float64(allowedCount)/b.Elapsed().Seconds(), "actualRequestsAllowed/sec")
	b.Logf("Completed with %d/%d\n", allowedCount, droppedCount+allowedCount)
}
