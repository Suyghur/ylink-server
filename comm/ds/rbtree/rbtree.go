//@File     rbtree.go
//@Time     2022/05/12
//@Author   #Suyghur,

package rbtree

import (
	"fmt"
	"ylink/comm/utils/comparator"
	"ylink/comm/utils/visitor"
)

var defaultKeyComparator = comparator.BuiltinTypeComparator

type Options struct {
	keyCmp comparator.Comparator
}

// Option is a function type used to set Options
type Option func(option *Options)

//WithKeyComparator is used to set RbTree's key comparator
func WithKeyComparator(cmp comparator.Comparator) Option {
	return func(option *Options) {
		option.keyCmp = cmp
	}
}

type RbTree struct {
	root   *Node
	size   int
	keyCmp comparator.Comparator
}

// New creates a new RbTree
func New(opts ...Option) *RbTree {
	option := Options{
		keyCmp: defaultKeyComparator,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &RbTree{keyCmp: option.keyCmp}
}

// Clear clears the RbTree
func (tree *RbTree) Clear() {
	tree.root = nil
	tree.size = 0
}

// Find finds the first node that the key is equal to the passed key, and returns its value
func (tree *RbTree) Find(key interface{}) interface{} {
	n := tree.findFirstNode(key)
	if n != nil {
		return n.value
	}
	return nil
}

// FindNode the first node that the key is equal to the passed key and return it
func (tree *RbTree) FindNode(key interface{}) *Node {
	return tree.findFirstNode(key)
}

// Begin returns the node with minimum key in the RbTree
func (tree *RbTree) Begin() *Node {
	return tree.First()
}

// First returns the node with minimum key in the RbTree
func (tree *RbTree) First() *Node {
	if tree.root == nil {
		return nil
	}
	return minimum(tree.root)
}

// RBegin returns the Node with maximum key in the RbTree
func (tree *RbTree) RBegin() *Node {
	return tree.Last()
}

// Last returns the Node with maximum key in the RbTree
func (tree *RbTree) Last() *Node {
	if tree.root == nil {
		return nil
	}
	return maximum(tree.root)
}

// IterFirst returns the iterator of first node
func (tree *RbTree) IterFirst() *RbTreeIterator {
	return NewIterator(tree.First())
}

// IterLast returns the iterator of first node
func (tree *RbTree) IterLast() *RbTreeIterator {
	return NewIterator(tree.Last())
}

// Empty returns true if Tree is empty,otherwise returns false.
func (tree *RbTree) Empty() bool {
	if tree.size == 0 {
		return true
	}
	return false
}

// Size returns the size of the rbtree.
func (tree *RbTree) Size() int {
	return tree.size
}

// Insert inserts a key-value pair into the RbTree.
func (tree *RbTree) Insert(key, value interface{}) {
	x := tree.root
	var y *Node

	for x != nil {
		y = x
		if tree.keyCmp(key, x.key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node{parent: y, color: RED, key: key, value: value}
	tree.size++

	if y == nil {
		z.color = BLACK
		tree.root = z
		return
	} else if tree.keyCmp(z.key, y.key) < 0 {
		y.left = z
	} else {
		y.right = z
	}
	tree.rbInsertFixup(z)
}

func (tree *RbTree) rbInsertFixup(z *Node) {
	var y *Node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					tree.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					tree.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.leftRotate(z.parent.parent)
			}
		}
	}
	tree.root.color = BLACK
}

// Delete deletes node from the RbTree
func (tree *RbTree) Delete(node *Node) {
	z := node
	if z == nil {
		return
	}

	var x, y *Node
	if z.left != nil && z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	xparent := y.parent
	if x != nil {
		x.parent = xparent
	}
	if y.parent == nil {
		tree.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.key = y.key
		z.value = y.value
	}

	if y.color == BLACK {
		tree.rbDeleteFixup(x, xparent)
	}
	tree.size--
}

func (tree *RbTree) rbDeleteFixup(x, parent *Node) {
	var w *Node
	for x != tree.root && getColor(x) == BLACK {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			x, w = tree.rbFixupLeft(x, parent, w)
		} else {
			x, w = tree.rbFixupRight(x, parent, w)
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (tree *RbTree) rbFixupLeft(x, parent, w *Node) (*Node, *Node) {
	w = parent.right
	if w.color == RED {
		w.color = BLACK
		parent.color = RED
		tree.leftRotate(parent)
		w = parent.right
	}
	if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
		w.color = RED
		x = parent
	} else {
		if getColor(w.right) == BLACK {
			if w.left != nil {
				w.left.color = BLACK
			}
			w.color = RED
			tree.rightRotate(w)
			w = parent.right
		}
		w.color = parent.color
		parent.color = BLACK
		if w.right != nil {
			w.right.color = BLACK
		}
		tree.leftRotate(parent)
		x = tree.root
	}
	return x, w
}

func (tree *RbTree) rbFixupRight(x, parent, w *Node) (*Node, *Node) {
	w = parent.left
	if w.color == RED {
		w.color = BLACK
		parent.color = RED
		tree.rightRotate(parent)
		w = parent.left
	}
	if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
		w.color = RED
		x = parent
	} else {
		if getColor(w.left) == BLACK {
			if w.right != nil {
				w.right.color = BLACK
			}
			w.color = RED
			tree.leftRotate(w)
			w = parent.left
		}
		w.color = parent.color
		parent.color = BLACK
		if w.left != nil {
			w.left.color = BLACK
		}
		tree.rightRotate(parent)
		x = tree.root
	}
	return x, w
}

func (tree *RbTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (tree *RbTree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// findNode finds the node that its key is equal to the passed key, and returns it.
func (tree *RbTree) findNode(key interface{}) *Node {
	x := tree.root
	for x != nil {
		if tree.keyCmp(key, x.key) < 0 {
			x = x.left
		} else {
			if tree.keyCmp(key, x.key) == 0 {
				return x
			}
			x = x.right
		}
	}
	return nil
}

// findNode finds the first node that its key is equal to the passed key, and returns it
func (tree *RbTree) findFirstNode(key interface{}) *Node {
	node := tree.FindLowerBoundNode(key)
	if node == nil {
		return nil
	}
	if tree.keyCmp(node.key, key) == 0 {
		return node
	}
	return nil
}

// FindLowerBoundNode finds the first node that its key is equal or greater than the passed key, and returns it
func (tree *RbTree) FindLowerBoundNode(key interface{}) *Node {
	return tree.findLowerBoundNode(tree.root, key)
}

func (tree *RbTree) findLowerBoundNode(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}
	if tree.keyCmp(key, x.key) <= 0 {
		ret := tree.findLowerBoundNode(x.left, key)
		if ret == nil {
			return x
		}
		if tree.keyCmp(ret.key, x.key) <= 0 {
			return ret
		}
		return x
	}
	return tree.findLowerBoundNode(x.right, key)
}

// FindUpperBoundNode finds the first node that its key is greater than the passed key, and returns it
func (tree *RbTree) FindUpperBoundNode(key interface{}) *Node {
	return tree.findUpperBoundNode(tree.root, key)
}

func (tree *RbTree) findUpperBoundNode(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}
	if tree.keyCmp(key, x.key) >= 0 {
		return tree.findUpperBoundNode(x.right, key)
	}
	ret := tree.findUpperBoundNode(x.left, key)
	if ret == nil {
		return x
	}
	if tree.keyCmp(ret.key, x.key) <= 0 {
		return ret
	}
	return x
}

// Traversal traversals elements in the RbTree, it will not stop until to the end of RbTree or the visitor returns false
func (tree *RbTree) Traversal(visitor visitor.KvVisitor) {
	for node := tree.First(); node != nil; node = node.Next() {
		if !visitor(node.key, node.value) {
			break
		}
	}
}

// IsRbTree is a function use to test whether t is a RbTree or not
func (tree *RbTree) IsRbTree() (bool, error) {
	// Properties:
	// 1. Each node is either red or black.
	// 2. The root is black.
	// 3. All leaves (NIL) are black.
	// 4. If a node is red, then both its children are black.
	// 5. Every path from a given node to any of its descendant NIL nodes contains the same number of black nodes.
	_, property, ok := tree.test(tree.root)
	if !ok {
		return false, fmt.Errorf("violate property %v", property)
	}
	return true, nil
}

func (tree *RbTree) test(n *Node) (int, int, bool) {

	if n == nil { // property 3:
		return 1, 0, true
	}

	if n == tree.root && n.color != BLACK { // property 2:
		return 1, 2, false
	}
	leftBlackCount, property, ok := tree.test(n.left)
	if !ok {
		return leftBlackCount, property, ok
	}
	rightBlackCount, property, ok := tree.test(n.right)
	if !ok {
		return rightBlackCount, property, ok
	}

	if rightBlackCount != leftBlackCount { // property 5:
		return leftBlackCount, 5, false
	}
	blackCount := leftBlackCount

	if n.color == RED {
		if getColor(n.left) != BLACK || getColor(n.right) != BLACK { // property 4:
			return 0, 4, false
		}
	} else {
		blackCount++
	}

	if n == tree.root {
		//fmt.Printf("blackCount:%v \n", blackCount)
	}
	return blackCount, 0, true
}

// getColor returns the node's color
func getColor(n *Node) Color {
	if n == nil {
		return BLACK
	}
	return n.color
}
