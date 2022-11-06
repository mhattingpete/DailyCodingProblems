package main

import (
    "fmt"
    "example.com/doublyLinkedList"
)

/*
Given an array of integers and a number k, where 1 <= k <= length of the array, compute the maximum values of each subarray of length k.

For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should get: [10, 7, 8, 8], since:

10 = max(10, 5, 2)
7 = max(5, 2, 7)
8 = max(2, 7, 8)
8 = max(7, 8, 7)
Do this in O(n) time and O(k) space. You can modify the input array in-place and you do not need to store the results. You can simply print them out as you compute them.


*/

func maxInArr(arr []int) (int, int) {
	maxVal := arr[0]
	maxIdx := 0
	for i := 0; i < len(arr); i++ {
		if e := arr[i]; e > maxVal {
			maxVal = e
			maxIdx = i
		}
	}
	return maxVal, maxIdx
}

func maxInArrK(arr []int, k int) []int {
	var maxArr []int
	var maxVal, maxIdx int
	maxIdx = -1
	for i := 0; i <= len(arr)-k; i++ {
		if e := arr[i+k-1]; maxIdx >= i && e >= maxVal {
			maxVal = e
			maxIdx = i+k-1
		} else {
			maxVal, maxIdx = maxInArr(arr[i:i+k])
		}
		maxArr = append(maxArr, maxVal)
	}
	return maxArr
}

func maxInArrKDeque(arr []int, k int) []int {
	var maxArr []int
	l := &doublyLinkedList.DoublyLinkedList{}
	i := 0
	for ; i < k; i++ {
		for l.IsEmpty() != true && arr[i] >= arr[l.GetLast()] {
			l.DeleteLast()
		}
		l.InsertLast(i)
	}
	for ; i < len(arr); i++ {
		maxArr = append(maxArr, arr[l.GetFirst()])
		for l.IsEmpty() != true && l.GetFirst() <= i - k {
			l.DeleteFirst()
		}
		for l.IsEmpty() != true && arr[i] >= arr[l.GetLast()] {
			l.DeleteLast()
		}
		l.InsertLast(i)
	}
	maxArr = append(maxArr, arr[l.GetFirst()])
	return maxArr
}

func main() {
	arr := []int{10, 5, 2, 7, 8, 7}
	k := 3
	maxArr := maxInArrK(arr, k)
	fmt.Println(arr, k, maxArr)
	maxArr = maxInArrKDeque(arr, k)
	fmt.Println(arr, k, maxArr)
}