package main

import (
	"github.com/mwalto7/go-dsa/tree"
	"fmt"
)

func main() {
	bst := tree.NewBST()
	bst.Insert([]int{2, 1, 3}...)
	bst.TraverseFunc("levelorder", func(data int) {
		fmt.Printf("%d ", data)
	})
}
