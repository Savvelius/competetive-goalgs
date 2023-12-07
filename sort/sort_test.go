package sort

import (
	"math/rand"
	"slices"
	"testing"
)

func TestSorts(t *testing.T) {
	a := rand.Perm(1000)
	Choice(a)
	if !slices.IsSorted(a) {
		t.Errorf("choice sort failed with result %v", a[:10])
	}

	a = rand.Perm(1000)
	Bubble(a)
	if !slices.IsSorted(a) {
		t.Errorf("bubble sort failed with result %v", a[:10])
	}

	a = rand.Perm(1000)
	Insertion(a)
	if !slices.IsSorted(a) {
		t.Errorf("insertion sort failed with result %v", a[:10])
	}

	a = rand.Perm(1000)
	a = Heap(a)
	if !slices.IsSorted(a) {
		t.Errorf("heap sort failed with result %v", a[:10])
	}

	a = rand.Perm(1000)
	Count(a, 0, 999)
	if !slices.IsSorted(a) {
		t.Errorf("count sort failed with result %v", a[:15])
	}

	a = rand.Perm(1000)
	a = Merge(a)
	if !slices.IsSorted(a) {
		t.Errorf("merge sort failed with result %v", a[:15])
	}

	a = rand.Perm(1000)
	Quick(a)
	if !slices.IsSorted(a) {
		t.Errorf("quick sort failed with result %v", a[:15])
	}
}
