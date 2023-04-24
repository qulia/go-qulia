package ratelimiter

type RateLimiter interface {
	Allow() bool
	Close()
}

type RateLimiterBuffered interface {
	Allow() (<-chan interface{}, bool)
	Close()
}
