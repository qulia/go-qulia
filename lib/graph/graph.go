package graph

// Metadata to append properties,tags to Graph, Node, Edge
type Metadata map[string]interface{}

// Graph with Nodes and metadata
type Graph struct {
	Nodes map[string]*Node
	MData *Metadata
}

// Node with name, object holding the data, connections in and out, and metadata
type Node struct {
	Name string

	// The data associated with Node, e.g. person
	Data interface{}

	// Outgoing edges Target node.Name is the key, current node is the source
	EdgesOut map[string]*Edge

	// Source node.Name is the key, current node is the target
	EdgesIn map[string]*Edge

	// This is free form map to set values while traversing graph, e.g. "isVisited = true"
	// Normally would be used for more advanced scenarios
	MData *Metadata
}

// Edge from source to target with metadata
type Edge struct {
	Source *Node
	Target *Node
	Metadata *Metadata
}

func NewNode(name string, data interface{}) *Node{
	node := Node{
		Name:     name,
		Data:     data,
		MData:    &Metadata{},
		EdgesIn:  make(map[string]*Edge),
		EdgesOut: make(map[string]*Edge),
	}

	return &node
}

func NewNodeWithMetadata(name string, data interface{}, mData *Metadata) *Node{
	node := NewNode(name, data)
	node.MData = mData
	return node
}

func NewGraph() *Graph {
	g := Graph{}
	g.Nodes = make(map[string]*Node)
	return &g
}

func NewGraphWithMetadata(mData *Metadata) *Graph {
	g := NewGraph()
	g.MData = mData
	return g
}

// Add relation to the graph, directional from source node to target node
func (g *Graph) Add(source, target *Node) {
	g.AddWithMetadata(source, target, &Metadata{})
}

// Add relation to the graph, directional from source node to target node with metadata
func (g *Graph) AddWithMetadata(source, target *Node, mData *Metadata) {
	g.Nodes[source.Name] = source

	if target != nil {
		g.Nodes[target.Name] = target
		g.addEdge(source, target, mData)
	}
}

// Add relation to the graph, bidirectional with node1 and node2, each direction with its own metadata
func (g *Graph) AddBidirectional(node1, node2 *Node) {
	g.AddBidirectionalWithMetadata(node1, node2, &Metadata{}, &Metadata{})
}

// Add relation to the graph, bidirectional with node1 and node2, each direction with its own metadata
func (g *Graph) AddBidirectionalWithMetadata(node1, node2 *Node, mData12, mData21 *Metadata) {
	g.Nodes[node1.Name] = node1
	g.Nodes[node2.Name] = node2
	g.addEdge(node1, node2, mData12)
	g.addEdge(node2, node1, mData21)
}

func (g *Graph) addEdge(from, to *Node, mData *Metadata) {
	if to != nil && from != nil {
		v := &Edge{
			Source: from,
			Target: to,
			Metadata:   mData,
		}

		g.Nodes[from.Name].EdgesOut[to.Name] = v
		g.Nodes[to.Name].EdgesIn[from.Name] = v
	}
}