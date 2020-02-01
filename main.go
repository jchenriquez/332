package main

import (
	"container/heap"
	"fmt"
)

type Itinerary []string

func (hp Itinerary) Len() int {
	return len(hp)
}

func (hp Itinerary) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp Itinerary) Less(i, j int) bool {
	return hp[i] < hp[j]
}

func (hp *Itinerary) Push(val interface{}) {
  (*hp) = append(*hp, val.(string))
}

func (hp *Itinerary) Pop() interface{} {
	ret := (*hp)[len(*hp)-1]
	*hp = (*hp)[:len(*hp)-1]

	return ret
}

func graphIt(tickets [][]string) (map[string]Itinerary, int) {
	graph := make(map[string]Itinerary)
  var numOfEdges int

	for _, ticket := range tickets {
		itinerary, in := graph[ticket[0]]

		if !in {
			itinerary = make(Itinerary, 0)
		}
		ptr := &itinerary
		heap.Init(ptr)
		heap.Push(ptr, ticket[1])
    numOfEdges++
		graph[ticket[0]] = *ptr
	}
	return graph, numOfEdges
}

func edgeKey(airportA, airportB string) string {
	return fmt.Sprintf("%s->%s", airportA, airportB)
}

func dFsIt(graph map[string]Itinerary, memoized map[string]bool, currAirport string, totalSumOfEdges int) []string {
  
  flightItinerary := graph[currAirport] 
  ptr := &flightItinerary

  recursedBack := make([][]string, 0, totalSumOfEdges+1)
  var oneWay []string

  for len(*ptr) > 0 {
    nFlight := heap.Pop(ptr).(string)
    edge := edgeKey(currAirport, nFlight)
    if _, saw := memoized[edge]; !saw {
      memoized[edge] = true
      retItinerary := dFsIt(graph, memoized, nFlight, totalSumOfEdges)

      switch {
       case retItinerary[len(retItinerary)-1] == currAirport:
        recursedBack = append(recursedBack, retItinerary)
       case oneWay == nil:
        oneWay = retItinerary
      }
    }
  }
  ret := make([]string, 0, totalSumOfEdges+1)

  ret = append(ret, currAirport)

  for _, reIti := range recursedBack {
    ret = append(ret, reIti...)
  }

  return append(ret, oneWay...)
}


func findItinerary(tickets [][]string) []string {
	graph, numOfEdges := graphIt(tickets)
  fmt.Printf("graph %v\n", graph)
	return dFsIt(graph, make(map[string]bool), "JFK", numOfEdges)
}

func main() {
	fmt.Printf("%v\n", findItinerary([][]string{{"EZE","AXA"},{"TIA","ANU"},{"ANU","JFK"},{"JFK","ANU"},{"ANU","EZE"},{"TIA","ANU"},{"AXA","TIA"},{"TIA","JFK"},{"ANU","TIA"},{"JFK","TIA"}}))
}
