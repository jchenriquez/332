package main

import ("sort"
        "fmt"
        "strings")

func graphIt(tickets [][]string) map[string][]string {
    graph := make(map[string][]string)

    for _, itinerary := range tickets {
      flightTo, in := graph[itinerary[0]]

      if !in {
        flightTo = make([]string, 0)
      }

      flightTo = append(flightTo, itinerary[1])

      graph[itinerary[0]] = flightTo
    }
    return graph
}

func toPathKey (airportA, airportB string) string {
  return fmt.Sprintf("%s->%s", airportA, airportB)
}

func toLexico(airline string) int {
  var value int

  for i := 0; i < len(airline); i++ {
    value += int(airline[i])
  } 

  return value
}

func compareTo(collected1, collected2 []string) int {
  temp := strings.Join(collected1, "")
  orig := strings.Join(collected2, "")

  fmt.Printf("temp %s\n", temp)
  fmt.Printf("orig %s\n", orig)

  switch {
    case temp < orig:
      return -1
    case temp > orig:
      return 1
    default:
      return 0
  }
}

func dFsIt(graph map[string][]string, currAirport string, seen map[string]bool, path *[]string) {
  airport, in := graph[currAirport]
  
  if in {
    airportCpy := make([]string, len(airport))
    copy(airportCpy, airport)
    sort.Strings(airportCpy)
    for i := 0; i < len(airportCpy); i++ {
      flight := airportCpy[i]
      key := toPathKey(currAirport, flight)

      if saw, _ := seen[key]; !saw {
        seen[key] = true
        *path = append(*path, flight)
        dFsIt(graph, flight, seen, path)
      }
    }

  }

}

func findItinerary(tickets [][]string) []string {
    memoized := make(map[string]bool)
    graph := graphIt(tickets)
    minItinerary := &[]string{"JFK"}
    
    dFsIt(graph, "JFK", memoized, minItinerary)
    return *minItinerary
}

func main() {
  fmt.Printf("%v\n", findItinerary([][]string{[]string{"JFK","SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"}, []string{"ATL", "JFK"}, []string{"ATL", "SFO"}}))
  fmt.Printf("%d", compareTo([]string{"JFK", "KUL", "NRT", "JFK"}, []string{"JFK", "NRT", "JFK", "KUL"}))
}