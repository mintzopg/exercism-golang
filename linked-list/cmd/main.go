package main

import (
	"linkedlist"
)

func main() {
	l1 := linkedlist.NewList(1, 2, 3)
	l1.Reverse()
	l1.PrintList()

}
