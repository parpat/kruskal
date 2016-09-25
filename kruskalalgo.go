package kruskal

import "github.com/parpat/kruskal/unionFind"

func kruskal(g Graph) EdgeList {

	mstEdges := *new(EdgeList)

	components := unionFind.InitializeUF(len(g.Vertices))
	gsize := len(g.Vertices)

	for i := 0; len(mstEdges) != gsize-1; i++ {
		e := g.Edges[i]
		if !components.Connected(e.Start.(int), e.End.(int)) {
			components.Union(e.Start.(int), e.End.(int))
			mstEdges = append(mstEdges, e)
		}

	}

	return mstEdges
}
