package tree

type RBTree struct {
	root *Node
}

func NewRBTree() *RBTree {
	return &RBTree{nil}
}

func (r *RBTree) Root() *Node {
	return r.root
}

func (r *RBTree) Insert(values ...int) {

}

func (r *RBTree) insert(n *Node, value int) {

}
