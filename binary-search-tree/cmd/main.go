package main

import (
	"binarysearchtree"
	"strconv"
)

func main() {
	bst := binarysearchtree.Bst(4)
	bst.MapString(strconv.Itoa)

	binarysearchtree.PrintTree(bst)
}
