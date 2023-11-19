package testhelper

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/algo/ratelimiter"
)

func RunWorkers(t *testing.T, rl ratelimiter.RateLimiter) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			tc := time.NewTicker(time.Millisecond * 200)
			done := time.After(time.Second * 5)
			defer tc.Stop()
			defer fmt.Printf("Exiting worker %d\n", i)
			defer wg.Done()
			for {
				select {
				case <-tc.C:
					if rl.Allow() {
						fmt.Printf("allowed %d:%d %d\n",
							time.Now().Minute(), time.Now().Second(), i)
					} else {
						fmt.Printf("not allowed %d:%d %d\n",
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

func RunWorkersBuffered(t *testing.T, rl ratelimiter.RateLimiterBuffered) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			tc := time.NewTicker(time.Millisecond * 200)
			done := time.After(time.Second * 5)
			receiverWg := &sync.WaitGroup{}
			defer tc.Stop()
			defer fmt.Printf("Exiting worker %d\n", i)
			defer receiverWg.Wait()
			defer wg.Done()
			for {
				select {
				case <-tc.C:
					if ch, ok := rl.Allow(); ok {
						fmt.Printf("allowed %d:%d %d\n",
							time.Now().Minute(), time.Now().Second(), i)
						receiverWg.Add(1)
						go func(ch <-chan interface{}) {
							<-ch
							fmt.Printf("received %d:%d %d\n",
								time.Now().Minute(), time.Now().Second(), i)
							receiverWg.Done()
						}(ch)

					} else {
						fmt.Printf("not allowed %d:%d %d\n",
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
