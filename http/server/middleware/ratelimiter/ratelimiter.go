package ratelimiter

import (
	"net/http"
	"time"

	"github.com/qulia/go-qulia/ratelimiter"
	"github.com/qulia/go-qulia/ratelimiter/fixedwindowcounter"
	"github.com/qulia/go-qulia/ratelimiter/leakybucket"
	"github.com/qulia/go-qulia/ratelimiter/slidingwindowcounter"
	"github.com/qulia/go-qulia/ratelimiter/slidingwindowlog"
	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
)

func TokenBucket(capacity int, fillAmount int, fillPeriod time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	tb := tokenbucket.NewTokenBucket(capacity, fillAmount, fillPeriod)
	return handle(tb, next, doneCh)
}

func LeakyBucket(capacity int, leakAmount int, leakPeriod time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	lb := leakybucket.NewLeakyBucket(capacity, leakAmount, leakPeriod)
	return handleBuffered(lb, next, doneCh)
}

func FixedWindowCounter(threshold int, window time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	fwc := fixedwindowcounter.NewFixedWindowCounter(threshold, window)
	return handle(fwc, next, doneCh)
}

func SlidingWindowCounter(threshold int, window time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	swc := slidingwindowcounter.NewSlidingWindowCounter(threshold, window)
	return handle(swc, next, doneCh)
}

func SlidingWindowLog(threshold int, window time.Duration, next http.Handler, doneCh <-chan interface{}) http.Handler {
	swl := slidingwindowlog.NewSlidingWindowLog(threshold, window)
	return handle(swl, next, doneCh)
}

func handle(rl ratelimiter.RateLimiter, next http.Handler, doneCh <-chan interface{}) http.Handler {
	go func(rl ratelimiter.RateLimiter) {
		<-doneCh
		rl.Close()
	}(rl)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleBuffered(lb ratelimiter.RateLimiterBuffered, next http.Handler, doneCh <-chan interface{}) http.Handler {
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
