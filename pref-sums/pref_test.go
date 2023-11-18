package prefsums

import "testing"

func TestHasSubArray(t *testing.T) {
	arr := []int{1, 4, 5, 2, -1, 2, 3}
	prefs := GetPrefixSums(arr)

	if !HasSubArray(prefs, 11) {
		t.Errorf("expected array %v to contain subarray with sum %d", arr, 11)
	}
	if !HasSubArray(prefs, 6) {
		t.Errorf("expected array %v to contain subarray with sum %d", arr, 6)
	}
	if HasSubArray(prefs, 17) {
		t.Errorf("expected array %v to not contain subarray with sum %d", arr, 17)
	}
}
