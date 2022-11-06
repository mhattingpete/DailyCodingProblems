package main

import (
    "fmt"
    "strconv"
)

/*
Run-length encoding is a fast and simple method of encoding strings. 
The basic idea is to represent repeated successive characters as a single count and character. 
For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".

Implement run-length encoding and decoding. 
You can assume the string to be encoded have no digits and consists solely of alphabetic characters. 
You can assume the string to be decoded is valid.
*/

type Node struct {
	char byte
	count int
}

func createNode(c byte) *Node {
	return &Node{char:c, count:0}
}

func runLengthEncode(s string) string {
	rs := ""
	j := 0
	r := make(map[int]*Node)
	r[j] = createNode(s[0])
	r[j].count++
	for i := 1; i < len(s); i++ {
		if s[i] == r[j].char {
			r[j].count++
		} else {
			j++
			r[j] = createNode(s[i])
			r[j].count++
		}
	}
	for k := 0; k < len(r); k++ {
		rs += strconv.Itoa(r[k].count) + string(r[k].char)
	}
	return rs
}

func runLengthDecode(s string) string {
	rs := ""
	num := 0
	var char string
	for i := 0; i < len(s); i += 2 {
		num, _ = strconv.Atoi(string(s[i]))
		char = string(s[i+1])
		for j := 0; j < num; j++ {
			rs += char
		}
	}
	return rs
}

func main() {
	s := "AAAABBBCCDAA"
	es := "4A3B2C1D2A"
	fmt.Println(es == runLengthEncode(s))
	fmt.Println(s == runLengthDecode(runLengthEncode(s)))
}