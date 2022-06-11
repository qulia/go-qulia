package graph_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib/graph"
	"github.com/qulia/go-qulia/lib/set"
	"github.com/stretchr/testify/assert"
)

func TestGraphBasic(t *testing.T) {
	testGraph := graph.NewGraph[string]()
	testGraph.Add("node1", "node2")
	testGraph.AddNode("node3")
	testGraph.AddBidirectional("node4", "node5")
	testGraph.AddBidirectional("node3", "node4")

	expected := set.NewSet[string]()
	expected.FromSlice([]string{"node1", "node2", "node3", "node4", "node5"})

	assert.Equal(t, expected, testGraph.GetNodes())
	assert.True(t, testGraph.Adjacencies("node3")["node4"])
	assert.True(t, testGraph.Adjacencies("node4")["node3"])

	// Generage graph if dot exists
	_, _ = dotToImageGraphviz("TestGraphBasic", "svg", []byte(graphToDot(testGraph)))
}
