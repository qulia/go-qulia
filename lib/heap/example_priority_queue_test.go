// This example demonstrates a job priority queue built using the heap interface.
package heap_test

import (
	"fmt"

	"github.com/qulia/go-qulia/lib/heap"
)

type job struct {
	priority   int
	name       string
	department string
}

func (j job) Less(other job) bool {
	return j.priority < other.priority
}

// This example initializes the max heap with slice of ints
// With < operator as comparison the values are extracted in descending order
// Any type met by constraints.Ordered can be used for the content of the heap
func ExampleHeap() {
	iHeap := heap.NewMaxHeap([]int{3, 7, 4, 4, 1}) // Heap[int]
	iHeap.Insert(9)

	for !iHeap.IsEmpty() {
		fmt.Printf("Out: %d\n", iHeap.Extract())
	}

	// Output:
	// Out: 9
	// Out: 7
	// Out: 4
	// Out: 4
	// Out: 3
	// Out: 1
}

// This example initializes the heap with list of jobs and pushes another one with Insert method
// With the provided comparison method Less on the type implementing lib.Lesser[T]
// depending on the heap type (min/max) the jobs will be extracted in order
func ExampleHeapCustomComp() {
	jobs := []job{
		{
			priority:   4,
			name:       "JobA",
			department: "DeptA",
		},
		{
			priority:   1,
			name:       "JobB",
			department: "DeptA",
		},
		{
			priority:   0,
			name:       "JobZ",
			department: "DeptC",
		},
		{
			priority:   7,
			name:       "JobH",
			department: "DeptA",
		},
	}

	jobMinHeap := heap.NewMinHeapCustomComp(jobs) // HeapCustomComp[job]
	jobMaxHeap := heap.NewMaxHeapCustomComp(jobs) // HeapCustomComp[job]

	fj := job{
		priority:   5,
		name:       "JobJ",
		department: "DeptX",
	}
	jobMinHeap.Insert(fj)
	jobMaxHeap.Insert(fj)

	for jobMinHeap.Size() != 0 {
		fmt.Printf("Current job %v\n", jobMinHeap.Extract())
	}

	for jobMaxHeap.Size() != 0 {
		fmt.Printf("Current job %v\n", jobMaxHeap.Extract())
	}

	// Output:
	// Current job {0 JobZ DeptC}
	// Current job {1 JobB DeptA}
	// Current job {4 JobA DeptA}
	// Current job {5 JobJ DeptX}
	// Current job {7 JobH DeptA}
	// Current job {7 JobH DeptA}
	// Current job {5 JobJ DeptX}
	// Current job {4 JobA DeptA}
	// Current job {1 JobB DeptA}
	// Current job {0 JobZ DeptC}
}
