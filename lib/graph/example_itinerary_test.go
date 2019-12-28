package graph_test

import (
	"fmt"

	"github.com/qulia/go-qulia/lib/graph"
)

// This is a common graph problem where we need to find the itinerary using the tickets
// e.g. tickets [A, B] [B, C] [B,D] [C,B] itinerary A->B->C->B->D using all tickets
// Note that if we had picked A->B->D we would not be able to use all tickets, so not the right path
// Since we need to keep track of all edges we visited, metadata support on the edge can be used
func ExampleGraph_visitAllEdges() {
	tickets := [][]string{{"A", "B"}, {"B", "C"}, {"B", "D"}, {"C", "B"}}

	// Create graph with source, dest in tickets
	cGraph := graph.NewGraph()
	for _, ticket := range tickets {
		cGraph.Add(createOrGetNewNode(cGraph, ticket[0]), createOrGetNewNode(cGraph, ticket[1]))
	}
	//Find itinerary using all edges
	numEdges := len(tickets)
	var itinerary []string
	for _, node := range cGraph.Nodes {
		itinerary = append(itinerary, node.Name)
		if visitAll(node, numEdges, &itinerary) {
			break
		} else {
			itinerary = itinerary[:len(itinerary)-1]
		}
	}

	fmt.Printf("%v", itinerary)
	// Output:
	// [A B C B D]
}

func visitAll(node *graph.Node, numEdges int, itinerary *[]string) bool {
	if len(*itinerary) == numEdges+1 {
		return true
	}

	for _, edge := range node.EdgesOut {
		if visited, ok := edge.Metadata[isVisited]; !ok || !visited.(bool) {
			edge.Metadata[isVisited] = true
			target := edge.Target
			*itinerary = append(*itinerary, target.Name)
			if visitAll(target, numEdges, itinerary) {
				return true
			} else {
				// backtrack so it can be visited again
				edge.Metadata[isVisited] = false
				*itinerary = (*itinerary)[:len(*itinerary)-1]
				return false
			}
		}
	}

	return false
}

func createOrGetNewNode(cGraph *graph.Graph, name string) *graph.Node {
	if node, ok := cGraph.Nodes[name]; ok {
		return node
	}
	return graph.NewNode(name, nil)
}
