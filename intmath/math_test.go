package intmath

import "testing"

func TestNextPow2(t *testing.T) {
	if NextPow2(5) != 8 {
		t.Errorf("expected NextPow2(5)==%d, got %d", 8, NextPow2(5))
	}

	if NextPow2(8) != 8 {
		t.Errorf("expected NextPow2(8)==%d, got %d", 8, NextPow2(8))
	}
}
