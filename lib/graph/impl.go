package graph

import "github.com/qulia/go-qulia/lib/set"

type graphImpl[T comparable] struct {
	nodes   map[T]map[T]bool
	nodeSet set.Set[T]
}

func newGraphImpl[T comparable]() *graphImpl[T] {
	return &graphImpl[T]{map[T]map[T]bool{}, set.NewSet[T]()}
}

func (g *graphImpl[T]) AddNode(node T) {
	g.nodeSet.Add(node)
}

func (g *graphImpl[T]) Add(source, target T) {
	if g.nodes[source] == nil {
		g.nodes[source] = map[T]bool{}
	}
	g.nodes[source][target] = true
	g.AddNode(source)
	g.AddNode(target)
}

func (g *graphImpl[T]) AddBidirectional(node1, node2 T) {
	g.Add(node1, node2)
	g.Add(node2, node1)
}

func (g *graphImpl[T]) GetNodes() set.Set[T] {
	return g.nodeSet
}

func (g *graphImpl[T]) Adjacencies(node T) map[T]bool {
	return g.nodes[node]
}
