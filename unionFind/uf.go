package unionFind

//UF holds the IDs of components
type UF struct {
	ID []int
	Sz []int
}

//InitializeUF initializes the DS by putting each element in its own compoenent
func InitializeUF(n int) UF {
	u := UF{make([]int, n), make([]int, n)}

	for i := 0; i < n; i++ {
		u.ID[i] = i
		u.Sz[i] = 1
	}
	return u
}

//Find finds the root of the tree O(lg n) based on the depth
func (u *UF) Find(i int) int {
	for i != u.ID[i] {
		//While traversing towards the root, the node points to its
		//grandparent to reduce the path length to the root
		u.ID[i] = u.ID[u.ID[i]]
		i = u.ID[i]
	}
	return i
}

//Connected checks if the elements have the same parent
func (u *UF) Connected(e1, e2 int) bool {

	return u.Find(e1) == u.Find(e2)
}

//Union merges the components of the two elements.
//Weighted Union tries to avoID tall trees by checking the size before union.
//The smaller tree changes its root.
func (u *UF) Union(e1, e2 int) {
	e1R := u.Find(e1)
	e2R := u.Find(e2)
	if u.Sz[e1] < u.Sz[e2] {
		u.ID[e1R] = e2R
		u.Sz[e2R] += u.Sz[e1R]
	} else {
		u.ID[e2R] = e1R
		u.Sz[e1R] += u.Sz[e2R]
	}
}

//GetNumComponents returns the total number of components
func (u *UF) GetNumComponents() int {
	componentTrees := make(map[int]int)
	for _, c := range u.ID {
		componentTrees[u.Find(c)]++
	}
	return len(componentTrees)
}
