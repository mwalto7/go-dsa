package tree

// BSTree represents a binary search tree.
type BSTree struct {
	root *Node
}

// NewBST constructs an empty binary search tree.
func NewBST() *BSTree {
	return &BSTree{nil}
}

func (b *BSTree) Search(value int) (*Node, error) {
	found := b.search(b.root, value)
	if found == nil {
		return nil, MissingValueError
	}
	return found, nil
}

func (b *BSTree) search(n *Node, value int) *Node {
	if n == nil || value == n.Value {
		return n
	}
	if value < n.Value {
		return b.search(n.left, value)
	} else {
		return b.search(n.right, value)
	}
}

// Insert inserts the specified values into the tree
func (b *BSTree) Insert(values ...int) {
	for _, v := range values {
		b.root = insert(b.root, v)
	}
}

// insert inserts a new node into the tree and keeps track of the parent node
func insert(n *Node, value int) *Node {
	if n == nil {
		return NewNode(value)
	}
	if value < n.Value {
		n.left = insert(n.left, value)
		n.left.parent = n
	} else {
		n.right = insert(n.right, value)
		n.right.parent = n
	}
	return n
}

func (b *BSTree) Delete(nodes ...*Node) error {
	for _, node := range nodes {
		if err := b.delete(b.root, node.Value); err != nil {
			return err
		}
	}
	return nil
}

func (b *BSTree) delete(n *Node, value int) error {

	return nil
}

// InOrder traverses the tree in order and applies function f at each node.
func (b *BSTree) InOrder(f func(*Node)) {
	inOrder(b.root, f)
}

// PreOrder traverses the tree in pre-order and applies function f at each node.
func (b *BSTree) PreOrder(f func(*Node)) {
	preOrder(b.root, f)
}

// PostOrder traverses the tree in post-order and applies function f at each node.
func (b *BSTree) PostOrder(f func(*Node)) {
	postOrder(b.root, f)
}

// LevelOrder traverses the tree in level-order and applies function f at each node.
func (b *BSTree) LevelOrder(f func(*Node)) {
	levelOrder(b.root, f)
}
