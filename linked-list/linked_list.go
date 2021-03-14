package linkedlist

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

/*
Doubly Linked List (DLL); [pointer to prev][DATA][pointer to next] ...
						  <------------------ Node--------------->
e.g.	[nil][DATA][next] [prev][DATA][next] [prev][DATA][nil]
		 <---head-----> 				      <----Tail---->

		 ------ front----						---- back----
*/

// ErrEmptyList -  error definition for operations on empty list
var ErrEmptyList = errors.New("can't do that on an empty list")

// Node representation
type Node struct {
	Val        interface{}
	prev, next *Node
}

// List representation
type List struct {
	head   *Node
	length int
}

// Next returns the next node
func (n *Node) Next() *Node {
	if n != nil {
		return n.next
	}
	return nil
}

// Prev returns the previous node
func (n *Node) Prev() *Node {
	if n != nil {
		return n.prev
	}
	return nil
}

// First returns the first n in the list
func (n *Node) First() *Node {
	for n.prev != nil {
		n = n.prev
	}
	return n
}

// Last returns the last n in the list
func (n *Node) Last() *Node {
	for n.next != nil {
		n = n.next
	}
	return n
}

// NewList creates a new doubly linked list preserving the order of values
func NewList(args ...interface{}) *List {
	// first create an empty list
	l := &List{length: 0}
	// add all values given if any
	for _, v := range args {
		l.PushBack(v)
	}
	return l
}

// Len returns the list length. Helper method
func (l *List) Len() int {
	return l.length
}

// First returns a pointer to the head of the list
func (l *List) First() *Node {
	if l.length == 0 {
		return nil
	}
	return l.head
}

// Last returns a pointer to the tail of the list
func (l *List) Last() *Node {
	if l.length == 0 {
		return l.head
	}
	p := l.head
	for p.next != nil {
		p = p.next
	}
	return p
}

// PushBack inserts a value at the back of the list
func (l *List) PushBack(v interface{}) {
	n := &Node{Val: v}
	if l.length == 0 {
		l.head = n
		l.length++
	} else {
		n.prev = l.Last()
		l.Last().next = n
		l.length++
	}

}

// PopBack removes value at back
func (l *List) PopBack() (interface{}, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	if l.length == 1 {
		v := l.head.Val
		l.head = nil
		l.length = 0
		return v, nil
	}
	v := l.Last().Val
	l.Last().prev.next = nil
	l.length--
	return v, nil
}

// PushFront inserts value at front
func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v}
	if l.length == 0 {
		l.head = n
		l.length++
	} else {
		n.next = l.First()
		l.head.prev = n
		l.head = n
		l.length++
	}
}

// PopFront removes value at front
func (l *List) PopFront() (interface{}, error) {
	if l.length == 0 {
		return nil, ErrEmptyList
	}
	if l.length == 1 {
		v := l.head.Val
		l.head = nil
		l.length = 0
		return v, nil
	}
	v := l.head.Val
	l.First().next.prev = nil
	l.head = l.head.next
	l.length--
	return v, nil
}

// Reverse the list
func (l *List) Reverse() {
	if l.length == 0 {
		return
	}
	p := l.head
	for p.next != nil {
		p.prev, p.next = p.next, p.prev
		p = p.prev
	}
	p.prev, p.next = p.next, p.prev
	l.head = p
}

/*
******************** Helper functions  ******************
 */

// PrintList prints the nodes values in the list
func (l *List) PrintList() {
	p := l.head
	for p != nil {
		fmt.Printf("%+v ", p.Val)
		p = p.next
	}
	fmt.Println("List length = ", l.length)
}

// PrintNode prints a node's value
func (n *Node) PrintNode() {
	if n != nil {
		fmt.Println(n.Val)
	} else {
		log.Fatal(n)
	}
}

// PrintListDetails prints the linked list with both node's Val, next & prev pointers.
func (l *List) PrintListDetails() {
	buf := bytes.NewBuffer([]byte{'{'})
	buf.WriteString(fmt.Sprintf("First()= %p; ", l.First()))

	for cur := l.First(); cur != nil; cur = cur.Next() {
		buf.WriteString(fmt.Sprintf("[Prev()= %p, Val= %p (%v), Next()= %p] <-> ", cur.Prev(), cur, cur.Val, cur.Next()))
	}

	buf.WriteString(fmt.Sprintf("; Last()= %p; ", l.Last()))
	buf.WriteByte('}')

	fmt.Println(buf.String())
}
