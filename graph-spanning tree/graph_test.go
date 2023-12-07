package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := New("root")
	g0 := g.Emplace("node00", 3)
	g1 := g.Emplace("node01", -2)

	g1.Connect(g0, 10)
	g00 := g0.Emplace("nodeg00", 18)
	g00.Connect(g1, 5)

	g10 := g1.Emplace("nodeg10", -8)
	g10.Connect(g, 9)

	spanned := g.KruskalUF()
	for edge := range spanned {
		fmt.Printf("Edge(from=%v, to=%v, cost=%d)\n", edge.v1.value, edge.v2.value, edge.cost)
	}

	print("\n-----------------------\n")

	spanned = g.PrimRewrite()
	for edge := range spanned {
		fmt.Printf("Edge(from=%v, to=%v, cost=%d)\n", edge.v1.value, edge.v2.value, edge.cost)
	}

	// g.Pprint()
	// fmt.Print("\n\nOLD:\n")
	// g.Pprint()
}
