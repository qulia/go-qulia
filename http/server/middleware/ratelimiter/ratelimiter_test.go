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
	ratePerSecond = 8
)

var testMap = map[string]func(*http.ServeMux) http.Handler{
	"TokenBucket": func(mux *http.ServeMux) http.Handler {
		return ratelimiter.TokenBucket(ratePerSecond*2, ratePerSecond, time.Second, mux)
	},
}

func TestRateLimiterBasic(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	rateLimitedHandler := ratelimiter.TokenBucket(1, 1, time.Second*5, mux)
	ts := httptest.NewServer(rateLimitedHandler)
	defer ts.Close()
	time.Sleep(time.Second * 6)
	res, err := http.Get(ts.URL)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	res, err = http.Get(ts.URL)
	if err != nil || res.StatusCode != http.StatusTooManyRequests {
		log.Fatal(err)
	}
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkRateLimit(b *testing.B) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	for limiter, handler := range testMap {
		b.Run(limiter, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runBenchmark(handler(mux), b)
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
