package main

import (
    "fmt"
)

/*
You run an e-commerce website and want to record the last N order ids in a log. 

Implement a data structure to accomplish this, with the following API:
record(order_id): adds the order_id to the log
get_last(i): gets the ith last element from the log. 

i is guaranteed to be smaller than or equal to N.

You should be as efficient with time and space as possible.
*/

func record(log []int, order_id int) []int {
	return append(log, order_id)
}

func getLast(log []int, i int, l int) int {
	lastI := l - 1 - i
	return log[lastI]
}

func addNOrders(log []int, n int) []int {
	var e int
	l := len(log)
	for i := 0; i < n; i++ {
		e = getLast(log, i, l)
		log = record(log, e)
	}
	return log
}

func main() {
	n := 3
	log := []int{1,2,3,4,5,6}
	log = addNOrders(log, n)
	fmt.Println(log)
}