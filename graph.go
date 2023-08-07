package main

import (
	"fmt"
	"math"
)

// DirectedGraph is a directed, acyclic graph
type DirectedGraph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key   string
	Edges []*Edge
}

type Edge struct {
	To     string
	Weight int
}

func (g *DirectedGraph) AddVertex(key string) error {
	if _, ok := g.Find(key); ok {
		return fmt.Errorf("duplicate Key: %s", key)
	}

	g.Vertices = append(g.Vertices, &Vertex{Key: key})

	return nil
}

func (g *DirectedGraph) AddEdge(from, to string, weight int) error {
	// find Vertex
	fromVert, ok := g.Find(from)
	if !ok {
		return fmt.Errorf("from fromVert %s not found", from)
	}

	_, ok = g.Find(to)
	if !ok {
		return fmt.Errorf("from fromVert %s not found", to)
	}

	// add Edge
	fromVert.Edges = append(fromVert.Edges, &Edge{
		To:     to,
		Weight: weight,
	})

	return nil
}

func (g *DirectedGraph) RemoveEdge(from, to string) error {
	fromVert, ok := g.Find(from)
	if !ok {
		return fmt.Errorf("from fromVert %s not found", from)
	}

	var edge *Edge
	edgeIdx := -1 // theoretically can panic
	for i, e := range fromVert.Edges {
		if e.To == to {
			edge = e
			edgeIdx = i
		}
	}

	if edge == nil {
		if !ok {
			return fmt.Errorf("edge (%s, %s) not found", from, to)
		}
	}

	fromVert.Edges = append(fromVert.Edges[:edgeIdx], fromVert.Edges[edgeIdx+1:]...)

	return nil
}

func (g *DirectedGraph) Find(key string) (*Vertex, bool) {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v, true
		}
	}
	return nil, false
}

func (g *DirectedGraph) ShortestPath(from, to string) ([]*Vertex, int) {
	fromVert, ok := g.Find(from)
	if !ok {
		return nil, -1
	}

	toVertex, ok := g.Find(to)
	if !ok {
		return nil, -1
	}

	distances := make(map[*Vertex]int)
	for _, v := range g.Vertices {
		distances[v] = math.MaxInt32
	}
	distances[fromVert] = 0

	parents := make(map[*Vertex]*Vertex)
	q := []*Vertex{fromVert}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, edge := range curr.Edges {
			next, _ := g.Find(edge.To)
			currDist := distances[curr] + edge.Weight
			if currDist < distances[next] {
				distances[next] = currDist
				parents[next] = curr
				q = append(q, next)
			}
		}
	}

	// make the path
	var path []*Vertex
	curr := toVertex
	for curr != nil {
		path = append([]*Vertex{curr}, path...)
		curr = parents[curr]
	}

	if path[0] != fromVert {
		return nil, -1
	}

	return path, distances[toVertex]
}
