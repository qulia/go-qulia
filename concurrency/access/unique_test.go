package access_test

import (
	"log"
	"strconv"
	"sync"
	"testing"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
)

type job struct {
	id   int
	name string
}

func TestUniqueBasic(t *testing.T) {
	// Create a queue can be accessed exclusively only by one go routine
	accessHolder := access.NewUnique(queue.NewQueue[job]())
	accessHolder.Release()
	// Run consumer
	go func() {
		defer log.Printf("Exiting consumer")
		for {
			// Acquire the job queue before consuming from
			jobQueue, ok := accessHolder.Acquire()
			if !ok {
				return
			}

			if !jobQueue.IsEmpty() {
				job := jobQueue.Dequeue()
				log.Printf("Processing job %v", job)
			}
			accessHolder.Release()
		}
	}()

	senderWg := &sync.WaitGroup{}
	senderWg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			jobQueue, ok := accessHolder.Acquire()
			if !ok {
				return
			}
			job := job{
				id:   i,
				name: "job" + strconv.Itoa(i),
			}
			log.Printf("Queuing job %v", job)
			jobQueue.Enqueue(job)
			accessHolder.Release()
			senderWg.Done()
		}(i)
	}

	senderWg.Wait()
	accessHolder.Close()
}
