package graph_test

import (
	"fmt"
	"sort"

	"github.com/qulia/go-qulia/lib/graph"
)

const (
	availableCount = "availableCount"
)

// This is a problem where we need to find the itinerary using the tickets
// e.g. tickets [A, B] [B, C] [B,D] [C,B] itinerary A->B->C->B->D using all tickets
// Note that if we had picked A->B->D we would not be able to use all tickets, so not the right path
// Since we need to keep track of all edges we visited, metadata support on the edge can be used
// There could be multiple tickets available for the same source, dest pair. So we can keep track of
// "available edges" in the metadata
func ExampleGraph_itinerary() {
	tickets := [][]string{{"A", "B"}, {"B", "C"}, {"B", "D"}, {"C", "B"}}
	fmt.Printf("%v\n", testWithTickets(tickets))

	tickets = [][]string{{"A", "B"}, {"A", "C"}, {"C", "A"}}
	fmt.Printf("%v\n", testWithTickets(tickets))

	// Note there are two tickets D,B so availableCount is initially 2 for this edge
	tickets = [][]string{{"E", "F"}, {"D", "B"}, {"B", "C"}, {"C", "B"}, {"B", "E"}, {"D", "B"}, {"F", "D"}, {"D", "C"}, {"B", "D"}, {"C", "D"}}
	fmt.Printf("%v", testWithTickets(tickets))
	//Output:
	//[A B C B D]
	//[A C A B]
	//[B C B D B E F D C D B]
}

func testWithTickets(tickets [][]string) []string {
	// Create graph with source, dest in tickets
	cGraph := graph.NewGraph()
	for _, ticket := range tickets {
		cGraph.Add(createOrGetNewNode(cGraph, ticket[0]), createOrGetNewNode(cGraph, ticket[1]))
		m := cGraph.Nodes[ticket[0]].EdgesOut[ticket[1]].Metadata
		if _, ok := m[availableCount]; !ok {
			m[availableCount] = 0
		}
		m[availableCount] = m[availableCount].(int) + 1
	}
	//Find itinerary using all edges
	numEdges := len(tickets)
	var itinerary []string

	// sort nodes to get lexical order in result
	var nodeNames []string
	for key := range cGraph.Nodes {
		nodeNames = append(nodeNames, key)
	}

	sort.Strings(nodeNames)
	for _, nodeName := range nodeNames {
		node := cGraph.Nodes[nodeName]
		itinerary = append(itinerary, node.Name)
		if visitAll(node, numEdges, &itinerary) {
			break
		} else {
			itinerary = itinerary[:len(itinerary)-1]
		}
	}

	return itinerary
}

func visitAll(node *graph.Node, numEdges int, itinerary *[]string) bool {
	if len(*itinerary) == numEdges+1 {
		return true
	}

	// sort targets to get lexical order in result
	var edgeNames []string
	for key := range node.EdgesOut {
		edgeNames = append(edgeNames, key)
	}

	sort.Strings(edgeNames)
	for _, edgeName := range edgeNames {
		edge := node.EdgesOut[edgeName]
		m := edge.Metadata
		if m[availableCount].(int) > 0 {
			m[availableCount] = m[availableCount].(int) - 1
			target := edge.Target
			*itinerary = append(*itinerary, target.Name)
			if visitAll(target, numEdges, itinerary) {
				return true
			} else {
				// backtrack so it can be visited again
				m[availableCount] = m[availableCount].(int) + 1
				*itinerary = (*itinerary)[:len(*itinerary)-1]
			}
		}
	}

	return false
}

func createOrGetNewNode(cGraph graph.Interface, name string) *graph.Node {
	if node, ok := cGraph.GetNodes()[name]; ok {
		return node
	}
	return graph.NewNode(name, nil)
}
