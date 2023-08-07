package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirectedGraph_AddVertexAndEdge(t *testing.T) {
	g := DirectedGraph{}

	require.NoError(t, g.AddVertex("a"))
	require.NoError(t, g.AddVertex("b"))
	require.NoError(t, g.AddEdge("a", "b", 5))

	vert, ok := g.Find("a")
	require.True(t, ok)
	require.NotNil(t, vert)
	require.Len(t, vert.Edges, 1)
}

func TestDirectedGraph_RemoveEdge(t *testing.T) {
	g := DirectedGraph{}

	require.NoError(t, g.AddVertex("a"))
	require.NoError(t, g.AddVertex("b"))
	require.NoError(t, g.AddEdge("a", "b", 5))
	require.NoError(t, g.RemoveEdge("a", "b"))

	vert, ok := g.Find("a")
	require.True(t, ok)
	require.NotNil(t, vert)
	require.Len(t, vert.Edges, 0)
}

func TestGraph_ShortestPath(t *testing.T) {
	g := DirectedGraph{}

	require.NoError(t, g.AddVertex("a"))
	require.NoError(t, g.AddVertex("b"))
	require.NoError(t, g.AddVertex("c"))
	require.NoError(t, g.AddVertex("d"))
	require.NoError(t, g.AddVertex("f"))
	require.Error(t, g.AddVertex("f"))

	// test case: shortest path a -> c
	require.NoError(t, g.AddEdge("a", "b", 5))
	require.NoError(t, g.AddEdge("b", "c", 5))
	require.NoError(t, g.AddEdge("a", "d", 4))
	require.NoError(t, g.AddEdge("d", "c", 4))

	path, dist := g.ShortestPath("a", "c")
	require.Len(t, path, 3)
	require.Equal(t, 8, dist)

	// test case: shortest path a -> g
	require.NoError(t, g.AddVertex("g"))
	require.NoError(t, g.AddVertex("z"))
	require.NoError(t, g.AddEdge("a", "g", 100))
	require.NoError(t, g.AddEdge("d", "z", 1))
	require.NoError(t, g.AddEdge("z", "g", 1))

	path, dist = g.ShortestPath("a", "g")
	require.Len(t, path, 4)
	require.Equal(t, dist, 6)
}
