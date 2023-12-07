package wis

// Weigted Independent Set algorithm

type Vertice[T any] struct {
	value  T
	weight uint
}

func newVertice[T any](value T, weight uint) *Vertice[T] {
	return &Vertice[T]{
		value:  value,
		weight: weight,
	}
}

type Edge[T any] struct {
	v1 *Vertice[T]
	v2 *Vertice[T]
}

func newEdge[T any](v1, v2 *Vertice[T]) *Edge[T] {
	return &Edge[T]{
		v1: v1,
		v2: v2,
	}
}

func (edge *Edge[T]) otherThan(than *Vertice[T]) *Vertice[T] {
	if than == edge.v1 {
		return edge.v2
	} else if than == edge.v2 {
		return edge.v1
	} else {
		panic("edge doesn't contain 'than' vertice")
	}
}

type Graph[T any] struct {
	vertices map[*Vertice[T]]struct{}
	edges    map[*Edge[T]]struct{}

	vArray []*Vertice[T]
	eArray []*Edge[T]
}

func (graph *Graph[T]) HasVertice(v *Vertice[T]) bool {
	_, has := graph.vertices[v]
	return has
}

func (graph *Graph[T]) HasEdge(edge *Edge[T]) bool {
	_, has := graph.edges[edge]
	return has
}

func (graph *Graph[T]) AddVertice(v *Vertice[T]) {
	if graph.HasVertice(v) {
		panic("trying to insert existent vertice")
	}
	graph.vertices[v] = struct{}{}
	graph.vArray = append(graph.vArray, v)
}

func (graph *Graph[T]) AddEdge(edge *Edge[T]) {
	if graph.HasEdge(edge) {
		panic("trying to insert existent edge")
	}
	graph.edges[edge] = struct{}{}
	graph.eArray = append(graph.eArray, edge)
}

func (graph *Graph[T]) EmpalceVertice(value T, weight uint) *Vertice[T] {
	vertice := newVertice(value, weight)
	graph.AddVertice(vertice)
	return vertice
}

func (graph *Graph[T]) Connect(v1, v2 *Vertice[T]) *Edge[T] {
	if graph.FindEdge(v1, v2) != nil {
		panic("trying to connect already connected vertices")
	}
	if !graph.HasVertice(v1) || !graph.HasVertice(v2) {
		panic("trying to connect to non-present vertice")
	}
	edge := newEdge(v1, v2)
	graph.AddEdge(edge)
	return edge
}

func (graph *Graph[T]) FindVertice(value T, equal func(T, T) bool) *Vertice[T] {
	for _, vertice := range graph.vArray {
		if equal(vertice.value, value) {
			return vertice
		}
	}

	return nil
}

func (graph *Graph[T]) FindEdge(v1, v2 *Vertice[T]) *Edge[T] {
	for _, edge := range graph.eArray {
		areEqual := (edge.v1 == v1) && (edge.v2 == v2)
		areReverseEqual := (edge.v2 == v2) && (edge.v1 == v2)

		if areEqual || areReverseEqual {
			return edge
		}
	}

	return nil
}

func (graph *Graph[T]) WIS() (map[*Vertice[T]]struct{}, uint) {
	A := make([]uint, len(graph.vertices)+1)
	A[0] = 0
	A[1] = graph.vArray[0].weight

	for i, v := range graph.vArray[1:] {
		Aindex := i + 2
		A[Aindex] = max(A[Aindex-2]+v.weight, A[Aindex-1])
	}

	answer := map[*Vertice[T]]struct{}{}
	for i := len(A) - 1; i >= 2; i-- {
		curVertice := graph.vArray[i-1]
		if A[i-2]+curVertice.weight == A[i] {
			answer[curVertice] = struct{}{}
		}
	}

	return answer, A[len(A)-1]
}
