package tree

// BSTree represents a binary search tree.
type BSTree struct {
	root *Node
}

// NewBST constructs an empty binary search tree.
func NewBST() *BSTree {
	return &BSTree{nil}
}

// Root returns the root node of the binary search tree.
func (b *BSTree) Root() *Node {
	return b.root
}

// Insert inserts the specified values into the tree
func (b *BSTree) Insert(values ...int) {
	for _, value := range values {
		b.insert(NewNode(value))
	}
}

// insert inserts a new node into the tree and keeps track of the parent node
func (b *BSTree) insert(newNode *Node) {
	var parent *Node
	currNode := b.root
	for currNode != nil {
		parent = currNode
		if newNode.value < currNode.value {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}
	newNode.parent = parent
	if parent == nil {
		b.root = newNode
	} else if newNode.value < parent.value {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}

// InOrder traverses the tree in order and applies function f at each node.
func (b *BSTree) InOrder(f func(int)) {
	inOrder(b.root, f)
}

// PreOrder traverses the tree in pre-order and applies function f at each node.
func (b *BSTree) PreOrder(f func(int)) {
	preOrder(b.root, f)
}

// PostOrder traverses the tree in post-order and applies function f at each node.
func (b *BSTree) PostOrder(f func(int)) {
	postOrder(b.root, f)
}

// LevelOrder traverses the tree in level-order and applies function f at each node.
func (b *BSTree) LevelOrder(f func(int)) {
	levelOrder(b.root, f)
}
