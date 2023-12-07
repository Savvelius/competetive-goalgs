package bitarray

import "testing"

func assert(t *testing.T, expr bool) {
	if !expr {
		t.Fail()
	}
}

func TestArray(t *testing.T) {
	arr := New(10)
	assert(t, arr.Len() == 0)
	arr.Append(true)
	assert(t, arr.Len() == 1)
	arr.Append(false)
	assert(t, arr.Len() == 2)
	arr.Append(true)
	assert(t, arr.Len() == 3)
	assert(t, arr.Get(0) == true)

	assert(t, arr.Get(1) == false)
	arr.Set(1, true)
	assert(t, arr.Get(1) == true)

	assert(t, arr.Get(2) == true)
}
