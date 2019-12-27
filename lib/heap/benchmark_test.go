package heap_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/qulia/go-qulia/lib/heap"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var sliceRand *rand.Rand
const (
	numsDefaultMin = -100000
	numsDefaultMax = 100000
	numsDefaultSize = 10000
)
func init() {
	sliceRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func BenchmarkHeapBasic(b *testing.B) {
	genInput := generateInput(numsDefaultSize, numsDefaultMin, numsDefaultMax)
	var genInputInterface []interface{}
	for _, elem := range genInput {
		genInputInterface = append(genInputInterface, elem)
	}
	log.Infof("Size %d", len(genInputInterface))
	//log.Info("Generated input %s", genInput)
	b.ResetTimer()
	b.Run("Create heap from slice", func(b *testing.B) {
		h := heap.NewMaxHeap(genInputInterface, intCompFunc)
		b.StopTimer()
		checkHeap(b, genInput, h)
		b.StartTimer()
	})
}

func BenchmarkHeapPush(b *testing.B) {
	genInput := generateInput(numsDefaultSize, numsDefaultMin, numsDefaultMax)

	log.Infof("Size %d", len(genInput))
	//log.Info("Generated input %s", genInput)
	b.ResetTimer()
	b.Run("Create heap from slice", func(b *testing.B) {
		h := heap.NewMaxHeap(nil, intCompFunc)
		for _,elem := range genInput {
			h.Insert(elem)
		}
		b.StopTimer()
		checkHeap(b, genInput, h)
		b.StartTimer()
	})
}

func generateInput(size, min, max int) []int {
	var result []int
	for i := 0; i < size; i++ {
		val := sliceRand.Intn(max - min) + min
		result = append(result, val)
	}

	return result
}

func checkHeap(b *testing.B, genInput []int, h heap.Interface) {
	buffer := make([]int, len(genInput))
	copy(buffer, genInput)
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i] > buffer[j]
	})

	assert.Equal(b, len(buffer), h.Size())
	index := 0
	for h.Size() > 0{
		assert.Equal(b, buffer[index], h.Extract().(int))
		index++
	}
}