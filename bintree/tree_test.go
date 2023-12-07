package bintree

import (
	"runtime"
	"testing"
)

func TestTree(t *testing.T) {
	// 4 2 6
	tree := NewNode[int, string](3, "root")
	tree.InsertOrUpdate(5, "five")

	testEqual(t, tree.Get(10), nil)
	testEqual(t, tree.Get(5).Val, "five")
	testEqual(t, tree.Get(3).Val, "root")

	tree.InsertOrUpdate(4, "four")
	tree.InsertOrUpdate(2, "two")
	tree.InsertOrUpdate(6, "six")

	testEqual(t, tree.Get(10), nil)
	testEqual(t, tree.Get(5).Val, "five")
	testEqual(t, tree.Get(3).Val, "root")

	tree.Delete(5)
	testEqual(t, tree.Get(5), nil)
	testEqual(t, tree.Get(6).Val, "six")
	testEqual(t, tree.Get(4).Val, "four")
	testEqual(t, tree.Get(3).Val, "root")
	testEqual(t, tree.Get(2).Val, "two")

	// 4 2 6 left
	tree.Delete(3)
	testEqual(t, tree.Get(6).Val, "six")
	testEqual(t, tree.Get(3), nil)
	testEqual(t, tree.Get(2).Val, "two")
	testEqual(t, tree.Get(4).Val, "four")

	tree.InsertOrUpdate(4, "FOUR")
	testEqual(t, tree.Get(6).Val, "six")
	testEqual(t, tree.Get(3), nil)
	testEqual(t, tree.Get(2).Val, "two")
	testEqual(t, tree.Get(4).Val, "FOUR")

	testEqual(t, tree.Get(6).Predecessor().Val, "FOUR")
	testEqual(t, tree.Get(2).Predecessor(), nil)
	testEqual(t, tree.Get(4).Predecessor().Val, "two")
}

func testEqual[T comparable](t *testing.T, got, expected T) {
	if got != expected {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("Test line %d: Expected %v, got %v.", line, expected, got)
	}
}
