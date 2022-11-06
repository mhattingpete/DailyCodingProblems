package main

import (
    "fmt"
    "example.com/linkedList"
)

/*
Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical.

For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return the node with value 8.

In this example, assume nodes with the same value are the exact same node objects.

Do this in O(M + N) time (where M and N are the lengths of the lists) and constant space.
*/

func findIntersection(l1 *linkedList.LinkedList, l2 *linkedList.LinkedList) int {
	head1 := l1.First
	head2 := l2.First
	ptr1 := head1
	ptr2 := head2
	if ptr1 == nil || ptr2 == nil {
		return -1
	} else {
		for ptr1.Value != ptr2.Value {
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
			// if intersection is found
			if ptr1 != nil && ptr2 != nil { 
				if ptr1.Value == ptr2.Value {
					return ptr1.Value
				}
			}
			// if we reach end of one of the lists
			if ptr1 == nil {
				ptr1 = head2
			}
			if ptr2 == nil {
				ptr2 = head1
			}
		}
		return ptr1.Value
	}
}

func main() {
	l1 := linkedList.LinkedList{}
	l1.InsertLast(3)
	l1.InsertLast(7)
	l1.InsertLast(8)
	l1.InsertLast(10)
	l1.Print()
	l2 := linkedList.LinkedList{}
	l2.InsertLast(99)
	l2.InsertLast(98)
	l2.InsertLast(1)
	l2.InsertLast(8)
	l2.InsertLast(10)
	l2.Print()
	v := findIntersection(&l1, &l2)
	fmt.Printf("Intersection found with: %d", v)

}