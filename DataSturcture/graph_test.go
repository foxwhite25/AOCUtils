package DataSturcture

import "testing"

func TestGraph_ShortestPath(t *testing.T) {
	graph := NewGraph[int]()
	graph.AddNode("1", 1)
	graph.AddNode("2", 2)
	graph.AddNode("3", 3)
	graph.AddNode("4", 4)
	graph.AddNode("5", 5)

	graph.AddEdge("1", "2", 1)
	graph.AddEdge("4", "1", 2)
	graph.AddEdge("2", "3", 2)
	graph.AddEdge("1", "3", 5)
	graph.AddEdge("3", "5", 1)

	path := graph.ShortestPath("1", "5")
	t.Log(len(path))
}
