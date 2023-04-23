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
		return ratelimiter.TokenBucket(1, 1, time.Second*5, mux, doneCh)
	},
	"LeakyBucket": func(mux *http.ServeMux, doneCh <-chan interface{}) http.Handler {
		return ratelimiter.LeakyBucket(1, 1, time.Second*10, mux, doneCh)
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
		time.Sleep(time.Second * 6)
		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			time.Sleep(time.Second * 2)
			res, err := http.Get(ts.URL)
			if err != nil || res.StatusCode != http.StatusTooManyRequests {
				fmt.Printf("tn:%s\n", tn)
				log.Fatal(err)
			}
			res.Body.Close()
			wg.Done()
		}()
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
