package ratelimiter_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/http/server/middleware/ratelimiter"
)

var testMap = map[string]func(*http.ServeMux, <-chan interface{}) http.Handler{
	"TokenBucket": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.TokenBucket(1, 1, time.Second*10, mux, doneCh)
	},
	"LeakyBucket": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.LeakyBucket(1, 1, time.Second*10, mux, doneCh)
	},
	"FixedWindowCounter": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.FixedWindowCounter(1, time.Second*10, mux, doneCh)
	},
	"SlidingWindowLog": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.SlidingWindowLog(1, time.Second*10, mux, doneCh)
	},
	"SlidingWindowCounter": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.SlidingWindowCounter(1, time.Second, mux, doneCh)
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
		ts := httptest.NewServer(rateLimitedHandlerFunc(mux, doneCh))
		defer ts.Close()

		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(tn string, ts *httptest.Server) {
			defer wg.Done()
			time.Sleep(time.Second * 2)
			res, err := http.Get(ts.URL)
			if err != nil {
				log.Fatal(err)
			}
			if res.StatusCode == http.StatusTooManyRequests {
				t.Logf("not allowed in test:%s\n", tn)
			}
			res.Body.Close()
		}(tn, ts)
		res, err := http.Get(ts.URL)
		if err != nil || res.StatusCode != http.StatusOK {
			log.Fatal(err)
		}

		wg.Wait()
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	doneCh <- struct{}{}
}
