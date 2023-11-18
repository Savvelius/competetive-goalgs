package trie

import "testing"

func TestTrie(t *testing.T) {
	root := NewVertex()
	root.AddString("Hello")
	root.AddString("Hola")

	if !root.Contains("Hello") {
		t.Errorf("Expected trie to contain %s", "Hello")
	}

	if !root.Contains("Hola") {
		t.Errorf("Expected trie to contain %s", "Hola")
	}

	if root.Contains("Halo") {
		t.Errorf("Expected trie to not contain %s", "Halo")
	}

	root.Erase("Hola")
	if root.Contains("Hola") {
		t.Errorf("Expected trie to not contain %s", "Hola")
	}
}
