package main

import (
    "fmt"
)

func contains(i int, arr []int) bool {
	for _, k := range arr {
		if i == k {
			return true
		}
	}
	return false
}

func findCombinationsUtil(countPointer *int, arr []int, index int, num int, reducedNum int, options []int) {
	if reducedNum >= 0 {
		if reducedNum == 0 {
			// if combination is print it
			fmt.Println(arr[:index])
			*countPointer = (*countPointer + 1)
		} else {
			// find previous number stored in arr
			prev := 1
			if index != 0 {
				prev = arr[index - 1]
			}
			for k := prev; k <= num; k++ {
				if contains(k, options) == true {
					// next element of arr is k
					arr[index] = k
					findCombinationsUtil(countPointer, arr, index + 1, num, reducedNum - k, options)
				}
			}
		}
	}
}

func findCombinations(numberToSum int, options []int) int {
	countPointer := new(int)
	arr := make([]int, numberToSum)
	findCombinationsUtil(countPointer, arr, 0, numberToSum, numberToSum, options)
	return *countPointer
}


func main() {
	n := 4
	options := []int{1,2}
	r := findCombinations(n, options)
	fmt.Println(r)
}