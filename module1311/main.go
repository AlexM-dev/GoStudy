package main

import (
	"fmt"
	"mm/graph"
)

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Println("TREE TEST STARTS")
	tree := graph.NewTree(7)

	tree.Add(2000)
	tree.Add(12)
	tree.Add(2)
	tree.Add(22)
	tree.Add(1)
	tree.Add(122)
	tree.Print()
	fmt.Println(tree.Find(22))
	fmt.Println(tree.Delete(22))
	fmt.Println(tree.Find(22))
	fmt.Println(tree.Delete(22))
	tree.Print()
	tree.Add(-13)
	tree.Add(12202)
	tree.Print()
	fmt.Println("TREE TEST ENDS")

	//Графы без весов просто для удобства, добавляются они при необходимости в несколько строчек
	fmt.Println("-----------------------------------------")

	fmt.Println("NON-ORT-GRAPH TEST STARTS")
	g1 := graph.NewGraph()

	for i := 1; i <= 10; i++ {
		g1.AddVertex(i)
	}
	g1.AddEdge(1, 3)
	g1.AddEdge(3, 4)
	g1.AddEdge(4, 2)
	g1.AddEdge(2, 8)
	g1.AddEdge(2, 7)
	g1.BFS(1)
	fmt.Println("NON-ORT-GRAPH TEST ENDS")

	fmt.Println("-----------------------------------------")

	fmt.Println("ORT-GRAPH TEST STARTS")
	g := graph.NewGraph()

	for i := 1; i <= 10; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 2)
	g.AddEdge(2, 8)
	g.AddEdge(2, 7)
	g.FindPath(1, 7)
	fmt.Println("ORT-GRAPH TEST ENDS")

	fmt.Println("-----------------------------------------")
}
