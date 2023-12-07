package bintree

import "cmp"

type Node[K cmp.Ordered, V any] struct {
	Key    K
	Val    V
	parent *Node[K, V]
	less   *Node[K, V]
	more   *Node[K, V]
}

func NewNode[K cmp.Ordered, V any](key K, val V) *Node[K, V] {
	return &Node[K, V]{
		Key: key,
		Val: val,
	}
}

func (node *Node[K, V]) isRoot() bool {
	return node.parent == nil
}

func (node *Node[K, V]) numChildren() int {
	cnt := 0
	if node.less != nil {
		cnt++
	}
	if node.more != nil {
		cnt++
	}
	return cnt
}

func (node *Node[K, V]) isLessChild() bool {
	if node.parent == nil {
		return false
	}
	return node.parent.less == node
}

// Applies f to less and more child if they are not nil.
// Second parameter - is less
func (node *Node[K, V]) forChild(f func(*Node[K, V], bool)) {
	if node.less != nil {
		f(node.less, true)
	}
	if node.more != nil {
		f(node.more, false)
	}
}

func (node *Node[K, V]) Min() *Node[K, V] {
	if node.less == nil {
		return node
	}
	return node.less.Min()
}

func (node *Node[K, V]) Max() *Node[K, V] {
	if node.more == nil {
		return node
	}
	return node.more.Max()
}

func (node *Node[K, V]) Get(key K) *Node[K, V] {
	switch cmp.Compare(key, node.Key) {
	case 0:
		return node
	case -1: // less
		if node.less == nil {
			return nil
		}
		return node.less.Get(key)
	default: //more
		if node.more == nil {
			return nil
		}
		return node.more.Get(key)
	}
}

func (node *Node[K, V]) Predecessor() *Node[K, V] {
	if node.less != nil {
		return node.less.Max()
	}

	for node.parent != nil {
		// parent is less
		if node == node.parent.more {
			return node.parent
		}
		node = node.parent
	}

	return nil
}

func (node *Node[K, V]) InsertOrUpdate(key K, val V) {
	switch cmp.Compare(key, node.Key) {
	case 0:
		node.Val = val
	case -1: // less
		if node.less != nil {
			node.less.InsertOrUpdate(key, val)
			return
		}
		node.less = NewNode(key, val)
		node.less.parent = node
	default: //more
		if node.more != nil {
			node.more.InsertOrUpdate(key, val)
			return
		}
		node.more = NewNode(key, val)
		node.more.parent = node
	}
}

// TODO:
func (node *Node[K, V]) Delete(key K) {
	found := node.Get(key)
	if found == nil {
		return
	}

	switch found.numChildren() {
	case 0:
		if node.isRoot() {
			*node = Node[K, V]{}
			return
		}
		node.forChild(func(n *Node[K, V], isLess bool) {
			if isLess {
				n.parent.less = nil
			} else {
				n.parent.more = nil
			}
		})
	case 1:
		found.forChild(func(foundCh *Node[K, V], _ bool) {
			if found.isRoot() {
				*found = *foundCh
				found.forChild(func(n *Node[K, V], _ bool) {
					n.parent = found
				})
				return
			}

			foundCh.parent = found.parent
			if found.isLessChild() {
				found.parent.less = foundCh
			} else {
				found.parent.more = foundCh
			}
		})
	case 2:
		leaf := found.Predecessor()
		if leaf == nil {
			panic("cannot happen - Predecessor of node with two children is nil")
		}
		// swap their keys
		leaf.Key, found.Key = found.Key, leaf.Key
		// swap their values
		leaf.Val, found.Val = found.Val, leaf.Val
		// delete the leaf
		leaf.Delete(key)
	}
}
