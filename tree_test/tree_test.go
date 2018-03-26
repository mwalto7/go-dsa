package tree_test

import (
	"testing"
	"github.com/mwalto7/go-dsa/tree"
)

var nodes = []*tree.Node{
	{Value: 2},
	{Value: 1},
	{Value: 3},
	{Value: -1},
	{Value: 0},
	{Value: 1},
	{Value: 9},
}

func TestWalk(t *testing.T) {
	bst := tree.NewBST()
	bst.Insert(2, 3, 1)
	tree.Walk(bst.InOrder)
}
