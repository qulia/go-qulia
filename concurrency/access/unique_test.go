package access_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/qulia/go-qulia/concurrency/access"
	"github.com/qulia/go-qulia/lib/queue"
	log "github.com/sirupsen/logrus"
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
		for {
			// Acquire the job queue before consuming
			jobQueue := jobQueueUnique.Acquire()
			if jobQueue == nil {
				jobQueueUnique.Release()
				return
			}

			job := jobQueue.(*queue.Queue).Dequeue()
			log.Infof("Processing job %v", job)
			jobQueueUnique.Release()
			log.Infof("Done processing job %v", job)
		}

		log.Infof("Exiting consumer")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			jobQueue := jobQueueUnique.Acquire()
			job := job{
				id:   i,
				name: "job" + strconv.Itoa(i),
			}
			log.Infof("Queuing job %v", job)
			jobQueue.(*queue.Queue).Enqueue(job)
			jobQueueUnique.Release()
		}

		log.Infof("Done queuing jobs")
		jobQueueUnique.Done()
		log.Infof("Exiting producer")
	}()

	jobQueueUnique.Release()
	time.Sleep(time.Second * 2)
}
