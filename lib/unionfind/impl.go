package unionfind

import "fmt"

type unionFindImpl[T comparable] struct {
	m     map[T]*node[T]
	count int
}

type node[T comparable] struct {
	val, parent T
	rank        int
	size        int
}

func newUniondFindImpl[T comparable](input []T) *unionFindImpl[T] {
	unif := &unionFindImpl[T]{}
	unif.m = map[T]*node[T]{}
	for _, val := range input {
		unif.Find(val)
	}
	return unif
}

// Number of groups
func (unif *unionFindImpl[T]) Count() int {
	return unif.count
}

func (unif *unionFindImpl[T]) Find(v T) T {
	n := unif.m[v]
	if n == nil {
		unif.m[v] = &node[T]{v, v, 0, 1}
		n = unif.m[v]
		unif.count++
	}
	if n.parent != n.val {
		// Not root
		// Path compression
		n.parent = unif.Find(n.parent)
	}
	return n.parent
}

func (unif *unionFindImpl[T]) Union(x, y T) {
	rootX := unif.Find(x)
	rootY := unif.Find(y)

	if rootX == rootY {
		return
	}

	// Union by rank
	xRoot := unif.m[rootX]
	yRoot := unif.m[rootY]
	if xRoot.rank < yRoot.rank {
		xRoot, yRoot = yRoot, xRoot
	}

	yRoot.parent = xRoot.val
	if xRoot.rank == yRoot.rank {
		xRoot.rank += 1
		fmt.Printf("%v %v\n", xRoot, yRoot)
	}

	xRoot.size += yRoot.size
	unif.count--
}

func (unif *unionFindImpl[T]) Size(x T) int {
	node := unif.m[x]
	if node == nil {
		return 0
	}

	return unif.m[unif.Find(x)].size
}
