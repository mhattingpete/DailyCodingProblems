package main

import (
    "fmt"
)

/*
The edit distance between two strings refers to the minimum number of character insertions, deletions, and substitutions required to change one string to the other. 
For example, the edit distance between “kitten” and “sitting” is three: substitute the “k” for “s”, substitute the “e” for “i”, and append a “g”.

Given two strings, compute the edit distance between them.
*/

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func levenshteinEditDistance(s1 string, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	minLen := min(l1, l2)
	count := 0
	for i := 0; i < minLen; i++ {
		if s1[i] == s2[i] {
			count++
		}
	}
	return max(l1, l2) - count
}

func main() {
	s1 := "kitten"
	s2 := "sitting"
	dist := levenshteinEditDistance(s1, s2)
	fmt.Printf("The Levenshtein edit distance between %s and %s is: %d\n", s1, s2, dist)
	s2 = "kitchen"
	dist = levenshteinEditDistance(s1, s2)
	fmt.Printf("The Levenshtein edit distance between %s and %s is: %d", s1, s2, dist)
}