package main

import (
	"github.com/mwalto7/go-dsa/tree"
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 0, 9, 5}

	bst := tree.NewBST()
	bst.Insert(values...)

	preorder := tree.Slice(bst.PreOrder)
	fmt.Println(preorder)

	inorder := tree.Slice(bst.InOrder)
	fmt.Println(inorder)

	postorder := tree.Slice(bst.PostOrder)
	fmt.Println(postorder)

	tree.Walk(bst.InOrder)
}
