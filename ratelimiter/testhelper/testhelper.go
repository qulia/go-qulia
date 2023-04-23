package testhelper

import (
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/ratelimiter"
)

func RunWokers(t *testing.T, rl ratelimiter.RateLimiter) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			tc := time.NewTicker(time.Millisecond * 200)
			done := time.After(time.Second * 5)
			defer tc.Stop()
			defer t.Logf("Exiting worker %d", i)
			defer wg.Done()
			for {
				select {
				case <-tc.C:
					if rl.Allow() {
						t.Logf("allowed %d:%d %d",
							time.Now().Minute(), time.Now().Second(), i)
					} else {
						t.Logf("not allowed %d:%d %d",
							time.Now().Minute(), time.Now().Second(), i)
					}
				case <-done:
					return
				}
			}
		}(i)
	}

	wg.Wait()
}

func RunWokersBuffered(t *testing.T, rl ratelimiter.RateLimiterBuffered) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			tc := time.NewTicker(time.Millisecond * 200)
			done := time.After(time.Second * 5)
			receiverWg := &sync.WaitGroup{}
			defer tc.Stop()
			defer t.Logf("Exiting worker %d", i)
			defer receiverWg.Wait()
			defer wg.Done()
			for {
				select {
				case <-tc.C:
					if ch, ok := rl.Allow(); ok {
						t.Logf("allowed %d:%d %d",
							time.Now().Minute(), time.Now().Second(), i)
						receiverWg.Add(1)
						go func(ch <-chan interface{}) {
							<-ch
							t.Logf("received %d:%d %d",
								time.Now().Minute(), time.Now().Second(), i)
							receiverWg.Done()
						}(ch)

					} else {
						t.Logf("not allowed %d:%d %d",
							time.Now().Minute(), time.Now().Second(), i)
					}
				case <-done:
					return
				}
			}
		}(i)
	}

	wg.Wait()
}
