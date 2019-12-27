// This example demonstrates a job priority queue built using the heap interface.
package heap_test

import (
	"fmt"
	"github.com/qulia/go-qulia/lib/heap"
)

type job struct{
	priority int
	name string
	department string
}

var (
	jobCompFunc = func(first, second interface{}) int {
		firstJob :=first.(job)
		secondJob := second.(job)
		return heap.IntCompFunc(firstJob.priority, secondJob.priority)
	}
)

// This example initializes the heap with list of jobs and pushes another one with Insert method
// With the provided comparison method, the jobs with low priority ones are extracted first
func ExampleMinHeap()  {
	jobs := []interface{} {
		job{
			priority:   4,
			name:       "JobA",
			department: "DeptA",
		},
		job{
			priority:   1,
			name:       "JobB",
			department: "DeptA",
		},
		job{
			priority:   0,
			name:       "JobZ",
			department: "DeptC",
		},
		job{
			priority:   7,
			name:       "JobH",
			department: "DeptA",
		},
	}

	jobHeap := heap.NewMinHeap(jobs, jobCompFunc)

	jobHeap.Insert(job{
		priority:   5,
		name:       "JobJ",
		department: "DeptX",
	})

	for jobHeap.Size() != 0 {
		fmt.Printf("Current job %v\n", jobHeap.Extract().(job))
	}

	// Output:
	// Current job {0 JobZ DeptC}
	// Current job {1 JobB DeptA}
	// Current job {4 JobA DeptA}
	// Current job {5 JobJ DeptX}
	// Current job {7 JobH DeptA}
}