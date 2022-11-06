package main

import (
    "fmt"
)

// Node structure of a memory efficient doubly linked list
type xorNode struct {
	key int
	npx *xorNode
}

type xorLinkedList struct {
	head *xorNode
}

/* returns XORed value of the node addresses */
func xor(x *xorNode, y *xorNode) *xorNode {
	if x != nil && y != nil {
		return nil
	} else if x != nil && y == nil {
		return x
	} else if x == nil && y != nil {
		return y
	} else {
		return nil
	}
}

/* Insert a node at the begining of the XORed linked list and makes the
   newly inserted node as head */
func add(head_ref **xorNode, a int) {
	new_node := &xorNode{key: a}
	/* Since new node is being inserted at the begining, npx of new node
       will always be XOR of current head and NULL */
	new_node.npx = xor(*head_ref, nil)
	/* If linked list is not empty, then npx of current head node will be XOR 
       of new node and node next to current head */
	if *head_ref != nil {
		// *(head_ref)->npx is XOR of NULL and next. So if we do XOR of 
        // it with NULL, we get next
		next := xor((*head_ref).npx, nil)
		(*head_ref).npx = xor(new_node, next)
	}
	// Change head
	*head_ref = new_node
}

// prints contents of doubly linked list in forward direction
func display(head *xorNode) {
	var curr *xorNode = head
	var prev *xorNode = nil
	var next *xorNode
	for curr != nil {
		// get address of next node: curr->npx is next^prev, so curr->npx^prev
        // will be next^prev^prev which is next
		next = xor(prev, curr.npx)
		if next == nil {
			fmt.Printf("%+v", curr.key)
		} else {
			fmt.Printf("%+v -> ", curr.key)
		}
		// update prev and curr for next iteration
		prev = curr
		curr = next
	}
	fmt.Println()
}


func main() {
	var head *xorNode = nil
	add(&head,5)
	display(head)
	add(&head,9)
	display(head)
    add(&head,13)
    //add(&head,22)
    //add(&head,28)
    //add(&head,36)
    fmt.Println("\n==============================\n")
    display(head)
    fmt.Println("\n==============================\n")
}