package generic

import (
	"fmt"
)

/*
Separate Vertex struct from Tree struct
Got to do this in rust i guess.
Given : [l, r) which is in range - checked by Tree,
	query all segments which combine to [l, r)
1. Partial overlap -> query(left) and query(right)
2. No overlap - return neutral element
3. Total overlap - return value at this node
*/

// create for every 	monoid
type Vertex[T any] struct {
	left, right, parent *Vertex[T]

	l, r  int
	value T
}

func (v *Vertex[T]) len() int {
	return v.r - v.l
}

func (v *Vertex[T]) isLeaf() bool {
	return v.left == nil
}

func (v *Vertex[T]) forLeftRight(f func(*Vertex[T])) {
	if !v.isLeaf() {
		f(v.left)
		f(v.right)
	}
}

func (cur *Vertex[T]) buildVertex(array []T, binOp func(T, T) T) {
	if cur.len() == 1 {
		cur.value = array[cur.l]
		return
	}

	mid := cur.len() / 2

	cur.left = &Vertex[T]{l: cur.l, r: cur.l + mid, parent: cur}
	cur.right = &Vertex[T]{l: cur.l + mid, r: cur.r, parent: cur}

	cur.forLeftRight(func(v *Vertex[T]) {
		v.buildVertex(array, binOp)
	})

	cur.value = binOp(cur.left.value, cur.right.value)
}

const (
	PARTIAL_OVERLAP = iota
	NO_OVERLAP
	TOTAL_OVERLAP
)

// to be rewritten. probably...
func (v *Vertex[T]) overlap(l, r int) int {
	if l <= v.l && r >= v.r {
		return TOTAL_OVERLAP
	}

	if l >= v.r || r <= v.l {
		return NO_OVERLAP
	}

	return PARTIAL_OVERLAP
}

type Tree[T any] struct {
	root *Vertex[T]

	binOp   func(T, T) T
	neutral T

	// this is for recursive calls
	curV *Vertex[T]
}

func New[T any](binOp func(T, T) T, neutral T, array []T) Tree[T] {
	tree := Tree[T]{
		root:    &Vertex[T]{l: 0, r: len(array)},
		binOp:   binOp,
		neutral: neutral,
	}
	tree.curV = tree.root

	tree.root.buildVertex(array, binOp)

	return tree
}

func (tree *Tree[T]) Len() int {
	return tree.root.len()
}

func (tree *Tree[T]) Query(l, r int) T {
	switch tree.curV.overlap(l, r) {
	case TOTAL_OVERLAP:
		return tree.curV.value
	case NO_OVERLAP:
		return tree.neutral
	case PARTIAL_OVERLAP:
		savedCurV := tree.curV

		tree.curV = savedCurV.left
		leftValue := tree.Query(l, r)

		tree.curV = savedCurV.right
		rightValue := tree.Query(l, r)

		tree.curV = tree.root

		return tree.binOp(leftValue, rightValue)
	default:
		panic("can't happen due to exhaustive switch")
	}
}

// applies binOp with given value to all elements in a sub array
func (tree *Tree[T]) Apply(l, r int, value T) {
	overlap := tree.curV.overlap(l, r)

	if overlap == NO_OVERLAP {
		return
	}

	if overlap == TOTAL_OVERLAP && tree.curV.isLeaf() {
		tree.curV.value = tree.binOp(tree.curV.value, value)
	} else {
		savedCurV := tree.curV

		tree.curV = savedCurV.left
		tree.Apply(l, r, value)

		tree.curV = savedCurV.right
		tree.Apply(l, r, value)

		tree.curV = tree.root

		savedCurV.value = tree.binOp(savedCurV.left.value, savedCurV.right.value)
	}
}

func (tree *Tree[T]) Pprint() {
	lvl := 1
	lvlSize := 1
	curLvlSize := 0
	switched := true

	queue := []*Vertex[T]{tree.root}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		curLvlSize++

		if switched {
			fmt.Printf("\nLevel %d:\n", lvl)
			switched = false
		}

		if curLvlSize == lvlSize {
			lvl++
			lvlSize *= 2
			curLvlSize = 0
			switched = true
		}

		fmt.Printf("Node(val=%v, l=%d, r=%d)  ", cur.value, cur.l, cur.r)
		cur.forLeftRight(func(v *Vertex[T]) {
			queue = append(queue, v)
		})
	}
}
