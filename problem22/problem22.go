package main

import (
    "fmt"
    "strings"
)

/*
Given a dictionary of words and a string made up of those words (no spaces), return the original sentence in a list. 
If there is more than one possible reconstruction, return any of them. If there is no possible reconstruction, then return null.

For example, given the set of words 'quick', 'brown', 'the', 'fox', and the string "thequickbrownfox", you should return ['the', 'quick', 'brown', 'fox'].

Given the set of words 'bed', 'bath', 'bedbath', 'and', 'beyond', and the string "bedbathandbeyond", return either ['bed', 'bath', 'and', 'beyond] or ['bedbath', 'and', 'beyond'].
*/

func reconstructSentence(s string, dict []string) []string {
	for _, w := range dict {
		if strings.Contains(s, w) {
			s = strings.ReplaceAll(s, w, w+";")
		}
	} 
	return strings.Split(s[0:len(s)-1], ";")
}

func main() {
	dict := []string{"quick", "brown", "the", "fox"}
	text := "thequickbrownfox"
	rec := reconstructSentence(text, dict)
	fmt.Printf("The string: %s, was changed into: %v, using the dictionary: %v", text, rec, dict)

	dict = []string{"bed", "bath", "bedbath", "and", "beyond"}
	text = "bedbathandbeyond"
	rec = reconstructSentence(text, dict)
	fmt.Printf("The string: %s, was changed into: %v, using the dictionary: %v", text, rec, dict)
}