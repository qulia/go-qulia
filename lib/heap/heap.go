package heap

type Interface interface {
	Insert(interface{})
	Extract()interface{}
	Length() int
	IsEmpty() bool
}

type CompareFunc func(first, second interface{}) int

type MinHeap struct {
}

func NewMinHeap(input []interface{}, compareToFunc CompareFunc) Interface {
	buffer := make([]interface{}, len(input))
	copy(buffer, input)
	return initHeap(buffer, compareToFunc, false)
}

type MaxHeap struct {
}

func NewMaxHeap(input []interface{}, compareToFunc CompareFunc) Interface{
	buffer := make([]interface{}, len(input))
	copy(buffer, input)
	return initHeap(buffer, compareToFunc, true)
}

type heap struct {
	maxOnTop bool
	buffer []interface{}
	compareFunc CompareFunc
}

func initHeap(buffer []interface{}, compareToFunc CompareFunc, maxOnTop bool) Interface {
	h := heap{buffer:buffer, compareFunc:compareToFunc, maxOnTop:maxOnTop}
	h.heapify()
	return &h
}

func (h *heap) Insert(elem interface{}) {
	// Insert at the end, swift up
	h.buffer = append(h.buffer, elem)
	h.swiftUp(h.Length() - 1)
}

func (h *heap) Extract() interface{} {
	if h.IsEmpty() {
		return nil
	}

	// Capture first, swap with last, shrink 1, swift down from top
	first := h.buffer[0]
	h.swap(0, h.Length() - 1)
	h.buffer = h.buffer[:h.Length() - 1]
	h.swiftDown(0)

	return first
}

func (h *heap) IsEmpty() bool {
	return h.Length() == 0
}

func (h * heap) Length() int {
	return len(h.buffer)
}

func (h * heap) swiftUp(index int) {
	// If we are already at the root, nothing to do
	if index == 0 {
		return
	}

	current := index
	parent := current / 2

	top, equal := h.findTop(current, parent)
	if equal {
		return
	}

	if top != parent {
		h.swap(top, parent)
		h.swiftUp(parent)
	}
}

func (h * heap) swiftDown(index int) {
	// If at the leaf, done
	if index >= h.Length() / 2 {
		return
	}

	parent := index
	left := 2 * index + 1
	right := 2 * index + 2

	top, equal := h.findTop(parent, left)
	if equal {
		top, equal = h.findTop(parent, right)
		if equal {
			return
		}
	} else {
		top, _ = h.findTop(top, right)
	}

	if top != parent {
		h.swap(top, parent)
		h.swiftDown(top)
	}
}

func (h *heap) findTop(first int, second int) (int, bool) {
	if first >= h.Length() {
		return second, false
	}

	if second >= h.Length() {
		return first, false
	}

	var top int
	multiplier := 1
	if !h.maxOnTop {
		multiplier = -1
	}

	comp := h.compareFunc(h.buffer[first], h.buffer[second]) * multiplier
	if comp > 0 {
		top = first
	} else if comp < 0 {
		top = second
	} else {
		// Equal
		return -1, true
	}
	return top, false
}

func (h * heap) swap(i,j int) {
	tmp := h.buffer[i]
	h.buffer[i] = h.buffer[j]
	h.buffer[j] = tmp
}

func (h * heap) heapify() {
	if h.Length() <= 1 {
		return
	}

	// leaf nodes are already heaps
	// Start at first non-leaf node and go up to the root swifting up as needed
	// Leaf nodes start at n/2 goes to n-1
	for i := h.Length() / 2 - 1; i >= 0; i-- {
		h.swiftDown(i)
	}
}
