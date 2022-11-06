package main

import (
    "fmt"
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

/* Utility function that puts all 
non-positive (0 and negative) numbers on left 
side of arr[] and return count of such numbers */
func Segregate(arr []int) int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= 0 {
			Swap(&arr[i], &arr[j])
			// increment count of non-positive integers
			j++
		}
	}
	return j
}

func FindSmallestMissing(arr []int) int {
	var i, kabs int
	shift := Segregate(arr)
	size := len(arr)-shift
	arr = arr[shift:]
	// mark elements in the array as visited by setting arr[arr[i]-1] as negative when arr[i] is visited
	for i = 0; i < size; i++ {
		kabs = Abs(arr[i]) - 1
		if kabs < size && arr[kabs] > 0 {
			arr[kabs] = -arr[kabs]
		}
	}
	for i = 0; i < size; i++ {
		if arr[i] > 0 {
			return i + 1
		}
	}

	return size + 1
}

func main() {
	a := []int{3,4,-1,1,1}
	fmt.Println("The smallest positive missing number is",FindSmallestMissing(a))
	a = []int{1,2,0}
	fmt.Println("The smallest positive missing number is",FindSmallestMissing(a))
}