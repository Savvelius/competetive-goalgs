package array

import (
	"testing"
)

func TestSeq(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 3, 6, 7}
	a2 := []int{4, 3, 6, 7, 1, 2, 3, 10}

	v := MCS(a1, a2)
	if v != 4 {
		t.Errorf("expected mcs of a1, a2 to be %d, got %d", 4, v)
	}

	a1 = []int{3, 1, 2, 3, 4, 5, 0, -1, 5, 10, 12, 12, 13}

	v = MCIS(a1)
	if v != 8 {
		t.Errorf("expected mcis of a1 to be %d, got %d", 8, v)
	}
}
