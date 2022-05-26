package graph

import "github.com/qulia/go-qulia/lib/set"

type graphComparable[T comparable] struct {
	nodes   map[T]map[T]bool
	nodeSet set.Set[T]
}

func (g *graphComparable[T]) AddNode(node T) {
	g.nodeSet.Add(node)
}

func (g *graphComparable[T]) Add(source, target T) {
	if g.nodes[source] == nil {
		g.nodes[source] = map[T]bool{}
	}
	g.nodes[source][target] = true
	g.AddNode(source)
	g.AddNode(target)
}

func (g *graphComparable[T]) AddBidirectional(node1, node2 T) {
	g.Add(node1, node2)
	g.Add(node2, node1)
}

func (g *graphComparable[T]) GetNodes() set.Set[T] {
	return g.nodeSet
}

func (g *graphComparable[T]) Adjacencies(node T) map[T]bool {
	return g.nodes[node]
}
