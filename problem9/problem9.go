package main

import (
    "fmt"
)

func sumNonAdjacent(arr []int) int {
	var new_excl int
	incl := arr[0]
	excl := 0
	for i := 1; i < len(arr); i++ {
		if incl > excl {
			new_excl = incl
		} else {
			new_excl = excl
		}
		incl = excl + arr[i]
		excl = new_excl
	}
	if incl > excl {
		return incl
	} else {
		return excl
	}
}

func main() {
	arr := []int{2, 4, 6, 2, 5}
	fmt.Println("Max sum of non adjacent in: %v, is: %d", arr, sumNonAdjacent(arr))
	arr = []int{5, 1, 1, 5}
	fmt.Println("Max sum of non adjacent in: %v, is: %d", arr, sumNonAdjacent(arr))
}