package main

import (
	"binarysearchtree"
)

func main() {
	bst := binarysearchtree.SearchTreeData{}
	bst.Insert(3)

	binarysearchtree.PrintTree(&bst)
}
