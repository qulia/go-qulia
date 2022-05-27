package unionfind

// Ref: https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type UnionFind[T comparable] interface {
	// Union two clusters
	Union(T, T)

	// Find the cluster of value, if it does not belong to any
	// create one
	Find(T) T

	// Cluster size
	Size(T) int

	// Number of clusters
	Count() int
}

func NewUnionFind[T comparable](input []T) UnionFind[T] {
	return newUniondFindImpl(input)
}
