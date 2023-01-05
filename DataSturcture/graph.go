package DataSturcture

type Node[T any] struct {
	key   string
	value T
}

type Edge[T any] struct {
	from   string
	to     string
	weight int
}

type Graph[T any] struct {
	nodes []Node[T]
	edges []Edge[T]
}

func NewGraph[T any]() *Graph[T] {
	return &Graph[T]{nodes: make([]Node[T], 0), edges: make([]Edge[T], 0)}
}

func (g *Graph[T]) AddNode(key string, value T) {
	g.nodes = append(g.nodes, Node[T]{key: key, value: value})
}

func (g *Graph[T]) AddEdge(from string, to string, weight int) {
	g.edges = append(g.edges, Edge[T]{from: from, to: to, weight: weight})
}

func (g *Graph[T]) GetNode(key string) Node[T] {
	for _, node := range g.nodes {
		if node.key == key {
			return node
		}
	}
	return Node[T]{}
}

func (g *Graph[T]) GetEdge(from string, to string) Edge[T] {
	for _, edge := range g.edges {
		if edge.from == from && edge.to == to {
			return edge
		}
	}
	return Edge[T]{}
}

func (g *Graph[T]) GetNodes() []Node[T] {
	return g.nodes
}

func (g *Graph[T]) GetEdges() []Edge[T] {
	return g.edges
}

func (g *Graph[T]) GetNeighbors(key string) []Node[T] {
	neighbors := make([]Node[T], 0)
	for _, edge := range g.edges {
		if edge.from == key {
			neighbors = append(neighbors, g.GetNode(edge.to))
		}
	}
	return neighbors
}

func getMinKey(visited []string, dist map[string]int) string {
	min := 2147483647
	var minKey string
	for _, key := range visited {
		if dist[key] < min {
			min = dist[key]
			minKey = key
		}
	}
	return minKey
}

func remove(s []string, i string) []string {
	for index, value := range s {
		if value == i {
			return append(s[:index], s[index+1:]...)
		}
	}
	return s
}

func (g *Graph[T]) ShortestPath(from string, to string) []Node[T] {
	// Dijkstra algorithm
	dist := make(map[string]int)
	prev := make(map[string]string)
	visited := make([]string, 0)
	for _, node := range g.nodes {
		prev[node.key] = ""
		dist[node.key] = 1<<31 - 1
		visited = append(visited, node.key)
	}
	dist[from] = 0

	for len(visited) > 0 {
		u := getMinKey(visited, dist)
		if u == to || u == "" {
			break
		}
		visited = remove(visited, u)
		neighbors := g.GetNeighbors(u)
		for _, neighbor := range neighbors {
			for _, s := range visited {
				if neighbor.key != s {
					continue
				}
				alt := dist[u] + g.GetEdge(u, neighbor.key).weight
				if alt < dist[neighbor.key] {
					dist[neighbor.key] = alt
					prev[neighbor.key] = u
				}
			}
		}
	}

	// Get path
	path := make([]Node[T], 0)
	u := to
	for u != "" {
		path = append(path, g.GetNode(u))
		u = prev[u]
	}
	path = append(path, g.GetNode(from))
	return path
}
