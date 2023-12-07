package huff

import (
	"fmt"
	"math"

	"github.com/Savveius/gogreedyalgs/sort"
	"github.com/Savvelius/graph-algs/heap"
)

type Node struct {
	label string
	path  string
	prob  float64
	zero  *Node
	one   *Node
}

type Leaf struct {
	label string
	prob  float64
}

type Input struct {
	freq  uint
	label string
}

func inputsToLeafs(inputs []Input) []Leaf {
	sumFreq := float64(0)
	for i := range inputs {
		sumFreq += float64(inputs[i].freq)
	}

	leafs := sort.Map(inputs, func(i Input) Leaf {
		return Leaf{label: i.label, prob: float64(i.freq) / sumFreq}
	})

	return leafs
}

func NewNode(prob float64, label string, zero *Node, one *Node) *Node {
	if prob < 0 || prob > 1 {
		panic("probability should be in bounds [0, 1]")
	}
	return &Node{
		label: label,
		prob:  prob,
		zero:  zero,
		one:   one,
	}
}

func NewLeafNode(l Leaf) *Node {
	return NewNode(l.prob, l.label, nil, nil)
}

func (zero *Node) Merge(one *Node) *Node {
	return NewNode(zero.prob+one.prob, zero.label+one.label, zero, one)
}

func (root *Node) setPaths(path string) {
	if root.zero != nil {
		zeroPath := path + "0"
		root.zero.path += zeroPath
		root.zero.setPaths(zeroPath)
	}
	if root.one != nil {
		onePath := path + "1"
		root.one.path += onePath
		root.one.setPaths(onePath)
	}
}

func build(nodes heap.Heap[*Node]) *Node {
	if nodes.Len() == 0 {
		panic("shouldn't happen in correct implementation")
	}

	min1 := nodes.ExtractMin()

	if nodes.Len() == 0 {
		return min1
	}

	min2 := nodes.ExtractMin()

	merged := min1.Merge(min2)
	nodes.Insert(merged)

	return build(nodes)
}

func Build(inputs []Input) *Node {

	if len(inputs) == 0 {
		return nil
	}

	leafs := inputsToLeafs(inputs)

	sumProb := float64(0)

	nodes := heap.New[*Node](func(n1, n2 *Node) bool { return n1.prob < n2.prob }, len(inputs))
	for i := range leafs {
		sumProb += leafs[i].prob

		newNode := NewLeafNode(leafs[i])
		nodes.Insert(newNode)
	}

	f64EqThreshold := 1e-8
	if sumProb > 1+f64EqThreshold || sumProb < 1-f64EqThreshold {
		panic("sum of leaf probabilities should be 1")
	}

	tree := build(nodes)
	tree.setPaths("")
	return tree
}

func (root *Node) Pprint() {
	fmt.Printf("Node(label=%s, path=%s, prob=%f)\n", root.label, root.path, root.prob)

	if root.zero != nil {
		root.zero.Pprint()
	}

	if root.one != nil {
		root.one.Pprint()
	}
}

func (node *Node) isLeaf() bool {
	return (node.one == nil) && (node.zero == nil)
}

func (root *Node) addLeafs(leafs map[string]float64) {
	if root.isLeaf() {
		leafs[root.path] = root.prob
		return
	}

	if root.one != nil {
		root.one.addLeafs(leafs)
	}

	if root.zero != nil {
		root.zero.addLeafs(leafs)
	}
}

func (root *Node) GetLanguage() map[string]float64 {
	leafs := map[string]float64{}
	root.addLeafs(leafs)
	return leafs
}

func (root *Node) LanguageInformation() float64 {
	leafs := root.GetLanguage()

	information := float64(0)
	for k, v := range leafs {
		information += v * float64(len(k))
	}

	return information
}

func (root *Node) LanguageEntropy() float64 {
	leafs := root.GetLanguage()

	entropy := float64(0)
	for _, v := range leafs {
		entropy += -v * math.Log2(v)
	}

	return entropy
}
