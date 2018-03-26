package tree

import (
	"fmt"
	"errors"
)

var MissingValueError = errors.New("tree: value not in tree")

// Node represents a tree node.
type Node struct {
	left   *Node // a node's left child
	right  *Node // a node's right child
	parent *Node // a node's parent node
	Value  int   // the value of a node
}

// NewNode returns a new node with the specified value.
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// Left returns a node's left child node.
func (n *Node) Left() *Node {
	return n.left
}

// Right returns a node's right child node.
func (n *Node) Right() *Node {
	return n.right
}

// Parent returns a node's parent node.
func (n *Node) Parent() *Node {
	return n.parent
}

// String returns the string representation of a node.
func (n *Node) String() string {
	if n != nil {
		return fmt.Sprintf("%d", n.Value)
	}
	return "<nil>"
}

// Tree represents a tree ADT
type Tree interface {
	// Search searches a tree for a node containing the specified value.
	Search(value int) *Node

	// Insert inserts the specified values into a tree.
	Insert(values ...int)

	// Delete deletes the specified nodes from a tree.
	Delete(nodes ...*Node) error
}

// preOrder recursively traverses a tree in pre-order and applies func
// f at each node.
func preOrder(n *Node, f func(*Node)) {
	if n != nil {
		f(n)
		preOrder(n.left, f)
		preOrder(n.right, f)
	}
}

// inOrder recursively traverses a tree in order (least to greatest)
// and applies func f at each node.
func inOrder(n *Node, f func(*Node)) {
	if n != nil {
		inOrder(n.left, f)
		f(n)
		inOrder(n.right, f)
	}
}

// postOrder recursively traverses a tree in reverse (greatest to least)
// and applies func f at each node.
func postOrder(n *Node, f func(*Node)) {
	if n != nil {
		postOrder(n.right, f)
		f(n)
		postOrder(n.left, f)
	}
}

// levelOrder traverses a tree level-by-level and applies func f at each node.
func levelOrder(n *Node, f func(*Node)) {
	if n != nil {
		var queue []*Node
		queue = append(queue, n)
		for {
			nodeCount := len(queue)
			if nodeCount == 0 {
				return
			}
			for nodeCount > 0 {
				tmp := queue[0]
				f(tmp)
				queue[0] = nil
				queue = queue[1:]
				if tmp.left != nil {
					queue = append(queue, tmp.left)
				}
				if tmp.right != nil {
					queue = append(queue, tmp.right)
				}
				nodeCount--
			}
			fmt.Println()
		}
	}
}

// TraversalFunc defines a function that is applied at each node during tree traversal.
type TraversalFunc func(f func(*Node))

// Traverse prints a tree's values in the specified traversal order.
func Walk(order TraversalFunc) {
	order(func(n *Node) {
		fmt.Println(n)
	})
}

// TraverseFunc traverses a tree in the specified order and applies function f
// at each node.
func WalkFunc(order TraversalFunc, f func(*Node)) {
	order(f)
}

// Slice traverses a tree in the specified order and appends each node's
// value to a slice.
func Slice(order TraversalFunc) []*Node {
	var nodes []*Node
	order(func(n *Node) {
		nodes = append(nodes, n)
	})
	return nodes
}
