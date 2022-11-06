package main

import (
    "fmt"
)

/*
Write an algorithm to justify text. Given a sequence of words and an integer line length k, return a list of strings which represents each line, fully justified.

More specifically, you should have as many words as possible in each line. There should be at least one space between each word. Pad extra spaces when necessary so that each line has exactly length k. Spaces should be distributed as equally as possible, with the extra spaces, if any, distributed starting from the left.

If you can only fit one word on a line, then you should pad the right-hand side with spaces.

Each word is guaranteed not to be longer than k.

For example, given the list of words ["the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"] and k = 16, you should return the following:
["the  quick brown", # 1 extra space on the left
"fox  jumps  over", # 2 extra spaces distributed evenly
"the   lazy   dog"] # 4 extra spaces distributed evenly
*/

func justifySentence(sentence []string, sentenceLength int, maxLength int) string {
	s := ""
	spaces := len(sentence)-1
	if spaces < 0 {
		spaces = 0
	}
	spacesNeeded := maxLength - (sentenceLength + spaces)
	spacesPerWord := spacesNeeded / (len(sentence)-1)
	spacesLeft := spacesNeeded - (spacesPerWord * (len(sentence)-1))
	for i := 0; i < len(sentence)-1; i++ {
		for j := 0; j < spacesPerWord; j++ {
			sentence[i] += " "
		}
	}
	for j := 0; j < spacesLeft; j++ {
		sentence[0] += " "
	}
	for i, w := range sentence {
		if i < len(sentence)-1 {
			s += w + " "
		} else {
			s += w
		}
	}
	return s
}

func justify(words []string, maxLength int) []string {
	sentences := []string{}
	sentence := []string{}
	wordlength := 0
	sentenceLength := 0
	spaces := 0
	for _, w := range words {
		wordlength = len(w)
		spaces = len(sentence)-1
		if sentenceLength + spaces + wordlength < maxLength {
			sentence = append(sentence, w)
			sentenceLength += wordlength
		} else {
			sentences = append(sentences, justifySentence(sentence, sentenceLength, maxLength))
			sentence = []string{w}
			sentenceLength = wordlength
		}
	}
	sentences = append(sentences, justifySentence(sentence, sentenceLength, maxLength))
	return sentences
}


func main() {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	k := 16
	sentences := justify(words, k)
	for _, s := range sentences {
		fmt.Printf("Sentence: '%s' with length: %d\n", s, len(s))
	}
}