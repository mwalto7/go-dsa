package tree

type BSTree struct {
	root *Node
}

func NewBST() *BSTree {
	return &BSTree{nil}
}

func (b *BSTree) Root() *Node {
	return b.root
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

func (b *BSTree) InOrder(f func(int)) {
	inOrder(b.root, f)
}

func (b *BSTree) PreOrder(f func(int)) {
	preOrder(b.root, f)
}

func (b *BSTree) PostOrder(f func(int)) {
	postOrder(b.root, f)
}

func (b *BSTree) LevelOrder(f func(int)) {
	levelOrder(b.root, f)
}
