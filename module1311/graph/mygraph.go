package graph

import (
	"fmt"
)

type graph struct {
	adj map[int][]int
}

func NewGraph() *graph {
	return &graph{adj: make(map[int][]int)}
}

func (g *graph) AddVertex(vertex int) {
	g.adj[vertex] = []int{}
}

func (g *graph) AddEdge(vertex1, vertex2 int) {
	g.adj[vertex1] = append(g.adj[vertex1], vertex2)
}

func (g *graph) AddEdgeN(vertex1, vertex2 int) {
	g.adj[vertex1] = append(g.adj[vertex1], vertex2)
	g.adj[vertex2] = append(g.adj[vertex2], vertex1)
}

func (g *graph) BFS(startingNode int) {
	visited := make(map[int]bool, len(g.adj))
	for k := range g.adj {
		visited[k] = false
	}
	var q []int

	visited[startingNode] = true
	q = append(q, startingNode)

	for len(q) > 0 {
		var current int
		current, q = q[0], q[1:]
		fmt.Println(current)
		for _, node := range g.adj[current] {
			if !visited[node] {
				q = append(q, node)
				visited[node] = true
			}
		}
	}
}

func (g *graph) FindPath(firstNode, secondNode int) {
	visited := make(map[int]bool, len(g.adj))
	for k := range g.adj {
		visited[k] = false
	}
	var path, q []int
	q = append(q, firstNode)
	visited[firstNode] = true

	for len(q) > 0 {
		var currentNode int
		currentNode, q = q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adj[currentNode]
		if contains(edges, secondNode) {
			path = append(path, secondNode)
			for i := 0; i < len(path)-1; i++ {
				fmt.Print(path[i], "->")
			}
			fmt.Println(path[len(path)-1])
			return
		}

		for _, node := range g.adj[currentNode] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}
	fmt.Println("no path found")
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
