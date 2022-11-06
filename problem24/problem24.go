package main

import (
    "fmt"
)

/*
Implement locking in a binary tree. A binary tree node can be locked or unlocked only if all of its descendants or ancestors are not locked.

Design a binary tree node class with the following methods:

is_locked, which returns whether the node is locked
lock, which attempts to lock the node. If it cannot be locked, then it should return false. Otherwise, it should lock it and return true.
unlock, which unlocks the node. If it cannot be unlocked, then it should return false. Otherwise, it should unlock it and return true.

You may augment the node to add parent pointers or any other property you would like. 

You may assume the class is used in a single-threaded program, so there is no need for actual locks or mutexes. Each method should run in O(h), where h is the height of the tree.
*/

type Node struct {
	value string
	locked bool
	left *Node
	right *Node
}

func (t *Node) addLeftNode(v string) {
	n := &Node{value: v, locked: false}
	t.left = n
}

func (t *Node) addRightNode(v string) {
	n := &Node{value: v, locked: false}
	t.right = n
}

func (n *Node) is_locked() bool {
	return n.locked
}

func (t *Node) descendantsLocked() bool {
	if t != nil {
		if t.is_locked() == false {
			return t.left.descendantsLocked() || t.right.descendantsLocked()
		} else {
			return true
		}
	} else {
		return false
	}
}

func (root *Node) lock() bool {
	locked := root.left.descendantsLocked() || root.right.descendantsLocked()
	if locked == true {
		return false
	} else {
		root.locked = true
		return true
	}
}

func (root *Node) unlock() bool {
	locked := root.left.descendantsLocked() || root.right.descendantsLocked()
	if locked == true {
		return false
	} else {
		root.locked = false
		return true
	}
}

func main() {
	tree := &Node{value: "root", locked: false}
	tree.addLeftNode("left")
	tree.addRightNode("right")
	tree.left.addLeftNode("left.left")
	tree.left.addRightNode("left.right")
	tree.right.addLeftNode("right.left")
	tree.right.addRightNode("right.right")
	locked := tree.right.right.lock()
	fmt.Printf("Locked root.right.right: %v\n", locked)
	locked = tree.right.lock()
	fmt.Printf("Locked root.right: %v\n", locked)
	locked = tree.right.right.unlock()
	fmt.Printf("Unlocked root.right.right: %v\n", locked)
	locked = tree.right.lock()
	fmt.Printf("Locked root.right: %v\n", locked)
	locked = tree.right.unlock()
	fmt.Printf("Unlocked root.right: %v\n", locked)
	locked = tree.lock()
	fmt.Printf("Locked root: %v\n", locked)
}