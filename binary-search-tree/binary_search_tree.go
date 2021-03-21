package binarysearchtree

import (
	"fmt"
)

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// func Bst(int) *SearchTreeData
func Bst(v int) *SearchTreeData {
	return &SearchTreeData{data: v}
}

// func (*SearchTreeData) Insert(int)
func (t *SearchTreeData) Insert(v int) {
	t = insert(t, v)
}

// func (*SearchTreeData) MapString(func(int) string) []string
func (t *SearchTreeData) MapString(f func(int) string) []string {
	s := []string{}
	traverseString(t, f, &s)
	return s
}

// func (*SearchTreeData) MapInt(func(int) int) []int
func (t *SearchTreeData) MapInt(f func(int) int) []int {
	a := []int{}
	traverseInt(t, f, &a)
	return a
}

func insert(t *SearchTreeData, v int) *SearchTreeData {
	if t == nil {
		return &SearchTreeData{nil, v, nil}
	}
	// All data in the left subtree is less than or equal to the
	// ... current node's data, and all data in the right subtree is greater than
	// ... the current node's data.
	if v <= t.data {
		t.left = insert(t.left, v)
		return t
	}
	t.right = insert(t.right, v)
	return t
}

func traverseString(t *SearchTreeData, f func(int) string, s *[]string) {
	if t == nil {
		return
	}
	traverseString(t.left, f, s)
	*s = append(*s, f(t.data))
	traverseString(t.right, f, s)
}

func traverseInt(t *SearchTreeData, f func(int) int, s *[]int) {
	if t == nil {
		return
	}
	traverseInt(t.left, f, s)
	*s = append(*s, f(t.data))
	traverseInt(t.right, f, s)
}

// helper func
func PrintTree(t *SearchTreeData) {
	if t == nil {
		return
	}
	PrintTree(t.left)
	fmt.Println(t.data, " ")
	PrintTree(t.right)
}
