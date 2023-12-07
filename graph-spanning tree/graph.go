package graph

import (
	"cmp"
	"fmt"
	"math"
	"slices"

	// "github.com/Savveius/gogreedyalgs/sort"
	unionfind "github.com/Savvelius/competetive-goalgs/union-find"
	"github.com/Savvelius/graph-algs/heap"
)

// Contains Prim and Kruskal
// min spanning tree algorithms algorithms

type Edge[T any] struct {
	cost int
	v1   *Vertice[T]
	v2   *Vertice[T]
}

type Vertice[T any] struct {
	value T
	edges map[*Edge[T]]struct{}
}

func New[T any](value T) *Vertice[T] {
	return &Vertice[T]{
		value: value,
		edges: map[*Edge[T]]struct{}{},
	}
}

func newEdge[T any](cost int, v1 *Vertice[T], v2 *Vertice[T]) *Edge[T] {
	return &Edge[T]{
		cost: cost,
		v1:   v1,
		v2:   v2,
	}
}

// panics if vertice already has that edge
func (v1 *Vertice[T]) addEdge(edge *Edge[T]) {
	if _, hasEdge := v1.edges[edge]; hasEdge {
		panic("vertice already has this edge")
	}
	v1.edges[edge] = struct{}{}
}

func (v1 *Vertice[T]) Connect(v2 *Vertice[T], cost int) {
	edge := newEdge(cost, v1, v2)
	v1.addEdge(edge)
	v2.addEdge(edge)
}

func (v1 *Vertice[T]) Emplace(value T, cost int) *Vertice[T] {
	v2 := New(value)
	v1.Connect(v2, cost)
	return v2
}

// panics if than not in edge
func (edge *Edge[T]) otherThan(than *Vertice[T]) *Vertice[T] {
	if edge.v1 == than {
		return edge.v2
	} else if edge.v2 == than {
		return edge.v1
	} else {
		panic("edge doesn't have given node")
	}
}

// not correct. FIXME
func (startV *Vertice[T]) Prim() map[*Edge[T]]struct{} {
	visited := map[*Vertice[T]]struct{}{}

	keys := map[*Vertice[T]]int{}
	keys[startV] = 0

	edges := map[*Edge[T]]struct{}{}

	heap := heap.New[*Vertice[T]](func(v1, v2 *Vertice[T]) bool {
		return keys[v1] < keys[v2]
	}, 16)
	startV.ForEachVertice(func(v *Vertice[T]) {
		if v != startV {
			keys[v] = math.MaxInt
		}
		heap.Insert(v)
	})

	for heap.Len() != 0 {
		minV := heap.ExtractMin()
		visited[minV] = struct{}{}

		for edge := range minV.edges {
			end := edge.otherThan(minV)
			// invariant: keys contains edges from X to V - X
			if _, isVisited := visited[end]; isVisited && (edge.cost == keys[minV]) {
				edges[edge] = struct{}{}
			}
			keys[end] = min(keys[end], edge.cost)
		}
	}

	return edges
}

func (root *Vertice[T]) Kruskal() map[*Edge[T]]struct{} {
	edges := []*Edge[T]{}
	root.ForEachEdge(func(e *Edge[T]) {
		edges = append(edges, e)
	})

	slices.SortFunc(edges, func(e1, e2 *Edge[T]) int {
		return cmp.Compare(e1.cost, e2.cost)
	})

	visitedEdges := map[*Edge[T]]struct{}{}

	// iterate through edges in ascending order
	for _, edge := range edges {
		if !edge.v1.isConnected(edge.v2, visitedEdges) {
			visitedEdges[edge] = struct{}{}
		}
	}

	return visitedEdges
}

func (root *Vertice[T]) KruskalUF() map[*Edge[T]]struct{} {
	edges := []*Edge[T]{}
	root.ForEachEdge(func(e *Edge[T]) {
		edges = append(edges, e)
	})

	slices.SortFunc(edges, func(e1, e2 *Edge[T]) int {
		return cmp.Compare(e1.cost, e2.cost)
	})

	vertices := []*Vertice[T]{}
	root.ForEachVertice(func(v *Vertice[T]) {
		vertices = append(vertices, v)
	})
	visitedEdges := map[*Edge[T]]struct{}{}

	connectedVertices := unionfind.New(vertices)

	// iterate through edges in increasing order
	for _, edge := range edges {
		if connectedVertices.Find(edge.v1) != connectedVertices.Find(edge.v2) {
			connectedVertices.Union(edge.v1, edge.v2)
			visitedEdges[edge] = struct{}{}
		}
	}

	return visitedEdges
}

func (root *Vertice[T]) PrimRewrite() map[*Edge[T]]struct{} {
	rootEdges := map[*Vertice[T]]*Edge[T]{}
	for edge := range root.edges {
		end := edge.otherThan(root)
		rootEdges[end] = edge
	}

	key := map[*Vertice[T]]int{}
	winner := map[*Vertice[T]]*Edge[T]{}
	heap := heap.New[*Vertice[T]](func(v1, v2 *Vertice[T]) bool {
		return key[v1] < key[v2]
	}, 16)
	root.ForEachVertice(func(v *Vertice[T]) {
		if v == root {
			return
		}

		heap.Insert(v)
		if rootEdge, hasRootEdge := rootEdges[v]; hasRootEdge {
			key[v] = rootEdge.cost
			winner[v] = rootEdge
		} else {
			key[v] = math.MaxInt
			winner[v] = nil
		}
	})
	visitedEdges := map[*Edge[T]]struct{}{}
	for heap.Len() != 0 {
		minV := heap.ExtractMin()
		visitedEdges[winner[minV]] = struct{}{}

		for edge := range minV.edges {
			end := edge.otherThan(minV)
			if edge.cost < key[end] {
				key[end] = edge.cost
				winner[end] = edge
			}
		}
	}

	return visitedEdges
}

func (v *Vertice[T]) isConnected(needle *Vertice[T], haystack map[*Edge[T]]struct{}) bool {
	if v == needle {
		return true
	}

	queue := []*Vertice[T]{v}
	visited := map[*Vertice[T]]struct{}{}

	for len(queue) != 0 {
		popped := queue[0]
		queue = queue[1:]
		if _, isVisited := visited[popped]; isVisited {
			continue
		}
		visited[popped] = struct{}{}

		for edge := range popped.edges {
			if _, inEdges := haystack[edge]; inEdges {
				end := edge.otherThan(popped)
				if end == needle {
					return true
				}
				queue = append(queue, end)
			}
		}
	}

	return false
}

func (v *Vertice[T]) ForEachVertice(f func(*Vertice[T])) {
	v.forEachVertice(f, map[*Vertice[T]]struct{}{})
}

func (v *Vertice[T]) forEachVertice(f func(*Vertice[T]), visited map[*Vertice[T]]struct{}) {
	f(v)
	visited[v] = struct{}{}
	for edge := range v.edges {
		end := edge.otherThan(v)
		if _, isVisited := visited[end]; isVisited {
			continue
		}
		visited[end] = struct{}{}
		end.forEachVertice(f, visited)
	}
}

func (v *Vertice[T]) ForEachEdge(f func(*Edge[T])) {
	v.forEachEdge(f, map[*Edge[T]]struct{}{})
}

func (v *Vertice[T]) forEachEdge(f func(*Edge[T]), visited map[*Edge[T]]struct{}) {
	for edge := range v.edges {
		if _, isVisited := visited[edge]; isVisited {
			continue
		}
		visited[edge] = struct{}{}
		f(edge)
		edge.otherThan(v).forEachEdge(f, visited)
	}
}

func (startV *Vertice[T]) Pprint() {
	queue := []*Vertice[T]{startV}
	visited := map[*Vertice[T]]struct{}{}

	for len(queue) != 0 {
		popped := queue[0]
		queue = queue[1:]
		if _, isVisited := visited[popped]; isVisited {
			continue
		}
		visited[popped] = struct{}{}

		fmt.Printf("Vertice(value=%v).Children:\n", popped.value)
		for edge := range popped.edges {
			poppedEnd := edge.otherThan(popped)
			queue = append(queue, poppedEnd)
			fmt.Printf("Vertice(value=%v, cost=%d)  ", poppedEnd.value, edge.cost)
		}
		fmt.Print("\n------------------------------------------------------------------------------\n")
	}
}
