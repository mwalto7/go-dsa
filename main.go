package main

import (
	"github.com/mwalto7/go-dsa/tree"
	"fmt"
)

func main() {
	bst := tree.NewBST()
	bst.Insert(2, 1, 3)
	values := tree.Slice(bst.PostOrder)
	fmt.Println(values)

	tree.TraverseFunc(bst.InOrder, func(value int) {
		fmt.Println(value)
	})
}
