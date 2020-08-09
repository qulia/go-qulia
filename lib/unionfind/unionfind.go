package unionfind

// Ref: https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type UnionFind struct {
	m map[int]*node
}

type node struct {
	val    int
	rank   int
	parent int
	size   int
}

func New(arr []int) *UnionFind {
	unif := &UnionFind{}
	unif.m = map[int]*node{}
	for _, val := range arr {
		unif.Find(val)
	}
	return unif
}

func (unif *UnionFind) Find(v int) int {
	n := unif.m[v]
	if n == nil {
		unif.m[v] = &node{v, 0, v, 1}
		n = unif.m[v]
	}
	if n.parent != n.val {
		// Not root
		// Path compression
		n.parent = unif.Find(n.parent)
	}
	return n.parent
}

func (unif *UnionFind) Union(x, y int) {
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
	}

	xRoot.size += yRoot.size
}

func (unif *UnionFind) Size(x int) int {
	node := unif.m[unif.Find(x)]
	if node == nil {
		return 0
	}

	return node.size
}
