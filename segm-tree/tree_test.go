package segmtree

import (
	"testing"
)

func TestTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	tree := New(arr)
	// tree.Add(2, 7)
	// tree.Pprint()

	tree.Assign(2, 4, 10)
	tree.Add(4, 10)
	tree.Pprint()
	// fmt.Println(tree.GetSum(0, 5))
}
