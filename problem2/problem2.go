package main

import (
    "fmt"
)

func ProductOfArray(arr []int) (int) {
	p := 1
	for _, e := range arr {
		p *= e
	}
	return p
}

func ProductWithDivision(arr []int) ([]int) {
	out := make([]int, len(arr))
	p := ProductOfArray(arr)
	for i, e := range arr {
		out[i] = p/e
	}
	return out
}

func ProductWithoutDivision(arr []int) ([]int) { // not done
	out := make([]int, len(arr))
	p := ProductOfArray(arr)
	for i, e := range arr {
		out[i] = p/e
	}
	return out
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	out := ProductWithDivision(arr)
    fmt.Println(out)
    out = ProductWithoutDivision(arr)
    fmt.Println(out)   
}