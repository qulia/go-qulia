package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {
	sl := New()

	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	sl.print("1,2,3 added")
	assert.False(t, sl.Search(0))
	sl.Add(4)
	assert.True(t, sl.Search(1))
	assert.False(t, sl.Erase(0))
	assert.True(t, sl.Erase(1))
	sl.print("1 erase")
	assert.False(t, sl.Search(1))
	sl.Add(0)
	assert.True(t, sl.Search(0))
	sl.Add(1)
	sl.print("0,1 added")
	assert.True(t, sl.Search(1))
	sl.Erase(2)
	assert.False(t, sl.Search(2))
	sl.Erase(4)
	assert.False(t, sl.Search(4))
	sl.print("2,4 erased")
}

func TestSkipList100(t *testing.T) {
	sl := New()
	for i := 100; i >= 0; i-- {
		sl.Add(i)
		sl.Add(i)
	}
	sl.print("0-100 added")
	for i := 0; i <= 100; i++ {
		sl.Erase(i)
		sl.Erase(i)
	}
	sl.print("0-100 erased")
}

func TestSkipListRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sl := New()
	added := []int{}
	n := 100
	for i := n - 1; i >= 0; i-- {
		val := rand.Int()
		sl.Add(val)
		added = append(added, val)
	}
	sl.print("0-n added")
	for _, val := range added {
		assert.True(t, sl.Search(val))
	}
	for _, val := range added {
		assert.True(t, sl.Erase(val))
	}
	sl.print("0-n erased")
	for _, val := range added {
		assert.False(t, sl.Search(val))
	}
	for _, val := range added {
		sl.Add(val)
	}
	sl.print("0-n added")
	for _, val := range added {
		assert.True(t, sl.Search(val))
	}
	for _, val := range added {
		assert.True(t, sl.Erase(val))
	}
	sl.print("0-n erased")
}

func (sl *SkipList) print(checkpoint string) {
	fmt.Printf("%s\n", checkpoint)
	for i := len(sl.levels) - 1; i >= 0; i-- {
		ls := sl.levels[i]
		fmt.Printf("%d: ", i)
		for el := ls.Front(); el != nil; el = el.Next() {
			fmt.Printf("%d[", el.Value.(*item).val)
			if el.Value.(*item).down != nil {
				fmt.Printf("down:%d ", el.Value.(*item).down.Value.(*item).val)
			} else {
				fmt.Printf("down:nil ")
			}

			if el.Value.(*item).up != nil {
				fmt.Printf("up:%d", el.Value.(*item).up.Value.(*item).val)
			} else {
				fmt.Printf("up:nil")
			}
			fmt.Printf("] ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
