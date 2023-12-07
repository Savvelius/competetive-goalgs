package array

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := New([]int{1, 2, 3, 4, 5}, func(i1, i2 int) int {
		return i1 + i2
	}, 0)

	q := tree.Query(0, 5)
	if q != 15 {
		t.Errorf("expected query(0, 5) to be %d got %d", 15, tree.Query(0, 5))
	}

	q = tree.Query(2, 4)
	if q != 7 {
		t.Errorf("expected query(2, 4) to be %d got %d", 7, tree.Query(2, 4))
	}

	tree.Apply(2, 4, 10)

	q = tree.Query(2, 4)
	if q != 27 {
		t.Errorf("expected query(2, 4) to be %d got %d", 27, tree.Query(2, 4))
	}

	q = tree.Query(0, 5)
	if q != 35 {
		t.Errorf("expected query(0, 5) to be %d got %d", 35, tree.Query(0, 5))
	}

	tree.Set(0, 5, 10)

	q = tree.Query(0, 5)
	if q != 50 {
		t.Errorf("expected query(0, 5) to be %d got %d", 50, tree.Query(0, 5))
	}

	tree.Set(2, 4, 0)

	q = tree.Query(0, 5)
	if q != 30 {
		t.Errorf("expected query(0, 5) to be %d got %d", 30, tree.Query(0, 5))
	}

	q = tree.Query(2, 4)
	if q != 0 {
		t.Errorf("expected query(0, 5) to be %d got %d", 0, tree.Query(2, 4))
	}
}
