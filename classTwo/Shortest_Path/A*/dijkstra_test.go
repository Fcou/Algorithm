package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	g := InitGraph(6)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 4, 15)
	g.AddEdge(1, 2, 15)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 5, 5)
	g.AddEdge(3, 2, 1)
	g.AddEdge(3, 5, 12)
	g.AddEdge(4, 5, 10)
	Dijkstra(0, 5, g)
}
