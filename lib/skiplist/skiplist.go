package skiplist

import (
	"container/list"
	"math"
	"math/rand"
	"time"
)

type item struct {
	val  int
	down *list.Element
	up   *list.Element
	ls   *list.List
}

type SkipList struct {
	levels []*list.List
}

func New() SkipList {
	sl := SkipList{}
	sl.levels = append(sl.levels, list.New())
	sl.levels[0].PushBack(&item{math.MinInt32, nil, nil, sl.levels[0]})
	rand.Seed(time.Now().UnixNano())
	return sl
}

func (sl *SkipList) Search(target int) bool {
	el, _ := sl.searchHelper(target)
	return el.Value.(*item).val == target
}

func (sl *SkipList) Add(num int) {
	el, path := sl.searchHelper(num)
	ls := el.Value.(*item).ls
	el = ls.InsertAfter(&item{num, nil, nil, ls}, el)
	sl.promote(el, path)
}

func (sl *SkipList) Erase(num int) bool {
	el, _ := sl.searchHelper(num)
	if el.Value.(*item).val != num {
		return false
	}

	for el != nil {
		up := el.Value.(*item).up
		el.Value.(*item).ls.Remove(el)
		el = up
	}
	return true
}

func (sl *SkipList) searchHelper(target int) (*list.Element, []*list.Element) {
	cur := sl.levels[len(sl.levels)-1].Front()
	var path []*list.Element
	for {
		if cur.Next() != nil {
			if cur.Next().Value.(*item).val >= target {
				if cur.Value.(*item).down != nil {
					path = append(path, cur)
					cur = cur.Value.(*item).down
				} else {
					if cur.Next().Value.(*item).val == target {
						cur = cur.Next()
					}
					break
				}
			} else {
				cur = cur.Next()
			}
		} else {
			if cur.Value.(*item).down != nil {
				path = append(path, cur)
				cur = cur.Value.(*item).down
			} else {
				break
			}
		}
	}

	return cur, path
}

func (sl *SkipList) promote(el *list.Element, path []*list.Element) {
	num := el.Value.(*item).val
	pathIndex := len(path) - 1
	for rand.Uint32()%2 == 1 {
		it := el.Value.(*item)
		var appendTo *list.Element
		if pathIndex < 0 {
			nl := list.New()
			sl.levels = append(sl.levels, nl)
			appendTo = nl.PushBack(&item{math.MinInt32, it.ls.Front(), nil, nl})
			it.ls.Front().Value.(*item).up = appendTo
		} else {
			appendTo = path[pathIndex]
			pathIndex--
		}
		upls := appendTo.Value.(*item).ls
		it.up = upls.InsertAfter(&item{num, el, nil, upls}, appendTo)
		el = it.up
	}
}
