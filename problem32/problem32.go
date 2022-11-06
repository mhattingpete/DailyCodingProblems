package main

import (
    "fmt"
    "math/rand"
)

/*
Suppose you are given a table of currency exchange rates, represented as a 2D array. 
Determine whether there is a possible arbitrage: 
that is, whether there is some sequence of trades you can make, starting with some amount A of any currency, 
so that you can end up with some amount greater than A of that currency.

There are no transaction costs and you can trade fractional quantities.
*/

// generates a random float32 number in the interval [minval;maxval)
func generateRandNum(minval float32, maxval float32) float32 {
	return (rand.Float32() + minval) * (maxval - minval)
}

// creates an empty matrix
func initializeMatrix(height int, width int) [][]float32 {
	matrix := make([][]float32, height)
	for i := range matrix {
		matrix[i] = make([]float32, width)
	}
	return matrix
}

// generates a matrix initialized with random numbers
func generateRandMatrix(height int, width int) [][]float32 {
	randMatrix := initializeMatrix(height, width)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i == j {
				randMatrix[i][j] = 1.0
			}
			else {
				randMatrix[i][j] = generateRandNum(0.0, 5.0)
			}
		}
	}
	return randMatrix
}

// determines if there is an arbitrage in a matrix
func determineArbitrage(matrix [][]float32) bool {
	
}

func main() {
	var startAmount float32 := 10.0
	matrix := generateRandMatrix(4, 4)

}