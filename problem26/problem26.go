package main

import (
    "fmt"
)

/*
Given a singly linked list and an integer k, remove the kth last element from the list. k is guaranteed to be smaller than the length of the list.

The list is very long, so making more than one pass is prohibitively expensive.

Do this in constant space and in one pass.
*/

type Node struct {
	Next *Node
	Value string
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) addNode(v string) {
	new_node := &Node{Value:v}
	if l.Head == nil {
		l.Head = new_node
	}
	if l.Tail != nil {
		l.Tail.Next = new_node
	}
	l.Tail = new_node
	l.Size++
}

func (l *LinkedList) print() string {
	s := ""
	n := l.Head
	s += n.Value
	for n.Next != nil {
		n = n.Next
		s += " -> " + n.Value
	}
	return s
}

func (l *LinkedList) removeLastK(k int) {
	size := l.Size
	n := l.Head
	for i := 0; i < (size-k-1); i++ {
		n = n.Next
	}
	n.Next = nil
}

func main() {
	l := &LinkedList{Size:0}
	l.addNode("1")
	l.addNode("2")
	l.addNode("3")
	l.addNode("4")
	l.addNode("5")
	l.addNode("6")
	s := l.print()
	fmt.Println(s)
	l.removeLastK(3)
	s = l.print()
	fmt.Println(s)
}