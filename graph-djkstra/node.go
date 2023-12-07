package graph

import (
	"fmt"
	"math"

	"github.com/Savvelius/graph-algs/heap"
)

/*
	visited := map[*Node[T]]uint{}
	keys := map[*Node[T]]uint{}
	keys[root] = 0

	heap := heap.New[*Node[T]](func(n1, n2 *Node[T]) bool {
		return keys[n1] < keys[n2]
	}, 16)

	root.ForEach(func(n *Node[T]) {
		if n != root {
			keys[n] = math.MaxUint
		}
		heap.Insert(n)
	}, map[*Node[T]]struct{}{})

	for heap.Len() != 0 {
		// invariant - keys[w] is smallest edge from X to heap
		w := heap.ExtractMin()
		visited[w] = keys[w]

		// preserve the invariant
		for y, path := range w.adjacent {
			keys[y] = min(keys[y], visited[w]+path)
		}
	}

	return keys
*/

type Node[T comparable] struct {
	value    T
	adjacent map[*Node[T]]uint
}

type Edge[T comparable] struct {
	n1   *Node[T]
	n2   *Node[T]
	cost uint
}

func (edge *Edge[T]) reverse() Edge[T] {
	return Edge[T]{
		n1:   edge.n2,
		n2:   edge.n1,
		cost: edge.cost,
	}
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{value: value, adjacent: map[*Node[T]]uint{}}
}

// panics on negative dist or else djkstra wouldn't work
func (node *Node[T]) Add(other *Node[T], dist uint) {
	if _, ok := node.adjacent[other]; ok {
		return
	}

	node.adjacent[other] = dist
}

func (node *Node[T]) Emplace(value T, dist uint) *Node[T] {
	emplaced := NewNode[T](value)
	node.Add(emplaced, dist)
	return emplaced
}

func (node *Node[T]) forEach(f func(*Node[T]), visited map[*Node[T]]struct{}) {
	f(node)
	visited[node] = struct{}{}
	for adj := range node.adjacent {
		if _, isVisited := visited[adj]; isVisited {
			continue
		}
		adj.forEach(f, visited)
	}
}

func (node *Node[T]) ForEach(f func(*Node[T])) {
	node.forEach(f, map[*Node[T]]struct{}{})
}

func (node *Node[T]) DFS(value T) *Node[T] {
	if node.value == value {
		return node
	}

	for adj := range node.adjacent {
		if found := adj.DFS(value); found != nil {
			return found
		}
	}

	return nil
}

func (node *Node[T]) BFS(value T) *Node[T] {
	queue := []*Node[T]{node}

	for len(queue) != 0 {

		current := queue[0]
		if current.value == value {
			return current
		}
		queue = queue[1:]

		for adj := range current.adjacent {
			queue = append(queue, adj)
		}
	}

	return nil
}

func (root *Node[T]) topoSort(visited map[*Node[T]]struct{}, slice []T) {
	visited[root] = struct{}{}
	slice = append(slice, root.value)
	for node := range root.adjacent {
		if _, isVisited := visited[node]; isVisited {
			continue
		}
		visited[node] = struct{}{}
		slice = append(slice, node.value)
		node.topoSort(visited, slice)
	}
}

func (root *Node[T]) TopoSort() []T {
	slice := []T{}
	root.topoSort(map[*Node[T]]struct{}{}, slice)
	return slice
}

// O(n*log(n)) time with heap's assistance
func (root *Node[T]) Djkstra() map[*Node[T]]uint {
	// visited := map[*Node[T]]uint{}
	// visited[root] = 0
	keys := map[*Node[T]]uint{} // invariant for edges from visited to unvisited
	keys[root] = 0

	heap := heap.New[*Node[T]](func(n1, n2 *Node[T]) bool { return keys[n1] < keys[n2] }, 16)
	root.ForEach(func(n *Node[T]) {
		if n != root {
			keys[n] = math.MaxUint
		}
		heap.Insert(n)
	})

	for heap.Len() != 0 {
		start := heap.ExtractMin()
		startLen := keys[start]
		// visited[start] = startLen

		// update invariant
		for end, path := range start.adjacent {
			keys[end] = min(keys[end], startLen+path)
		}
	}

	return keys
}

// BFS
func (node *Node[T]) Pprint() {
	visited := map[*Node[T]]struct{}{}
	queue := []*Node[T]{node}

	for len(queue) != 0 {
		popped := queue[0]
		queue = queue[1:]
		if _, isVisited := visited[popped]; isVisited {
			continue
		}
		visited[popped] = struct{}{}

		fmt.Printf("Node(value=%v).Children:\n", popped.value)

		for adj, len := range popped.adjacent {
			fmt.Printf("Node(value=%v, len=%d)  ", adj.value, len)
			queue = append(queue, adj)
		}
		fmt.Print("\n------------------------------------------------\n")
	}
}
