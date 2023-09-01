package skiplist

import (
	"container/list"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

type item[T constraints.Ordered] struct {
	val  T
	down *list.Element
	up   *list.Element
	ls   *list.List
}

type skipListImpl[T constraints.Ordered] struct {
	levels         []*list.List
	minVal, maxVal T
}

func newSkipListImpl[T constraints.Ordered](minVal, maxVal T) *skipListImpl[T] {
	sl := skipListImpl[T]{levels: []*list.List{list.New()}, minVal: minVal, maxVal: maxVal}
	sl.levels[0].PushBack(&item[T]{minVal, nil, nil, sl.levels[0]})
	rand.Seed(time.Now().UnixNano())
	return &sl
}

func (sl *skipListImpl[T]) Search(target T) bool {
	el, _ := sl.searchHelper(target)
	return el.Value.(*item[T]).val == target
}

func (sl *skipListImpl[T]) Add(num T) {
	el, path := sl.searchHelper(num)
	ls := el.Value.(*item[T]).ls
	el = ls.InsertAfter(&item[T]{num, nil, nil, ls}, el)
	sl.promote(el, path)
}

func (sl *skipListImpl[T]) Remove(num T) bool {
	el, _ := sl.searchHelper(num)
	if el.Value.(*item[T]).val != num {
		return false
	}

	for el != nil {
		up := el.Value.(*item[T]).up
		el.Value.(*item[T]).ls.Remove(el)
		el = up
	}
	return true
}

func (sl *skipListImpl[T]) ToSlice() []T {
	res := []T{}
	level0 := sl.levels[0]
	for e := level0.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(*item[T]).val)
	}

	return res[1:]
}

func (sl *skipListImpl[T]) searchHelper(target T) (*list.Element, []*list.Element) {
	cur := sl.levels[len(sl.levels)-1].Front()
	var path []*list.Element
	for {
		if cur.Next() != nil {
			if cur.Next().Value.(*item[T]).val >= target {
				if cur.Value.(*item[T]).down != nil {
					path = append(path, cur)
					cur = cur.Value.(*item[T]).down
				} else {
					if cur.Next().Value.(*item[T]).val == target {
						cur = cur.Next()
					}
					break
				}
			} else {
				cur = cur.Next()
			}
		} else {
			if cur.Value.(*item[T]).down != nil {
				path = append(path, cur)
				cur = cur.Value.(*item[T]).down
			} else {
				break
			}
		}
	}

	return cur, path
}

func (sl *skipListImpl[T]) promote(el *list.Element, path []*list.Element) {
	num := el.Value.(*item[T]).val
	pathIndex := len(path) - 1
	for rand.Uint32()%2 == 1 {
		it := el.Value.(*item[T])
		var appendTo *list.Element
		if pathIndex < 0 {
			nl := list.New()
			sl.levels = append(sl.levels, nl)
			appendTo = nl.PushBack(&item[T]{sl.maxVal, it.ls.Front(), nil, nl})
			it.ls.Front().Value.(*item[T]).up = appendTo
		} else {
			appendTo = path[pathIndex]
			pathIndex--
		}
		upls := appendTo.Value.(*item[T]).ls
		it.up = upls.InsertAfter(&item[T]{num, el, nil, upls}, appendTo)
		el = it.up
	}
}
