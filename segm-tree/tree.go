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
		if root.left != nil {
			root.sum += root.left.sum
			root.left.parent = root
		}

		root.right = New(array[mid:])
		if root.right != nil {
			root.sum += root.right.sum
			root.right.parent = root
		}

		return root
	}
}

func (root *Vertex) isLeaf() bool {
	return root.left == nil && root.right == nil
}

func (root *Vertex) Add(idx, value int) {

}

func (root *Vertex) Set(idx, value int) {
	if idx >= root.length {
		panic("idx should be less than length")
	}

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

	if root.left != nil && idx < root.left.length {
		root.left.Set(idx, value)
		return
	}

	offset := root.length - root.right.length
	root.right.Set(idx-offset, value)
}

// returns sum in [l, r)
func (root *Vertex) GetSum(l, r int) int {
	if l > r {
		log.Fatalf("l=%d must less or equal to r=%d", l, r)
	}

	if l == r {
		return 0
	}

	if l == 0 && r == root.length {
		return root.sum
	}

	if root.isLeaf() {
		panic("sub array range out of bounds")
	}

	childSum := 0

	if root.left != nil && l < root.left.length {
		childSum += root.left.GetSum(l, min(r, root.left.length))
	}

	leftLen := root.length - root.right.length
	childSum += root.right.GetSum(0, r-leftLen)

	return childSum
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
		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
}

// DFS from left to right
func (root *Vertex) ForEach(f func(*Vertex)) {
	f(root)

	if root.left != nil {
		root.left.ForEach(f)
	}

	if root.right != nil {
		root.right.ForEach(f)
	}
}
