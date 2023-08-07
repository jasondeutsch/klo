# Flight Planner

## About this project

This project implements a toy flight planning system. The purpose of this project is to demonstrate the use of graph data structures.
 
This project is two components:
1. `DirectedGraph` data type with associated methods
2. `FlightPLanner` data type which utilizes the graph for its internal workings. 

### Features 
* Add and remove flights from the network.
* Find the shortest path (in terms of flight duration) between two given airports.
* Find all flights from a given airport.

## Usage

### Dependencies

1. Go >= 1.16 
2. Make (for convenience)
3. Docker (for portability) 

### Testing

```shell
make test
# or
go test -v ./...
```


## Demo 
A small demo program is provided to show case the requirements of the Flight Planner.
```shell
# only required dependency is Docker
make docker-build
make docker-demo

# or 

go run ./...
```

## Library Usage:

### `flightplanner.go`
Provides a type and methods for basic airport flight planning:
#### Example (error handling omitted for brevity):
```go
planner := NewPlanner()

// Add(from, to, duration)
planner.Add("JFK", "LAX", 42)
planner.Add("LAX", "AUS", 55)

// List flights for airpoirt
flights, _ := planner.Flights("JFK")

// Find fastest route
path, dist := planner.ShortestTravelPath("JFK", "AUS")

// Remove a flight
planner.Remove("LAX", "AUS")

```

### `graph.go`
The underlying graph can be used directly
#### Example (error handling omitted for brevity):
```go
type DirectedGraph struct {
  Vertices []*Vertex
}

type Vertex struct {
  Key      string
  Adjacent []*Edge
}

type Edge struct {
  To     string
  Weight int
}

func (g *DirectedGraph) AddVertex(key string) error 
func (g *DirectedGraph) AddEdge(from, to string, weight int) error 
func (g *DirectedGraph) RemoveEdge(from, to string) error 
func (g *DirectedGraph) Find(key string) (*Vertex, bool) 
func (g *DirectedGraph) ShortestPath(from, to string) ([]*Vertex, int)
```


## Shortcomings and TODOs
* Airport codes should be case-insensitive
* `flightplanner.go`  deserves unit tests of its own (omitted because of demo command)