package main

import (
    "fmt"
)

/*
You are given an array of non-negative integers that represents a two-dimensional elevation map where each element is unit-width wall and the integer is the height. 
Suppose it will rain and all spots between two walls get filled up.

Compute how many units of water remain trapped on the map in O(N) time and O(1) space.

For example, given the input [2, 1, 2], we can hold 1 unit of water in the middle.

Given the input [3, 0, 1, 3, 0, 5], we can hold 3 units in the first index, 2 in the second, and 3 in the fourth index (we cannot hold 5 since it would run off to the left), 
so we can trap 8 units of water.
*/

func inArr(e int, arr []int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if e == arr[i] {
			return i, true
		}
	}
	return -1, false
}

func findSaddlepoints(elevationMap []int) []int {
	var saddlepoints []int
	saddlepoints = append(saddlepoints, 0)
	for i := 1; i < len(elevationMap)-1; i++ {
		if (elevationMap[i] < elevationMap[i+1] && elevationMap[i] < elevationMap[i-1]) || (elevationMap[i] > elevationMap[i+1] && elevationMap[i] > elevationMap[i-1]) {
			saddlepoints = append(saddlepoints, i)
		}
	}
	saddlepoints = append(saddlepoints, len(elevationMap)-1)
	return saddlepoints
}

func waterTrapped(elevationMap []int) int {
	unitsWater := 0
	saddlepoints := findSaddlepoints(elevationMap)
	for i := 1; i < len(elevationMap)-1; i++ {
		if elevationMap[i] < elevationMap[i+1] && elevationMap[i] < elevationMap[i-1] {
			if idx, in := inArr(i, elevationMap); in == true {
				//
			}
		}
	} 
	return unitsWater
}

func main() {
	elevationMap := []int{2,1,2}
	//fmt.Println(waterTrapped(elevationMap))
	elevationMap = []int{3,0,1,3,0,5}
	water := waterTrapped(elevationMap)
	fmt.Println()
	fmt.Println(water)
}