package unionFind

import (
	"fmt"
	"testing"
)

var (
	uftestSuite = []struct{ from, to int }{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{5, 0},
		{7, 2},
		{6, 1},
	}
)

//TODO: Proper testing
func TestUnion(t *testing.T) {
	ufTestT := InitializeUF(10)
	for _, pair := range uftestSuite {
		fmt.Println(ufTestT.Connected(pair.from, pair.to))
		ufTestT.Union(pair.from, pair.to)
		fmt.Println(ufTestT.Connected(pair.from, pair.to))
		fmt.Printf("Total components:%v\n", ufTestT.GetNumComponents())
	}

}
