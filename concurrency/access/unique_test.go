package access_test

import (
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
)

type job struct {
	id   int
	name string
}

func TestUniqueBasic(t *testing.T) {
	// Create a queue can be accessed exclusively only by one go routine
	key := access.NewUnique(queue.NewQueue[job]())
	key.Release()
	receiveTimeout := 5 * time.Second
	runSenders(key)

	// Run consumer
	receiverWg := &sync.WaitGroup{}
	receiverWg.Add(1)
	go func() {
		defer log.Printf("Exiting consumer")
		defer receiverWg.Done()
		for {
			// Acquire the job queue before consuming from
			jobQueue, open := key.Acquire()
			if !open {
				return
			}
			if !jobQueue.IsEmpty() {
				job := jobQueue.Dequeue()
				log.Printf("Processing job %v", job)
			} else {
				key.Release()
				<-time.After(receiveTimeout)
				jobQueue, open := key.Acquire()
				if !open {
					return
				}

				if jobQueue.IsEmpty() {
					key.Release()
					break
				}
			}
			key.Release()
		}
	}()

	// drain the queue
	receiverWg.Wait()

	// close access to the queue for both send and receive
	key.Close()
}

func runSenders(key *access.Unique[queue.Queue[job]]) {
	senderWg := &sync.WaitGroup{}
	senderWg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer senderWg.Done()
			jobQueue, ok := key.Acquire()
			if !ok {
				// key is closed, done
				return
			}
			job := job{
				id:   i,
				name: "job" + strconv.Itoa(i),
			}
			log.Printf("Queuing job %v", job)
			jobQueue.Enqueue(job)
			key.Release()
		}(i)
	}

	senderWg.Wait()
}
