package tree

type BinaryIndexTreeInterface interface {
	// add "val" at index "x"
	// 0 indexed bounded by the size of the index tree
	Update(x, val int)

	// returns the sum of elements up to and including index x
	Sum(x int) int
}

type BinaryIndexTree struct {
	buf []int
}

func (bit BinaryIndexTree) Update(x, val int) {
	x += 1
	for ; x > 0 && x < len(bit.buf); x += x & -x {
		bit.buf[x] += val
	}
}

func (bit BinaryIndexTree) Sum(x int) int {
	x += 1
	sum := 0
	for ; x > 0 && x < len(bit.buf); x -= x & -x {
		sum += bit.buf[x]
	}
	return sum
}

func NewBinaryIndexTree(size int) BinaryIndexTreeInterface {
	// bit buffer is 1 indexed
	bit := BinaryIndexTree{buf: make([]int, size+1)}
	return &bit
}
