package tree

type BinaryIndexTreeInterface interface {
	//add "val" at index "x"
	Update(x, val int)

	//returns the sum of first x elements
	Query(x int) int
}

type BinaryIndexTree struct {
	size int
	buf  []int
}

func (bit BinaryIndexTree) Update(x, val int) {
	for ; x <= bit.size; x += x & -x {
		bit.buf[x] += val
	}
}

func (bit BinaryIndexTree) Query(x int) int {
	sum := 0
	for ; x > 0; x -= x & -x {
		sum += bit.buf[x]
	}
	return sum
}

func NewBinaryIndexTree(size int) BinaryIndexTreeInterface {
	bit := BinaryIndexTree{size: size + 1, buf: make([]int, size+1)}
	return &bit
}
