package ratelimiter

import (
	"net/http"
	"time"

	"github.com/qulia/go-qulia/ratelimiter/tokenbucket"
)

func TokenBucket(capacity int, fillAmount int, fillPeriod time.Duration, next http.Handler) http.Handler {
	tb := tokenbucket.NewTokenBucket(capacity, fillAmount, fillPeriod)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tb.Take() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
