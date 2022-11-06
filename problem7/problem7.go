package main

import (
    "fmt"
    "strconv"
)

func isValidNum(num int) bool {
	maxnum := 26
	if num <= maxnum && num > 0 {
		return true
	} else {
		return false
	}
}

func isValid(message string, startNum int) bool {
	var m int
	l := len(message)
	if l <= 0 {
		return false
	}
	if l+startNum % 2 != 0 {
		l -= 1
	}
	for i := startNum; i < l; i += 2 {
		m, _ = strconv.Atoi(message[i:i+2])
		if isValidNum(m) == false {
			return false
		}
	}
	return true
}

func numDecodes(message string) int {
	var valid bool
	var n int
	startNum := 0
	combinations := 1
	valid = isValid(message, startNum)
	if valid == true {
		combinations += 2
	}
	n = int(message[0])
	if isValidNum(n) == true {
		startNum = 1
		valid = isValid(message, startNum)
		if valid == true {
			combinations++
		}
	}
	return combinations
}

func main() {
	m := "111213"
	n := numDecodes(m)
	fmt.Printf("The message %v, can be decoded in %d number of ways", m, n)
}