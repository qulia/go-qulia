package heap_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/qulia/go-qulia/lib/heap"
	"github.com/stretchr/testify/assert"

	contheap "container/heap"
)

var sliceRand *rand.Rand

const (
	numsDefaultMin  = -100000
	numsDefaultMax  = 100000
	numsDefaultSize = 10000
)

func init() {
	sliceRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

/* Copied from https://golang.org/pkg/container/heap/
for benchmark comparison
*/
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type JobHeap []job

func (h JobHeap) Len() int           { return len(h) }
func (h JobHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h JobHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *JobHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(job))
}

func (h *JobHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkHeapBasic(b *testing.B) {
	genInput := generateInput(numsDefaultSize, numsDefaultMin, numsDefaultMax)
	fmt.Printf("Size %d\n", len(genInput))
	//log.Info("Generated input %s", genInput)
	b.ResetTimer()
	b.Run("Create heap from slice", func(b *testing.B) {
		h := heap.NewMaxHeap(genInput)
		b.StopTimer()
		checkHeap(b, genInput, h)
		b.StartTimer()
	})
}

func BenchmarkHeapPush(b *testing.B) {
	genInput := generateInput(numsDefaultSize, numsDefaultMin, numsDefaultMax)

	fmt.Printf("Size %d\n", len(genInput))
	//log.Info("Generated input %s", genInput)
	b.ResetTimer()
	b.Run("Create heap from slice", func(b *testing.B) {
		h := heap.NewMaxHeap[int](nil)
		for _, elem := range genInput {
			h.Insert(elem)
		}
		b.StopTimer()
		checkHeap(b, genInput, h)
		b.StartTimer()
	})
}

func BenchmarkHeapCompareStdContainerHeap(b *testing.B) {
	genInput := generateInput(numsDefaultSize, numsDefaultMin, numsDefaultMax)

	fmt.Printf("Size %d\n", len(genInput))

	b.ResetTimer()
	b.Run("container/heap", func(b *testing.B) {
		stdh := &IntHeap{}
		contheap.Init(stdh)
		for _, elem := range genInput {
			contheap.Push(stdh, elem)
		}

		for stdh.Len() != 0 {
			contheap.Pop(stdh)
		}
	})
	b.StopTimer()

	b.ResetTimer()
	b.Run("go-qulia/lib/heap", func(b *testing.B) {
		h := heap.NewMaxHeap[int](nil)
		for _, elem := range genInput {
			h.Insert(elem)
		}

		for h.Size() != 0 {
			h.Extract()
		}
	})
	b.StopTimer()
	/*
		2022/06/11 17:06:55 Size 10000
		goos: darwin
		goarch: amd64
		pkg: github.com/qulia/go-qulia/lib/heap
		cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
		BenchmarkHeapCompareStdContainerHeap/container/heap-8         	1000000000	         0.002324 ns/op	       0 B/op	       0 allocs/op
		BenchmarkHeapCompareStdContainerHeap/go-qulia/lib/heap-8      	1000000000	         0.002898 ns/op	       0 B/op	       0 allocs/op
		PASS
	*/

}

func BenchmarkHeapCompareStdContainerHeapJob(b *testing.B) {
	b.ResetTimer()
	b.Run("container/heap", func(b *testing.B) {
		stdhj := &JobHeap{}
		contheap.Init(stdhj)
		for i := 0; i < b.N; i++ {
			contheap.Push(stdhj, job{priority: rand.Int()})
		}

		for stdhj.Len() != 0 {
			contheap.Pop(stdhj)
		}
	})
	b.StopTimer()

	b.ResetTimer()
	b.Run("go-qulia/lib/heap", func(b *testing.B) {
		h := heap.NewMaxHeapFlex[job](nil)
		for i := 0; i < b.N; i++ {
			h.Insert(job{priority: rand.Int()})
		}

		for h.Size() != 0 {
			h.Extract()
		}
	})
	b.StopTimer()
	/*
	   goos: darwin
	   goarch: amd64
	   pkg: github.com/qulia/go-qulia/lib/heap
	   cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
	   BenchmarkHeapCompareStdContainerHeapJob/container/heap-8         	 1353691	       901.5 ns/op	     295 B/op	       2 allocs/op
	   BenchmarkHeapCompareStdContainerHeapJob/go-qulia/lib/heap-8      	 1000000	      1245 ns/op	     215 B/op	       0 allocs/op
	   PASS
	*/
}

func generateInput(size, min, max int) []int {
	var result []int
	for i := 0; i < size; i++ {
		val := sliceRand.Intn(max-min) + min
		result = append(result, val)
	}

	return result
}

func checkHeap(b *testing.B, genInput []int, h heap.Heap[int]) {
	buffer := make([]int, len(genInput))
	copy(buffer, genInput)
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i] > buffer[j]
	})

	assert.Equal(b, len(buffer), h.Size())
	index := 0
	for h.Size() > 0 {
		assert.Equal(b, buffer[index], h.Extract())
		index++
	}
}
