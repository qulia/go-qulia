package graph_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib/graph"
	"github.com/stretchr/testify/assert"
)

func TestGraphBasic(t *testing.T) {
	testGraph := graph.NewGraph()
	testGraph.Add(graph.NewNode("node1", nil), graph.NewNode("node2", nil))
	testGraph.Add(graph.NewNode("node3", nil), nil)
	testGraph.AddBidirectional(graph.NewNode("node4", nil), graph.NewNode("node5", nil))

	// Check expected edges
	assert.NotNil(t, testGraph.Nodes["node1"].EdgesOut["node2"])
	assert.NotNil(t, testGraph.Nodes["node2"].EdgesIn["node1"])
	assert.Empty(t, testGraph.Nodes["node2"].EdgesOut)
	assert.Empty(t, testGraph.Nodes["node3"].EdgesOut)
	assert.NotNil(t, testGraph.Nodes["node4"].EdgesOut["node5"])
	assert.NotNil(t, testGraph.Nodes["node4"].EdgesIn["node5"])
	assert.NotNil(t, testGraph.Nodes["node5"].EdgesOut["node4"])
	assert.NotNil(t, testGraph.Nodes["node5"].EdgesIn["node4"])

	if generateViz {
		// Generage graph
		dotToImageGraphviz("TestGraphBasic", "svg", []byte(testGraph.Dot()))
	}
}
