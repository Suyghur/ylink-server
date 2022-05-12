//@File     node.go
//@Time     2022/05/12
//@Author   #Suyghur,

package rbtree

type Color bool

const (
	RED   = false
	BLACK = true
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  Color
	key    interface{}
	value  interface{}
}

func (node *Node) Key() interface{} {
	return node.key
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) SetValue(value interface{}) {
	node.value = value
}

func (node *Node) Next() *Node {
	return successor(node)
}

func (node *Node) Prev() *Node {
	return preSuccessor(node)
}

// successor returns the successor of the Node
func successor(node *Node) *Node {
	if node.right != nil {
		return minimum(node.right)
	}
	y := node.parent
	for y != nil && node == y.right {
		node = y
		y = node.parent
	}
	return y
}

// preSuccessor returns the preSuccessor of the Node
func preSuccessor(node *Node) *Node {
	if node.left != nil {
		return maximum(node.left)
	}
	if node.parent != nil {
		if node.parent.right == node {
			return node.parent
		}
		for node.parent != nil && node.parent.left == node {
			node = node.parent
		}
		return node.parent
	}
	return nil
}

// minimum finds the minimum Node of subtree n.
func minimum(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

// maximum finds the maximum Node of subtree n.
func maximum(node *Node) *Node {
	for node.right != nil {
		node = node.right
	}
	return node
}
