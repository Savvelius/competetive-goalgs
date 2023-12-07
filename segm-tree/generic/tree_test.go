package generic

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := New(func(i1, i2 int) int {
		return i1 + i2
	}, 0, []int{1, 2, 3, 4, 5})

	q := tree.Query(2, 4)
	if q != 7 {
		t.Errorf("expected query to produce %d, got %d", 7, q)
	}

	q = tree.Query(1, 5)
	if q != 14 {
		t.Errorf("expected query to produce %d, got %d", 7, q)
	}

	tree.Apply(0, 5, 10)
	tree.Pprint()
}
