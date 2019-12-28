package graph_test

import (
	"fmt"
	"math"
	"sort"

	"github.com/qulia/go-qulia/lib/graph"
)

const (
	isVisited = "isVisited"
)

type city struct {
	name, country string
	population    int
}

// This example uses graph for finding connected components in a collection based on certain criteria
func ExampleGraph_connected_components() {
	// Given a collection of cities, find the connected groups that are in the same country and population 1000 apart with another city
	cities := []city{
		{name: "c1", country: "country1", population: 10000},
		{name: "c2", country: "country1", population: 20000},
		{name: "c3", country: "country1", population: 10010},
		{name: "c4", country: "country1", population: 11010}, // c1 is connected with c3, c3 connected with c4, therefore c1, c3, c4 are in the same group
		{name: "c5", country: "country1", population: 20010},
		{name: "c6", country: "country2", population: 20010},
		{name: "c7", country: "country3", population: 20010},
		{name: "c8", country: "country3", population: 20011},
	}

	cityGraph := graph.NewGraph()
	// add cities to the graph
	for _, currentCity := range cities {
		newNode := graph.NewNode(currentCity.name, currentCity)
		for _, otherNode := range cityGraph.Nodes {
			otherCity := otherNode.Data.(city)
			if connected(currentCity, otherCity) {
				cityGraph.AddBidirectional(newNode, otherNode)
				// Finding one connected sufficient
				break
			}
		}

		// If not added yet, add as node
		if _, ok := cityGraph.Nodes[currentCity.name]; !ok {
			cityGraph.Add(newNode, nil)
		}
	}

	// Now find connected groups of cities
	var connected [][]city
	for _, node := range cityGraph.Nodes {
		if _, ok := node.MData[isVisited]; !ok {
			node.MData[isVisited] = true
			// Not visited yet
			connectedGroup := []city{node.Data.(city)}
			collectConnected(node, &connectedGroup)
			connected = append(connected, connectedGroup)
		}
	}

	// Sort before output for consistent result
	for _, group := range connected {
		sort.Slice(group, func(i, j int) bool {
			return group[i].name < group[j].name
		})
	}

	sort.Slice(connected, func(i, j int) bool {
		return connected[i][0].name < connected[j][0].name
	})

	fmt.Printf("%v", connected)

	// Output:
	// [[{c1 country1 10000} {c3 country1 10010} {c4 country1 11010}] [{c2 country1 20000} {c5 country1 20010}] [{c6 country2 20010}] [{c7 country3 20010} {c8 country3 20011}]]
}

// recursively visit all reachable nodes and add to connectedgroup
func collectConnected(node *graph.Node, connectedGroup *[]city) {
	if node == nil {
		return
	}

	for _, vertex := range node.EdgesOut {
		otherNode := vertex.Target
		if _, ok := otherNode.MData[isVisited]; !ok {
			// Not visited yet
			otherNode.MData[isVisited] = true
			*connectedGroup = append(*connectedGroup, otherNode.Data.(city))
			collectConnected(otherNode, connectedGroup)
		}
	}
}

func connected(c1, c2 city) bool {
	// assume case sensitive comparison
	return c1.country == c2.country && int(math.Abs(float64(c1.population-c2.population))) <= 1000
}
