package array

import (
	"fmt"

	"github.com/Savvelius/competetive-goalgs/intmath"
)

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

type Vertex[T any] struct {
	value T
	l, r  int
}

func (v *Vertex[T]) len() int {
	return v.r - v.l
}

func (v *Vertex[T]) isLeaf() bool {
	return v.len() == 1
}

type Tree[T any] struct {
	array   []Vertex[T]
	binOp   func(T, T) T // caller guarantees it's associativity
	neutral T

	// for self-iteration
	curIdx int
}

func (tree *Tree[T]) outOfBounds(i int) bool {
	return i >= len(tree.array)
}

func (tree *Tree[T]) at(i int) *Vertex[T] {
	// may explode
	if tree.outOfBounds(i) {
		return nil
	}
	return &tree.array[i]
}

func (tree *Tree[T]) left(i int) *Vertex[T] {
	return tree.at(left(i))
}

func (tree *Tree[T]) right(i int) *Vertex[T] {
	return tree.at(right(i))
}

func (tree *Tree[T]) isLeaf(i int) bool {
	return tree.at(i).len() == 1
}

func (tree *Tree[T]) updateValue(idx int) {
	if tree.isLeaf(idx) {
		return
	}

	tree.at(idx).value = tree.binOp(tree.left(idx).value, tree.right(idx).value)
}

func (tree *Tree[T]) build(array []T, curidx int) {
	cur := tree.at(curidx)

	if cur.len() == 1 {
		cur.value = array[cur.l]
		return
	}

	mid := cur.len() / 2

	*tree.left(curidx) = Vertex[T]{l: cur.l, r: cur.l + mid}
	*tree.right(curidx) = Vertex[T]{l: cur.l + mid, r: cur.r}

	tree.build(array, left(curidx))
	tree.build(array, right(curidx))

	tree.updateValue(curidx)
}

func New[T any](array []T, binOp func(T, T) T, neutral T) Tree[T] {
	tree := Tree[T]{
		array:   make([]Vertex[T], intmath.NextPow2(len(array))*2-1),
		binOp:   binOp,
		neutral: neutral,
		curIdx:  0,
	}

	tree.array[0] = Vertex[T]{l: 0, r: len(array)}
	tree.build(array, 0)

	return tree
}

func (tree *Tree[T]) Query(l, r int) T {
	curV := tree.at(tree.curIdx)

	if r <= curV.l || l >= curV.r {
		return tree.neutral
	}

	if curV.isLeaf() || l == curV.l && r == curV.r {
		return curV.value
	}

	savedIdx := tree.curIdx

	tree.curIdx = left(savedIdx)
	leftValue := tree.Query(l, r)

	tree.curIdx = right(savedIdx)
	rightValue := tree.Query(l, r)

	tree.curIdx = savedIdx

	return tree.binOp(leftValue, rightValue)
}

// applies binOp with given value to all elements in given range
func (tree *Tree[T]) Apply(l, r int, value T) {
	curV := tree.at(tree.curIdx)

	if r <= curV.l || l >= curV.r {
		return
	}

	if curV.isLeaf() {
		curV.value = tree.binOp(curV.value, value)
		return
	}

	savedIdx := tree.curIdx

	tree.curIdx = left(savedIdx)
	tree.Apply(l, r, value)

	tree.curIdx = right(savedIdx)
	tree.Apply(l, r, value)

	tree.curIdx = savedIdx
	tree.updateValue(savedIdx)
}

// sets all elements in range [l, r) to given value
func (tree *Tree[T]) Set(l, r int, value T) {
	curV := tree.at(tree.curIdx)

	if r <= curV.l || l >= curV.r {
		return
	}

	if curV.isLeaf() {
		curV.value = value
		return
	}

	savedIdx := tree.curIdx

	tree.curIdx = left(savedIdx)
	tree.Set(l, r, value)

	tree.curIdx = right(savedIdx)
	tree.Set(l, r, value)

	tree.curIdx = savedIdx
	tree.updateValue(savedIdx)
}

func (tree *Tree[T]) Len() int {
	return tree.array[0].len()
}

func (tree *Tree[T]) Pprint() {
	lvl := 1
	printed := false
	count := 0
	queue := []int{0}

	for len(queue) != 0 {
		if !printed {
			fmt.Printf("Level %d\n", lvl)
			printed = true
		}

		cur := queue[0]
		queue = queue[1:]
		count++

		curVertex := tree.at(cur)
		fmt.Printf("V(value=%v, l=%d, r=%d),  ", curVertex.value, curVertex.l, curVertex.r)

		if curVertex.isLeaf() {
			continue
		}

		queue = append(queue, left(cur))
		queue = append(queue, right(cur))

		if count == (1<<lvl)-1 {
			lvl++
			printed = false
			fmt.Print("\n----------------------------\n")
		}
	}
}
