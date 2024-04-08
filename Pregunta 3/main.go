package main

import "fmt"

/*
Implementación de estructura de datos grafo por lista de adyacencias
y funciones relacionadas
*/
type Graph struct {
	adyLs [][]int
}

type Edge struct {
	a int
	b int
}

func newGraph(n int) Graph {
	ls := make([][]int, n)

	for i := 0; i < n; i++ {
		ls[i] = make([]int, 0, n)
	}

	return Graph{
		adyLs: ls,
	}
}

func (g Graph) addEdge(a int, b int) {
	g.adyLs[a] = append(g.adyLs[a], b)
	g.adyLs[b] = append(g.adyLs[b], a)
}

func (g Graph) degree(a int) int {
	return len(g.adyLs[a])
}

func (g Graph) adyacents(a int) []int {
	return g.adyLs[a]
}

func (g Graph) edges() []Edge {
	n := len(g.adyLs)
	ls := make([]Edge, 0)

	for i := 0; i < n; i++ {
		for _, e := range g.adyLs[i] {
			ls = append(ls, Edge{i, e})
		}
	}

	return ls
}

/*
Implementación del algoritmo propuesto para aproximación de
la cobertura mínima de un grafo G.
*/
func minVertexAprox(g Graph) []int {
	n := len(g.adyLs)
	vertCover := make([]int, 0, n)
	eliminado := make([][]bool, n)

	for i := 0; i < n; i++ {
		eliminado[i] = make([]bool, n)
	}

	for _, e := range g.edges() {
		if !eliminado[e.a][e.b] {
			da := g.degree(e.a)
			db := g.degree(e.b)
			if da < db {
				vertCover = append(vertCover, e.b)
				for _, ad := range g.adyacents(e.b) {
					eliminado[e.b][ad] = true
					eliminado[ad][e.b] = true
				}
			} else {
				vertCover = append(vertCover, e.a)
				for _, ad := range g.adyacents(e.a) {
					eliminado[e.a][ad] = true
					eliminado[ad][e.a] = true
				}
			}
		}
	}

	return vertCover
}

func main() {
	//g := newGraph(6)

	/*
		g.addEdge(0, 1)
		g.addEdge(0, 2)
		g.addEdge(1, 2)
		g.addEdge(1, 3)
		g.addEdge(2, 4)
		g.addEdge(3, 4)
		g.addEdge(3, 5)
	*/

	//Grafo de ejmplo 2
	/*
		g.addEdge(0, 1)
		g.addEdge(0, 2)
		g.addEdge(1, 2)
		g.addEdge(1, 3)
		g.addEdge(1, 4)
		g.addEdge(1, 5)
	*/

	//Grafo de ejemplo 3
	g := newGraph(10)
	g.addEdge(1, 3)
	g.addEdge(1, 5)
	g.addEdge(5, 2)
	g.addEdge(5, 4)
	g.addEdge(2, 9)
	g.addEdge(9, 6)
	g.addEdge(5, 6)
	g.addEdge(4, 6)
	g.addEdge(3, 7)
	g.addEdge(4, 8)

	fmt.Printf("Minimum vertex cover for Graph g: %v\n", minVertexAprox(g))
}
