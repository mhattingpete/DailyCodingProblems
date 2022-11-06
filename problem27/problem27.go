package main

import (
    "fmt"
)

/*
Given a string of round, curly, and square open and closing brackets, return whether the brackets are balanced (well-formed).

For example, given the string "([])[]({})", you should return true.

Given the string "([)]" or "((()", you should return false.
*/

func wellFormed(s string) bool {
	open := 0
	opened := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			open++
			opened = append(opened, c)
		}
		if (opened[open-1] == '(' && c == ')') || (opened[open-1] == '[' && c == ']') || (opened[open-1] == '{' &&  c == '}') {
			open--
			opened = opened[:open]

		}
	}
	return open == 0
}

func main() {
	s := "([])[]({})"
	r := wellFormed(s)
	fmt.Printf("Is %s wellFormed: %v\n", s, r)
	s = "([)]"
	r = wellFormed(s)
	fmt.Printf("Is %s wellFormed: %v\n", s, r)
	s = "((()"
	r = wellFormed(s)
	fmt.Printf("Is %s wellFormed: %v\n", s, r)
}