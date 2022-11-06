package main

import (
    "fmt"
    "sort"
)

/*
Compute the running median of a sequence of numbers. 
That is, given a stream of numbers, print out the median of the list so far on each new element.

Recall that the median of an even-numbered list is the average of the two middle numbers.

For example, given the sequence [2, 1, 5, 7, 2, 0, 5], your algorithm should print out:
2
1.5
2
3.5
2
2
2
*/

func sortedCopy(list []int) []int {
	sorted := make([]int, len(list))
	copy(sorted, list)
	sort.Ints(sorted)
	return sorted
}

func calcMedian(list []int) float32 {
	l := len(list)
	c := sortedCopy(list)
	median := float32(c[0])
	if l % 2 == 0 {
		median = (float32(c[l/2-1])+float32(c[l/2]))/2.0
	} else {
		median = float32(c[l/2])
	}
	return median
}

func runningMedian(list []int) []float32 {
	n := len(list)
	medians := make([]float32, n)
	medians[0] = float32(list[0])
	for i := 1; i < n; i++ {
		medians[i] = calcMedian(list[0:i+1])
	}
	return medians
}

func main() {
	list := []int{2, 1, 5, 7, 2, 0, 5}
	medians := runningMedian(list)
	fmt.Println(medians)
}