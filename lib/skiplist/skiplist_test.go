package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList(math.MinInt32, math.MaxInt32)

	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	debugPrint(sl, "1,2,3 added")
	assert.False(t, sl.Search(0))
	sl.Add(4)
	assert.True(t, sl.Search(1))
	assert.False(t, sl.Remove(0))
	assert.True(t, sl.Remove(1))
	debugPrint(sl, "1 erase")
	assert.False(t, sl.Search(1))
	sl.Add(0)
	assert.True(t, sl.Search(0))
	sl.Add(1)
	debugPrint(sl, "0,1 added")
	assert.True(t, sl.Search(1))
	sl.Remove(2)
	assert.False(t, sl.Search(2))
	sl.Remove(4)
	assert.False(t, sl.Search(4))
	debugPrint(sl, "2,4 erased")
}

func TestSkipList100(t *testing.T) {
	sl := NewSkipList(math.MinInt32, math.MaxInt32)
	for i := 100; i >= 0; i-- {
		sl.Add(i)
		sl.Add(i)
	}
	debugPrint(sl, "0-100 added")
	for i := 0; i <= 100; i++ {
		sl.Remove(i)
		sl.Remove(i)
	}
	debugPrint(sl, "0-100 erased")
}

func TestSkipListRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sl := NewSkipList(math.MinInt32, math.MaxInt32)
	added := []int{}
	n := 100
	for i := n - 1; i >= 0; i-- {
		val := rand.Int()
		sl.Add(val)
		added = append(added, val)
	}
	debugPrint(sl, "0-n added")
	for _, val := range added {
		assert.True(t, sl.Search(val))
	}
	for _, val := range added {
		assert.True(t, sl.Remove(val))
	}
	debugPrint(sl, "0-n erased")
	for _, val := range added {
		assert.False(t, sl.Search(val))
	}
	for _, val := range added {
		sl.Add(val)
	}
	debugPrint(sl, "0-n added")
	for _, val := range added {
		assert.True(t, sl.Search(val))
	}
	for _, val := range added {
		assert.True(t, sl.Remove(val))
	}
	debugPrint(sl, "0-n erased")
}

func debugPrint(sli SkipList[int], checkpoint string) {
	sl := sli.(*skipListImpl[int])
	fmt.Printf("%s\n", checkpoint)
	for i := len(sl.levels) - 1; i >= 0; i-- {
		ls := sl.levels[i]
		fmt.Printf("%d: ", i)
		for el := ls.Front(); el != nil; el = el.Next() {
			fmt.Printf("%d[", el.Value.(*item[int]).val)
			if el.Value.(*item[int]).down != nil {
				fmt.Printf("down:%d ", el.Value.(*item[int]).down.Value.(*item[int]).val)
			} else {
				fmt.Printf("down:nil ")
			}

			if el.Value.(*item[int]).up != nil {
				fmt.Printf("up:%d", el.Value.(*item[int]).up.Value.(*item[int]).val)
			} else {
				fmt.Printf("up:nil")
			}
			fmt.Printf("] ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
