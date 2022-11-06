package main

import (
    "fmt"
)

/*
Given an array of time intervals (start, end) for classroom lectures (possibly overlapping), find the minimum number of rooms required.

For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.
*/

type Interval struct {
	Start int
	End int
	NumOverlaps int
}

func PrintIntervals(arr []Interval) {
	var intv *Interval
	for i := 0; i < len(arr); i++ {
		intv = &arr[i]
		fmt.Println(intv.Start, intv.End, intv.NumOverlaps)
	}
}

func overlaps(intv1 *Interval, intv2 *Interval) bool {
	return (intv2.Start < intv1.End && intv2.End > intv1.Start) || (intv1.Start < intv2.End && intv1.End > intv2.Start)
}

func findOverlaps(arr []Interval) int {
	var intv1, intv2 *Interval
	for i := 0; i < len(arr); i++ {
		intv1 = &arr[i]
		for j := i; j < len(arr); j++ {
			if i != j {
				intv2 = &arr[j]
				if overlaps(intv1, intv2) {
					intv1.NumOverlaps++
					intv2.NumOverlaps++
				}
			} 
		}
	}
	numOverlaps := 0
	for i := 0; i < len(arr); i++ {
		intv1 = &arr[i]
		if n := intv1.NumOverlaps; n > numOverlaps {
			numOverlaps = n
		}
	}
	return numOverlaps
}

func main() {
	var arr []Interval
	arr = append(arr, Interval{30, 75, 0})
	arr = append(arr, Interval{0, 50, 0})
	arr = append(arr, Interval{60, 150, 0})
	//arr = append(arr, Interval{55, 150, 0})
	fmt.Println(findOverlaps(arr))
	PrintIntervals(arr)
}