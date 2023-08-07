package main

import (
	"fmt"
)

func NewPlanner() FlightPlanner {
	return FlightPlanner{
		flights: &DirectedGraph{},
	}
}

type FlightPlanner struct {
	flights *DirectedGraph
}

func (p FlightPlanner) Add(from, to string, duration int) error {
	// add airports if not present
	if _, ok := p.flights.Find(from); !ok {
		_ = p.flights.AddVertex(from)
	}

	if _, ok := p.flights.Find(to); !ok {
		_ = p.flights.AddVertex(to)
	}

	// add flight
	if err := p.flights.AddEdge(from, to, duration); err != nil {
		return fmt.Errorf("failed To add flight %s -> %s (%d)", from, to, duration)
	}

	return nil
}

func (p FlightPlanner) Remove(from, to string) error {
	if err := p.flights.RemoveEdge(from, to); err != nil {
		return fmt.Errorf("failed To remove flight %s -> %s", from, to)
	}

	return nil
}

// TODO(handle error case of no path)
func (p FlightPlanner) ShortestTravelPath(from, to string) ([]string, int) {
	pathVerts, dist := p.flights.ShortestPath(from, to)

	path := make([]string, len(pathVerts))
	for i, v := range pathVerts {
		path[i] = v.Key
	}

	return path, dist
}

type Flight struct {
	From, To string
	Duration int
}

func (p FlightPlanner) Flights(from string) ([]Flight, error) {
	vert, ok := p.flights.Find(from)
	if !ok {
		return nil, fmt.Errorf("airpoirt %s not found", from)
	}

	flights := make([]Flight, len(vert.Edges))
	for i, v := range vert.Edges {
		flights[i] = Flight{
			From:     from,
			To:       v.To,
			Duration: v.Weight,
		}
	}

	return flights, nil
}
