package graph

import (
	"fmt"
	"testing"
)

func TestSearch(tester *testing.T) {
	// setup
	s := NewNode[string]("s")
	v := s.Emplace("v", 1)
	w := s.Emplace("w", 4)
	v.Add(w, 2)
	t := v.Emplace("t", 6)
	w.Add(t, 3)

	// search
	tBfs := s.BFS("t")
	if tBfs == nil {
		tester.Errorf("BFS returned nil")
	}

	tDfs := s.DFS("t")
	if tDfs == nil {
		tester.Errorf("DFS returned nil")
	}

	if tDfs != tBfs {
		tester.Errorf("BFS returned: %v, DFS returned: %v", *tBfs, *tDfs)
	}

	// s.Pprint()
	shortest := s.Djkstra()

	for k, v := range shortest {
		fmt.Printf("Node(value=%v, score=%d)\n", k.value, v)
	}
}
