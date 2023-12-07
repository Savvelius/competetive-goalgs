package bellford

import (
	"fmt"
	"testing"
)

func TestSearch(test *testing.T) {
	g := New[string]()

	s := g.Emplace("s")
	v := g.Emplace("v")
	t := g.Emplace("t")
	w := g.Emplace("w")
	u := g.Emplace("u")

	g.Connect(s, v, 4)
	g.Connect(s, u, 2)
	g.Connect(u, v, -1)
	g.Connect(u, w, 2)
	g.Connect(v, t, 4)
	g.Connect(w, t, 2)

	paths := g.MinPaths(s)
	for v, cost := range paths {
		fmt.Printf("Vertice(%v). Path=%f\n", v.value, cost)
	}
}
