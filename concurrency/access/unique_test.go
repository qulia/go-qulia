package access_test

import (
	"log"
	"strconv"
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
	jobQueueUnique := access.NewUnique(queue.NewQueue())

	// Run consumer
	go func() {
		defer log.Printf("Exiting consumer")
		for {
			// Acquire the job queue before consuming
			jobQueue := jobQueueUnique.Acquire()
			if jobQueue == nil {
				jobQueueUnique.Release()
				return
			}

			job := jobQueue.(*queue.Queue).Dequeue()
			log.Printf("Processing job %v", job)
			jobQueueUnique.Release()
			log.Printf("Done processing job %v", job)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			jobQueue := jobQueueUnique.Acquire()
			job := job{
				id:   i,
				name: "job" + strconv.Itoa(i),
			}
			log.Printf("Queuing job %v", job)
			jobQueue.(*queue.Queue).Enqueue(job)
			jobQueueUnique.Release()
		}

		log.Printf("Done queuing jobs")
		jobQueueUnique.Done()
		log.Printf("Exiting producer")
	}()

	jobQueueUnique.Release()
	time.Sleep(time.Second * 2)
}
