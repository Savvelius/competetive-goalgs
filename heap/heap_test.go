package heap

import (
	"math"
	"testing"
)

func TestInsertExtract(t *testing.T) {
	heap := New[int](func(i1, i2 int) bool { return i1 < i2 }, 5)
	if heap.Len() != 0 {
		t.Errorf("expected heap length to be 0, got %d", heap.Len())
	}
	heap.Insert(10)
	heap.Insert(-10)
	heap.Insert(20)
	heap.Insert(30)
	if heap.Len() != 4 {
		t.Errorf("expected heap length to be 4, got %d", heap.Len())
	}

	min := heap.ExtractMin()
	if min != -10 {
		t.Errorf("expected min value to be -10, got %d", min)
	}
	if heap.Len() != 3 {
		t.Errorf("expected heap length to be 3, got %d", heap.Len())
	}

	min = heap.ExtractMin()
	if min != 10 {
		t.Errorf("expected min value to be 10, got %d", min)
	}

	min = heap.ExtractMin()
	if min != 20 {
		t.Errorf("expected min value to be 20, got %d", min)
	}

	min = heap.ExtractMin()
	if min != 30 {
		t.Errorf("expected min value to be 30, got %d", min)
	}

	if heap.Len() != 0 {
		t.Errorf("expected heap length to be 0, got %d", heap.Len())
	}
}

func TestHeapifyToSlice(t *testing.T) {
	slice := []int{43, 52, -10, 432, 954, -90, 1, 2, 98542, -46238}
	heap := Heapify(slice, func(i1, i2 int) bool { return i1 < i2 })
	if heap.Len() != len(slice) {
		t.Errorf("expected heap len=%d to be same as slice len=%d", heap.Len(), len(slice))
	}

	slice = heap.ToSlice()
	if heap.Len() != 0 {
		t.Errorf("expected heap's length to be 0, got %d", heap.Len())
	}

	prev := math.MinInt
	for i, v := range slice {
		if v < prev {
			t.Errorf("Slice is not sorted at position %d", i)
		} else {
			prev = v
		}
	}
}
