package kruskal

import "sort"

//Vertex can be anything
type Vertex interface{}

//Edge is between to Vertices and has a weight
type Edge struct {
	Start  Vertex
	End    Vertex
	Weight int
}

//EdgeList will be the Edge list
type EdgeList []Edge

func (a EdgeList) Len() int           { return len(a) }
func (a EdgeList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EdgeList) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

//Graph is the main DS
type Graph struct {
	Edges    EdgeList
	Vertices []Vertex
}

//NewGraph creates a new graph with the specified number of Vertices
func NewGraph(v int) Graph {
	var g Graph
	for i := 0; i < v; i++ {
		g.AddVertex(i)
	}
	return g
}

//AddVertex used to add vertices
func (g *Graph) AddVertex(v Vertex) {
	g.Vertices = append(g.Vertices, v)
}

// AddEdge adds an edge to the graph. Then sorts the edges based on weights
func (g *Graph) AddEdge(e Edge) {
	g.Edges = append(g.Edges, e)
	sort.Sort(EdgeList(g.Edges))
}

func NewEdge(v1, v2 Vertex, w int) Edge {
	return Edge{v1, v2, w}
}
