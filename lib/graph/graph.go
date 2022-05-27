package graph

import (
	"github.com/qulia/go-qulia/lib/set"
)

type Graph[T comparable] interface {
	// GetNodes returns map of all nodes in the graph
	GetNodes() set.Set[T]

	// Add relation from source to target node
	Add(T, T)

	// Add relation from source to target node
	AddNode(T)

	// Add bidirectional relation between node1 and node2
	AddBidirectional(T, T)

	// Adjacent nodes from a source
	Adjacencies(T) map[T]bool // not using Set to allow range iteration
}

func NewGraph[T comparable]() Graph[T] {
	return newGraphImpl[T]()
}
