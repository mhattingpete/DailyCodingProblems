package main

import (
    "fmt"
)

type fun func(int, int) int
type cfun func(fun) int

func cons(a int, b int) cfun {
	pair := func(f fun) int {
		return f(a,b)
	}
	return pair
}

func car(cf cfun) int {
	getFirst := func(a int, b int) int {
		return a
	}
	return cf(getFirst)
}

func cdr(cf cfun) int {
	getLast := func(a int, b int) int {
		return b
	}
	return cf(getLast)
}

func main() {
	c := cons(3, 4)
	first := car(c)
	last := cdr(c)
	fmt.Printf("first: %d, last: %d\n", first, last)
}