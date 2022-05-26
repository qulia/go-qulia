package graph_test

import (
	"fmt"
	"sort"

	"github.com/qulia/go-qulia/lib/graph"
)

// This is a problem where we need to find the itinerary using the tickets
// e.g. tickets [A, B] [B, C] [B,D] [C,B] itinerary A->B->C->B->D using all tickets
// Note that if we had picked A->B->D we would not be able to use all tickets, so not the right path
// Since we need to keep track of all edges we visited, metadata support on the edge can be used
// There could be multiple tickets available for the same source, dest pair. So we can keep track of
// "available edges" in the metadata
func ExampleGraph_itinerary() {
	ticketsInputs := [][]ticket{
		{{"A", "B"}, {"B", "C"}, {"B", "D"}, {"C", "B"}},
		{{"A", "B"}, {"A", "C"}, {"C", "A"}},
		// Note there are two tickets D,B
		{{"E", "F"}, {"D", "B"}, {"B", "C"}, {"C", "B"}, {"B", "E"}, {"D", "B"}, {"F", "D"}, {"D", "C"}, {"B", "D"}, {"C", "D"}},
	}

	for _, input := range ticketsInputs {
		fmt.Printf("%v\n", testWithTickets(input))
	}

	//Output:
	//[A B C B D]
	//[A C A B]
	//[B C B D B E F D C D B]
}

func testWithTickets(tickets []ticket) []string {
	// Create graph with source, dest in tickets
	cGraph := graph.NewGraph[string]()
	ticketCount := map[ticket]int{}
	for _, ticket := range tickets {
		ticketCount[ticket]++
		cGraph.Add(ticket.src, ticket.dst)
	}

	var itinerary []string

	// sort nodes to get lexical order in result
	cities := cGraph.GetNodes().ToSlice()
	sort.Strings(cities)
	n := len(tickets)
	for _, c := range cities {
		itinerary = append(itinerary, c)
		if visitAll(c, cGraph, 0, n, ticketCount, &itinerary) {
			break
		} else {
			itinerary = itinerary[:len(itinerary)-1]
		}
	}

	return itinerary
}

func visitAll(c string, g graph.Graph[string], t, n int, ticketCount map[ticket]int, itinerary *[]string) bool {
	if t == n {
		return true
	}

	// sort targets to get lexical order in result
	var ds []string
	for n, _ := range g.Adjacencies(c) {
		ds = append(ds, n)
	}
	sort.Strings(ds)

	for _, d := range ds {
		tc := ticket{c, d}
		if ticketCount[tc] > 0 {
			ticketCount[tc]--
			*itinerary = append(*itinerary, d)
			if visitAll(d, g, t+1, n, ticketCount, itinerary) {
				return true
			} else {
				// backtrack so it can be chosen again
				ticketCount[tc]++
				*itinerary = (*itinerary)[:len(*itinerary)-1]
			}
		}
	}

	return false
}

type ticket struct {
	src, dst string
}
