package graph

import "github.com/qulia/go-qulia/lib"

type Interface interface {
	// GetNodes returns map of all nodes in the graph
	GetNodes() map[string]*Node

	// Add relation from source to target node
	// if target is nil just add the node with no connection
	Add(source, target *Node)

	// Add bidirectional relation between node1 and node2, equivalent to calling
	// Add(node1, node2) Add(node2, node1)
	AddBidirectional(node1, node2 *Node)
}

// Graph with Nodes and metadata
type Graph struct {
	Nodes map[string]*Node
	MData lib.Metadata
}

// Node with name, object holding the data, connections in and out, and metadata
type Node struct {
	Name string

	// The data associated with Node, e.g. person
	Data interface{}

	// Outgoing edges Target node.Name is the key, current node is the source
	EdgesOut map[string]Edge

	// Source node.Name is the key, current node is the target
	EdgesIn map[string]Edge

	// This is free form map to set values while traversing graph, e.g. "isVisited = true"
	// Normally would be used for more advanced scenarios
	MData lib.Metadata
}

// Edge from source to target with metadata
type Edge struct {
	Source   *Node
	Target   *Node
	Metadata lib.Metadata
}

// NewNode creates a node later to be added to the graph
func NewNode(name string, data interface{}) *Node {
	node := Node{
		Name:     name,
		Data:     data,
		MData:    lib.Metadata{},
		EdgesIn:  make(map[string]Edge),
		EdgesOut: make(map[string]Edge),
	}

	return &node
}

// NewGraph initialized graph structure
func NewGraph() *Graph {
	g := Graph{}
	g.Nodes = make(map[string]*Node)
	g.MData = lib.Metadata{}
	return &g
}

// Add relation to the graph, directional from source node to target node
func (g *Graph) Add(source, target *Node) {
	g.Nodes[source.Name] = source

	if target != nil {
		g.Nodes[target.Name] = target
		g.addEdge(source, target)
	}
}

// Add relation to the graph, bidirectional with node1 and node2, each direction with its own metadata
func (g *Graph) AddBidirectional(node1, node2 *Node) {
	g.Nodes[node1.Name] = node1
	g.Nodes[node2.Name] = node2
	g.addEdge(node1, node2)
	g.addEdge(node2, node1)
}

// GetNodes
func (g *Graph) GetNodes() map[string]*Node {
	return g.Nodes
}

func (g *Graph) addEdge(from, to *Node) {
	if to != nil && from != nil {
		if _, ok := g.Nodes[from.Name].EdgesOut[to.Name]; ok {
			// Already has edge from source to dest
			return
		}

		v := Edge{
			Source:   from,
			Target:   to,
			Metadata: lib.Metadata{},
		}

		g.Nodes[from.Name].EdgesOut[to.Name] = v
		g.Nodes[to.Name].EdgesIn[from.Name] = v
	}
}
