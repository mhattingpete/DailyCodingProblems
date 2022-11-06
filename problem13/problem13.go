package main

import (
    "fmt"
)

/*
Given an integer k and a string s, 
find the length of the longest substring that contains at most k distinct characters.

For example, 
given s = "abcba" and k = 2, 
the longest substring with k distinct characters is "bcb".
*/

func isValid(dict map[string]int, k int) bool {
	count := 0
	for _, value := range dict {
		if value > 0 {
			count++
		}
	}
	return k >= count
}

func findLongestSubstringWithKDistinct(s string, k int) string {
	start := 0
	end := 0
	max_window_size := 1
	max_window_start := 0
	var cs, ss string
	dict := make(map[string]int)
	for _, c := range s {
		cs = string(c)
		dict[cs] += 1
		end += 1
		for isValid(dict, k) != true {
			ss = string(s[start])
			dict[ss] -= 1
			start += 1
		}
		if window_size := end - start; window_size > max_window_size {
			max_window_size = window_size
			max_window_start = start
		}
	}
	return s[max_window_start:max_window_start + max_window_size]
}

func main() {
	s := "abcba"
	k := 2
	r := findLongestSubstringWithKDistinct(s, k)
	fmt.Println(r)
	s = "aabbcc"
	k = 1
	r = findLongestSubstringWithKDistinct(s, k)
	fmt.Println(r)
	k = 2
	r = findLongestSubstringWithKDistinct(s, k)
	fmt.Println(r)
	k = 3
	r = findLongestSubstringWithKDistinct(s, k)
	fmt.Println(r)
}