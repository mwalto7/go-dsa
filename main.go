package main

import (
	"github.com/mwalto7/go-dsa/tree"
	"fmt"
)

func main() {
	bst := tree.NewBST()
	values := []int{1, 2, 3, 4, 0, 9, 5}
	bst.Insert(values...)

	pre := tree.Slice(bst.PreOrder)
	fmt.Println(pre)
}
