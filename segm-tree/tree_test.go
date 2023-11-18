package segmtree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	tree := New(arr)
	fmt.Println(tree.GetSum(0, 3))
}
