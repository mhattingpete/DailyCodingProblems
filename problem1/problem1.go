package main

import (
    "fmt"
)

func TwoNumbersAddToK(arr []int, k int) (bool) {
	var nums []int 
	for _, e := range arr {
		for _, i := range nums {
			if i+e == k {
				fmt.Printf("Elements that add up to %v is: %v and %v",k,e,i)
				return true
			}
		}
		nums = append(nums, e)
	}
	return false
}

func main() {
	arr := []int{10, 15, 3, 7}
	k := 17
	b := TwoNumbersAddToK(arr, k)
    fmt.Printf("\nDoes the array contains two elements that add up to: %v? %v\n", k, b)
    arr = []int{10, 15, 3, 17, 19, -1, -3, -5, 0, 7}
	b = TwoNumbersAddToK(arr, k)
    fmt.Printf("\nDoes the array contains two elements that add up to: %v? %v\n", k, b)
    arr = []int{10, 15, 3, 17, 19, -1, -3, -5, -2, 7}
	b = TwoNumbersAddToK(arr, k)
    fmt.Printf("\nDoes the array contains two elements that add up to: %v? %v\n", k, b)
}