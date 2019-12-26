package heap_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"testing/quick"
	"time"

	"github.com/qulia/go-qulia/lib/heap"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var sliceRand *rand.Rand
func init() {
	sliceRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func BenchmarkHeapBasic(b *testing.B) {
	randSlice, ok := quick.Value(reflect.ValueOf([]int{}).Type(), sliceRand)
	assert.True(b, ok)
	genInput := randSlice.Interface().([]int)
	var genInputInterface []interface{}
	for _,elem := range genInput {
		genInputInterface = append(genInputInterface, elem)
	}
	log.Infof("Size %d", len(genInputInterface))
	//log.Info("Generated input %s", genInput)
	b.ResetTimer()
	b.Run("Create heap from slice", func(b *testing.B) {
		h := heap.NewMaxHeap(genInputInterface, intCompFunc)
		b.StopTimer()
		checkHeap(b, genInputInterface, h)
		b.StartTimer()
	})
}

func checkHeap(b *testing.B, genInput []interface{}, h heap.Interface) {
	buffer := make([]interface{}, len(genInput))
	copy(buffer, genInput)
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].(int) > buffer[j].(int)
	})

	assert.Equal(b, len(buffer), h.Length())
	index := 0
	for h.Length() > 0{
		assert.Equal(b, buffer[index], h.Extract().(int))
		index++
	}
}