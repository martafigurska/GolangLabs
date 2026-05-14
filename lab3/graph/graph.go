package graph

import (
	"math/rand"
	"errors"
)

type graph struct {
	edges map[int]map[int]bool
	n int
}

func CreateGraph(n int) *graph {
	g := &graph{
		edges: make(map[int]map[int]bool),
		n: n,
	}
	p := 2.0 / float64(n)

	for i := 0; i < n; i++ {
		g.edges[i] = make(map[int]bool)
		for j := 0; j < n; j++ {
			if i != j && rand.Float64() < p {
				g.edges[i][j] = true
			}
			// if g.edges[i][j] == true {

				// }
		}
	}
	return g
}

func (g *graph) String() string {
	result := ""
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			if g.edges[i][j] == true{
				result += "1 "
			} else {
				result += "0 "
			}
		}
		result += "\n"
	}
	return result
}

func (g *graph) AddEdge(i, j int) error{
	if i < 0 || i >= g.n || j < 0 || j >= g.n {
		return errors.New("Wrong idx")
	}
	if g.edges[i][j] != true {
		g.edges[i][j] = true
	} else {
		return errors.New("This edge already exists")
	}
	return nil
}

func (g *graph) CopyGraph() *graph {
	g2 := &graph{
		edges: make(map[int]map[int]bool),
		n: g.n,
	}
	for i := 0; i < g.n; i++ {
		g2.edges[i] = make(map[int]bool)
		for j, val := range g.edges[i] {
			g2.edges[i][j] = val
		}
	}
	return g2
}

func (g *graph) CompareGraphs(g2 *graph) bool {
	if g.n != g2.n {
		return false
	}

	for i := 0; i < g.n; i++ {
		if len(g.edges[i]) != len(g2.edges[i]) {
			return false
		}
		for j := range g.edges[i] {
			if g.edges[i][j] != g2.edges[i][j] {
				return false
			}
		}
	}

	return true
}

func (g *graph) DegreeMap() (map[int]int) {
	outDeg := make(map[int]int)
	// inDeg := make(map[int]int)

	for i := 0; i < g.n; i++ {
		outDeg[i] = len(g.edges[i])
	}

	return outDeg
}