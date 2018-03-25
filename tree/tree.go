package tree

import (
	"fmt"
)

// Node represents a tree node.
type Node struct {
	value  int   // the value of a node
	left   *Node // a node's left child
	right  *Node // a node's right child
	parent *Node // a node's parent node
}

// NewNode returns a new node with the specified value.
func NewNode(value int) *Node {
	return &Node{value, nil, nil, nil}
}

// Tree represents a tree ADT
type Tree interface {
	// Root returns the root node of a tree.
	Root() *Node

	// Insert inserts the specified values into a tree.
	Insert(values ...int)
}

// Traverse prints a tree's values in the specified traversal order.
func Walk(order TraversalFunc) {
	order(func(value int) {
		fmt.Println(value)
	})
}

// Slice traverses a tree in the specified order and appends each node's
// value to a slice.
func Slice(order TraversalFunc) []int {
	var values []int
	order(func(value int) {
		values = append(values, value)
	})
	return values
}

// TraverseFunc traverses a tree in the specified order and applies function f
// at each node.
func WalkFunc(order TraversalFunc, f func(int)) {
	order(f)
}

// TraversalFunc defines a function that is applied at each node during tree traversal.
type TraversalFunc func(f func(int))

// inOrder recursively traverses a tree in order (least to greatest)
// and applies func f at each node.
func inOrder(n *Node, f func(int)) {
	if n != nil {
		inOrder(n.left, f)
		f(n.value)
		inOrder(n.right, f)
	}
}

// preOrder recursively traverses a tree in pre-order and applies func
// f at each node.
func preOrder(n *Node, f func(int)) {
	if n != nil {
		f(n.value)
		preOrder(n.left, f)
		preOrder(n.right, f)
	}
}

// postOrder recursively traverses a tree in reverse (greatest to least)
// and applies func f at each node.
func postOrder(n *Node, f func(int)) {
	if n != nil {
		postOrder(n.right, f)
		f(n.value)
		postOrder(n.left, f)
	}
}

// levelOrder traverses a tree level-by-level and applies func f at each node.
func levelOrder(n *Node, f func(int)) {
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
				f(tmp.value)
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
