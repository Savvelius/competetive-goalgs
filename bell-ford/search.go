package bellford

import "math"

type Vertice[T any] struct {
	value T
	// edgesOut map[*Vertice[T]]float64
	edgesIn map[*Vertice[T]]float64
}

func newV[T any](value T) *Vertice[T] {
	return &Vertice[T]{
		value:   value,
		edgesIn: map[*Vertice[T]]float64{},
	}
}

// returns (nil, ...) if v has no edges in
// func (v *Vertice[T]) minEdgeIn() (*Vertice[T], float64) {
// 	var minV *Vertice[T]
// 	minCost := math.Inf(1)

// 	for from, cost := range v.edgesIn {
// 		if cost < minCost {
// 			minV = from
// 			minCost = cost
// 		}
// 	}

// 	return minV, minCost
// }

type Graph[T any] struct {
	vertices map[*Vertice[T]]int // maps each vertice into it's index in vArray
	vArray   []*Vertice[T]
}

func New[T any]() Graph[T] {
	return Graph[T]{
		vertices: map[*Vertice[T]]int{},
		vArray:   []*Vertice[T]{},
	}
}

func (graph *Graph[T]) Contains(vertice *Vertice[T]) bool {
	_, contains := graph.vertices[vertice]
	return contains
}

func (graph *Graph[T]) Add(vertice *Vertice[T]) {
	if graph.Contains(vertice) {
		panic("trying to add a contained vertice")
	}

	graph.vertices[vertice] = len(graph.vArray)
	graph.vArray = append(graph.vArray, vertice)
}

func (graph *Graph[T]) Emplace(value T) *Vertice[T] {
	vertice := newV(value)
	graph.Add(vertice)
	return vertice
}

func (graph *Graph[T]) Connect(from, to *Vertice[T], cost float64) {
	if !(graph.Contains(from) && graph.Contains(to)) {
		panic("graph doesn't contain vertice you are trying to connect")
	}

	_, contained := to.edgesIn[from]
	if contained {
		panic("trying to connect already connected vertice")
	}

	to.edgesIn[from] = cost
}

// returns nil on negative cycle
func (graph *Graph[T]) MinPaths(source *Vertice[T]) map[*Vertice[T]]float64 {
	if !graph.Contains(source) {
		panic("source is not in graph")
	}

	// L(i, v): i in [0, n], #v = len(graph.vArray)
	L := make([][]float64, len(graph.vArray)+1)
	for i := range L {
		L[i] = make([]float64, len(graph.vArray))
	}
	for i, v := range graph.vArray {
		if v != source {
			L[0][i] = math.Inf(1)
		}
	}
	for i := 1; i <= len(graph.vArray); i++ {
		stable := true

		for v, verticeTo := range graph.vArray {
			if verticeTo == source {
				continue
			}

			minVpath := math.Inf(1)

			for w, cost := range verticeTo.edgesIn {

				wIdx := graph.vertices[w]
				wPath := L[i-1][wIdx]

				if cost+wPath < minVpath {
					minVpath = cost + wPath
				}
			}

			L[i][v] = min(L[i-1][v], minVpath)

			if L[i][v] != L[i-1][v] {
				stable = false
			}
		}

		if stable {
			solution := make(map[*Vertice[T]]float64, len(graph.vArray))
			for v, vertice := range graph.vArray {
				solution[vertice] = L[i-1][v]
			}
			return solution
		}
	}

	return nil
}
