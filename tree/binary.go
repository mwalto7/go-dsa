package tree

import (
	"fmt"
	"os"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

type BSTree struct {
	root *Node
}

func NewBST() *BSTree {
	return &BSTree{nil}
}

// Insert inserts the specified values into the tree
func (b *BSTree) Insert(values ...int) {
	for _, value := range values {
		b.root = insert(b.root, value)
	}
}

// insert recursively inserts a new node with the specified value into the tree
// and returns the new node
func insert(n *Node, value int) *Node {
	if n == nil {
		return NewNode(value)
	}
	if n.value == value {
		return n
	}
	if n.value > value {
		n.left = insert(n.left, value)
	} else {
		n.right = insert(n.right, value)
	}
	return n
}

// TraverseFunc traverses the tree in the specified order then applies function f at each node.
// Valid traversal orders include in-order, pre-order, post-order, and level-order.
func (b *BSTree) TraverseFunc(order string, f func(int)) {
	switch order {
	case "inorder":
		inOrder(b.root, f)
	case "preorder":
		preOrder(b.root, f)
	case "postorder":
		postOrder(b.root, f)
	case "levelorder":
		levelOrder(b.root, f)
	default:
		fmt.Fprintln(os.Stderr, "tree: unkown traversal order")
	}
}

// inOrder recursively traverses the tree in order (least to greatest)
// and applies func f at each node
func inOrder(n *Node, f func(int)) {
	if n != nil {
		inOrder(n.left, f)
		f(n.value)
		inOrder(n.right, f)
	}
}

// preOrder recursively traverses the tree in pre-order and applies func
// f at each node
func preOrder(n *Node, f func(int)) {
	if n != nil {
		f(n.value)
		preOrder(n.left, f)
		preOrder(n.right, f)
	}
}

// postOrder recursively traverses the tree in reverse (greatest to least)
// and applies func f at each node
func postOrder(n *Node, f func(int)) {
	if n != nil {
		postOrder(n.right, f)
		f(n.value)
		postOrder(n.left, f)
	}
}

// levelOrder traverses the tree level-by-level and applies func f at each node
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
