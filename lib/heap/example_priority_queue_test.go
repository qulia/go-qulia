// This example demonstrates a job priority queue built using the heap interface.
package heap_test

import (
	"fmt"

	"github.com/qulia/go-qulia/v2/lib/heap"
)

// This example initializes the max heap with slice of ints
// With < operator as comparison the values are extracted in descending order
// Any type met by constraints.Ordered can be used for the content of the heap
func ExampleHeap() {
	iHeap := heap.NewMaxHeap([]int{3, 7, 4, 4, 1}) // Heap[int]
	iHeap.Insert(9)

	// Calculate the sum of (rank * val) rank: 1-n
	sum := 0
	for rank := 1; !iHeap.IsEmpty(); rank++ {
		cur := iHeap.Extract()
		fmt.Printf("Out: %d\n", cur)
		sum += cur * rank
	}

	fmt.Printf("Sum: %d\n", sum)

	// Output:
	// Out: 9
	// Out: 7
	// Out: 4
	// Out: 4
	// Out: 3
	// Out: 1
	// Sum: 72
}

// This example initializes the heap with list of jobs and pushes another one with Insert method
// With the provided comparison method Less on the type implementing lib.Comparer[T]
// depending on the heap type (min/max) the jobs will be extracted in order
func ExampleHeapFlex() {
	jobs := []job{
		{priority: 4, name: "JobA", department: "DeptA"},
		{priority: 1, name: "JobB", department: "DeptA"},
		{priority: 0, name: "JobZ", department: "DeptC"},
		{priority: 7, name: "JobH", department: "DeptA"},
	}

	jobMinHeap := heap.NewMinHeapFlex(jobs) // HeapFlex[job]
	jobMaxHeap := heap.NewMaxHeapFlex(jobs) // HeapFlex[job]

	fj := job{priority: 5, name: "JobJ", department: "DeptX"}
	jobMinHeap.Insert(fj)
	jobMaxHeap.Insert(fj)

	for jobMinHeap.Size() != 0 {
		j := jobMinHeap.Extract()
		fmt.Printf("Current job (pri, name) (%v, %v)\n", j.priority, j.name)
	}

	for jobMaxHeap.Size() != 0 {
		fmt.Printf("Current job %v\n", jobMaxHeap.Extract())
	}

	// Output:
	// Current job (pri, name) (0, JobZ)
	// Current job (pri, name) (1, JobB)
	// Current job (pri, name) (4, JobA)
	// Current job (pri, name) (5, JobJ)
	// Current job (pri, name) (7, JobH)
	// Current job {7 JobH DeptA}
	// Current job {5 JobJ DeptX}
	// Current job {4 JobA DeptA}
	// Current job {1 JobB DeptA}
	// Current job {0 JobZ DeptC}
}

type job struct {
	priority   int
	name       string
	department string
}

func (j job) Compare(other job) int {
	return j.priority - other.priority
}
