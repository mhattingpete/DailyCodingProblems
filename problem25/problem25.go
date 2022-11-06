package main

import (
    "fmt"
)

/*
Implement regular expression matching with the following special characters:

. (period) which matches any single character
* (asterisk) which matches zero or more of the preceding element
That is, implement a function that takes in a string and a valid regular expression and returns whether or not the string matches the regular expression.

For example, given the regular expression "ra." and the string "ray", your function should return true. The same regular expression on the string "raymond" should return false.

Given the regular expression ".*at" and the string "chat", your function should return true. The same regular expression on the string "chats" should return false.
*/

func charMatch(char string, re string) bool {
	if re == "." {
		return true
	} else if re == char {
		return true
	} else {
		return false
	}
}

func match(str string, re string) bool {
	j := 0
	r := ""
	s := ""
	for i := 0; i < len(str); i++ {
		if j < len(re) {
			r = string(re[j])
			s = string(str[i])
			fmt.Println(r,s,j,i)
			if r == "*" {
				if j+1 < len(re) {
					r = string(re[j+1])
					if charMatch(s, r) == true {
						j += 2
					} else {
						r = string(re[j-1])
						if charMatch(s, r) == false {
							return false
						}
					}
				} else {
					r = string(re[j-1])
					if charMatch(s, r) == false {
						return false
					} else {
						j++
					}
				}
			} else {
				if charMatch(s, r) == false {
					return false
				} else {
					j++
				}
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "ray"
	re := "ra."
	m := match(s, re)
	fmt.Printf("Does Str: %s, match with Re: %s: %v\n", s, re, m)
	s = "chaat"
	re = ".*at"
	m = match(s, re)
	fmt.Printf("Does Str: %s, match with Re: %s: %v", s, re, m)
}