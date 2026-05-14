package main

import (
	"fmt"
	"lab3/graph"
)

func main() {
	g := graph.CreateGraph(5)
	fmt.Println(g)

	err := g.AddEdge(1, 3) 
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(g)

	gCopy := g.CopyGraph()
	fmt.Println("Copied graph:")
	fmt.Println(gCopy)

	gDiff := graph.CreateGraph(5)

	if g.CompareGraphs(gCopy) == true {
		fmt.Println("The same")
	} else {
		fmt.Println("Not the same")
	}

	if g.CompareGraphs(gDiff) == true {
		fmt.Println("The same")
	} else {
		fmt.Println("Not the same")
	}

	outD := g.DegreeMap()
	fmt.Println("\nOut degrees in graph g")
	fmt.Println(outD)
}