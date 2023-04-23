package ratelimiter

import (
	"net/http"
	"time"

	"github.com/qulia/go-qulia/ratelimiter"
	"github.com/qulia/go-qulia/ratelimiter/fixedwindowcounter"
	"github.com/qulia/go-qulia/ratelimiter/leakybucket"
	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
)

func TokenBucket(capacity int, fillAmount int, fillPeriod time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	tb := tokenbucket.NewTokenBucket(capacity, fillAmount, fillPeriod)
	go func(tb ratelimiter.RateLimiter) {
		<-doneCh
		tb.Close()
	}(tb)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tb.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LeakyBucket(capacity int, leakAmount int, leakPeriod time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	lb := leakybucket.NewLeakyBucket[*http.Request](capacity, leakAmount, leakPeriod)
	go func(lb ratelimiter.RateLimiterBuffered) {
		<-doneCh
		lb.Close()
	}(lb)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ch, ok := lb.Allow(); ok {
			// This is a blocking call until it leaks from the bucket
			<-ch
		} else {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func FixedWindowCounter(threshold int, window time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	fwc := fixedwindowcounter.NewFixedWindowCounter(threshold, window)
	go func(fwc ratelimiter.RateLimiter) {
		<-doneCh
		fwc.Close()
	}(fwc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !fwc.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
