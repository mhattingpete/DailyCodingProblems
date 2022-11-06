package main

import (
    "fmt"
    "time"
)

type fun func(int) int

func funcToCall(n int) int {
	return 2 * n
}

func jobSchedule(f fun, n int) int {
	time.Sleep(time.Duration(n) * time.Second)
	return f(n)
}

func main() {
	waitTime := 2
	r := jobSchedule(funcToCall, waitTime)
	fmt.Printf("Waited %vs, and got: %d", waitTime, r)
}