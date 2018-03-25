package tree

type AVLTree struct {
	root *Node
}

func NewAVL() *AVLTree {
	return &AVLTree{nil}
}

func (a *AVLTree) Root() *Node {
	return a.root
}

func (a *AVLTree) Insert(values ...int) {
	//for _, value := range values {
	//	  insert(a.root, value)
	//}
}
