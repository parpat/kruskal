package kruskal

import "testing"

var (
	testEdges = EdgeList{
		NewEdge(4, 5, 35),
		NewEdge(4, 7, 37),
		NewEdge(5, 7, 28),
		NewEdge(0, 7, 16),
		NewEdge(1, 5, 32),
		NewEdge(0, 4, 38),
		NewEdge(2, 3, 17),
		NewEdge(1, 7, 19),
		NewEdge(0, 2, 26),
		NewEdge(1, 2, 36),
		NewEdge(1, 3, 29),
		NewEdge(2, 7, 34),
		NewEdge(6, 2, 40),
		NewEdge(3, 6, 52),
		NewEdge(6, 0, 58),
		NewEdge(6, 4, 93),
	}

	expectedMST = EdgeList{
		testEdges[3],
		testEdges[6],
		testEdges[7],
		testEdges[8],
		testEdges[2],
		testEdges[0],
		testEdges[12],
	}
)

func TestKruskal(t *testing.T) {
	g := NewGraph(8)
	for _, edge := range testEdges {
		g.AddEdge(edge)
	}

	mstEdges := kruskal(g)

	for i, mstEdge := range mstEdges {
		if mstEdge != expectedMST[i] {
			t.Fail()
			t.Logf("Got MST edge: W%d -- Expected MST edge: W%d", mstEdge.Weight, expectedMST[i].Weight)
		}
	}
}
