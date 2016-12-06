package mst

//Input: A connected graph G whose edges have distinct weights
//1   Initialize a forest T to be a set of one-vertex trees, one for each vertex of the graph.
//2   While T has more than one component:
// 3     For each component C of T:
// 4       Begin with an empty set of edges S
// 5       For each vertex v in C:
// 6         Find the cheapest edge from v to a vertex outside of C, and add it to S
// 7       Add the cheapest edge in S to T
// 8
//   Combine trees connected by edges to form bigger components
// 9   Output: T is the minimum spanning tree of G.

func boruvka(g *Graph) EdgeList {
	//mstEdges := *new(EdgeList)

	//forest := unionFind.InitializeUF(len(g.Vertices))

	// while(forest.GetNumComponents() > 1){
	// for _, tree :=
	//}

	return nil
}
