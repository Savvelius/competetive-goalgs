package segmtree

import (
	"fmt"
	"log"
)

type Vertex struct {
	parent *Vertex
	left   *Vertex
	right  *Vertex

	sum    int
	length int

	assign   int
	assigned bool
}

// Constructs segment tree from array in O(N) time and memory
func New(array []int) *Vertex {
	switch len(array) {
	case 0:
		return nil
	case 1:
		return &Vertex{sum: array[0], length: 1}
	default:
		root := &Vertex{length: len(array)}
		mid := len(array) / 2

		root.left = New(array[:mid])
		root.sum += root.left.sum
		root.left.parent = root

		root.right = New(array[mid:])
		root.sum += root.right.sum
		root.right.parent = root

		return root
	}
}

func (root *Vertex) isLeaf() bool {
	return root.left == nil
}

func (root *Vertex) Add(idx, value int) {
	root.tryPushAssign()
	root.sum += value

	if root.isLeaf() {
		return
	}

	if idx < root.left.length {
		root.left.Add(idx, value)
	} else {
		root.right.Add(idx-root.left.length, value)
	}
}

func (root *Vertex) Set(idx, value int) {
	if idx >= root.length {
		panic("idx should be less than length")
	}

	root.tryPushAssign()
	if root.isLeaf() {
		delta := value - root.sum
		root.sum = value

		cur := root
		for cur.parent != nil {
			cur.parent.sum += delta
			cur = cur.parent
		}
		return
	}

	if idx < root.left.length {
		root.left.Set(idx, value)
	} else {
		root.right.Set(idx-root.left.length, value)
	}
}

func (root *Vertex) ForLeftRight(f func(*Vertex)) {
	if !root.isLeaf() {
		f(root.left)
		f(root.right)
	}
}

// assigns value to all elements in [l, r)
func (root *Vertex) Assign(l, r, value int) {
	if l == 0 && r == root.length {
		root.assign = value
		root.assigned = true
		return
	}

	if root.isLeaf() {
		if l == r && l == 0 {
			root.assign = value
			root.assigned = true
			return
		}
		panic("sub array range out of bounds")
	}

	if r <= root.left.length {
		root.left.Assign(l, r, value)
	} else if l >= root.left.length {
		root.right.Assign(l-root.left.length, r-root.left.length, value)
	} else {
		root.left.Assign(l, root.left.length, value)
		root.right.Assign(0, r, value)
	}
}

// if root is assigned
func (root *Vertex) tryPushAssign() {
	if !root.assigned {
		return
	}
	root.sum = root.length * root.assign

	root.ForLeftRight(func(v *Vertex) {
		v.assign = root.assign
		v.assigned = true
	})
	root.assigned = false
}

// returns sum in [l, r)
func (root *Vertex) GetSum(l, r int) int {
	if l > r {
		log.Fatalf("l=%d must less or equal to r=%d", l, r)
	}

	if l == r || l > root.length {
		return 0
	}

	root.tryPushAssign()
	if l == 0 && r == root.length {
		return root.sum
	}

	if root.isLeaf() {
		panic("sub array range out of bounds")
	}

	return root.left.GetSum(l, min(r, root.left.length)) + root.right.GetSum(0, r-root.left.length)
}

func (root *Vertex) Pprint() {
	lvl := 1
	lvlSize := 1
	curLvlSize := 0
	switched := true

	queue := []*Vertex{root}

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

		fmt.Printf("Node(sum=%d, len=%d)  ", cur.sum, cur.length)
		cur.ForLeftRight(func(v *Vertex) {
			queue = append(queue, v)
		})
	}
}

// DFS from left to right
func (root *Vertex) ForEach(f func(*Vertex)) {
	f(root)

	root.ForLeftRight(func(v *Vertex) {
		v.ForEach(f)
	})
}
