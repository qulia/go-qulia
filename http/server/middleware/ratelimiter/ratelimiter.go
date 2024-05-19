package ratelimiter

import (
	"net/http"
	"time"

	"github.com/qulia/go-qulia/v2/algo/ratelimiter"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/fixedwindowcounter"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/leakybucket"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/slidingwindowcounter"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/slidingwindowlog"
	"github.com/qulia/go-qulia/v2/algo/ratelimiter/tokenbucket"
	"github.com/qulia/go-qulia/v2/lib/common"
)

func TokenBucket(capacity int, fillAmount int, fillPeriod time.Duration,
	next http.Handler, doneCh <-chan interface{}, mtp common.TimeProvider,
) http.Handler {
	tb := tokenbucket.NewTokenBucket(capacity, fillAmount, fillPeriod, mtp)
	return handle(tb, next, doneCh)
}

func LeakyBucket(capacity int, leakAmount int, leakPeriod time.Duration,
	next http.Handler, doneCh <-chan interface{}, mtp common.TimeProvider,
) http.Handler {
	lb := leakybucket.NewLeakyBucket(capacity, leakAmount, leakPeriod, mtp)
	return handleBuffered(lb, next, doneCh)
}

func FixedWindowCounter(threshold int, window time.Duration,
	next http.Handler, doneCh <-chan interface{}, mtp common.TimeProvider,
) http.Handler {
	fwc := fixedwindowcounter.NewFixedWindowCounter(threshold, window, mtp)
	return handle(fwc, next, doneCh)
}

func SlidingWindowCounter(threshold int, window time.Duration,
	next http.Handler, doneCh <-chan interface{}, mtp common.TimeProvider,
) http.Handler {
	swc := slidingwindowcounter.NewSlidingWindowCounter(threshold, window, mtp)
	return handle(swc, next, doneCh)
}

func SlidingWindowLog(threshold int, window time.Duration,
	next http.Handler, doneCh <-chan interface{}, mtp common.TimeProvider,
) http.Handler {
	swl := slidingwindowlog.NewSlidingWindowLog(threshold, window, mtp)
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
