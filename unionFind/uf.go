package unionFind

//UF holds the IDs of components
type UF struct {
	id []int
	sz []int
}

//InitializeUF initializes the DS by putting each element in its own compoenent
func InitializeUF(n int) UF {
	u := UF{make([]int, n), make([]int, n)}

	for i := 0; i < n; i++ {
		u.id[i] = i
		u.sz[i] = 1
	}
	return u
}

//Find finds the root of the tree O(lg n) based on the depth
func (u UF) Find(i int) int {
	for i != u.id[i] {
		//While traversing towards the root, the node points to its
		//grandparent to reduce the path length to the root
		u.id[i] = u.id[u.id[i]]
		i = u.id[i]
	}
	return i
}

//Connected checks if the elements have the same parent
func (u UF) Connected(e1, e2 int) bool {

	return u.Find(e1) == u.Find(e2)
}

//Union merges the components of the two elements.
//Weighted Union tries to avoid tall trees by checking the size before union.
//The smaller tree changes its root.
func (u UF) Union(e1, e2 int) {
	e1R := u.Find(e1)
	e2R := u.Find(e2)
	if u.sz[e1] < u.sz[e2] {
		u.id[e1R] = e2R
		u.sz[e2R] += u.sz[e1R]
	} else {
		u.id[e2R] = e1R
		u.sz[e1R] += u.sz[e2R]
	}
}
