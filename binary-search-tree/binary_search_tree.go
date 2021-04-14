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
<<<<<<< HEAD
	// t = insert(t, v)
=======
>>>>>>> master
	if v <= t.data {
		if t.left == nil {
			t.left = Bst(v)
		} else {
			t.left.Insert(v)
		}
	}
	if v > t.data {
		if t.right == nil {
			t.right = Bst(v)
		} else {
			t.right.Insert(v)
		}
	}
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
