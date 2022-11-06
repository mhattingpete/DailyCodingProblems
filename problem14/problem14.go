package main

import (
    "fmt"
    "math/rand"
    "math"
)

/*
The area of a circle is defined as Pi * r^2. 
Estimate Pi to 3 decimal places using a Monte Carlo method.

Hint: The basic equation of a circle is x2 + y2 = r2.

Draw a square around the circle and use that:

The area of the circle is Pi * r2,
The area of the square is width2 = (2r)2 = 4r2.

If we divide the area of the circle, by the area of the square we get Pi/4.

The same ratio can be used between the number of points within the square and the number of points within the circle.
So Pi = 4 * points in circle / points in total
*/

func inCircle(x float64, y float64) bool {
	r := 1.0
	xs := square(x)
	ys := square(y)
	if xs + ys <= r {
		return true
	} else {
		return false
	}
}

func square(x float64) float64 {
	return math.Pow(x, 2.0)
}

func uniform(lower float64, upper float64) float64 {
	b := lower
	a := upper + math.Abs(lower)
	return rand.Float64() * a + b
}

func drawPoints(numPoints int) float64 {
	pointsInCircle := 0
	var x, y float64
	for i := 0; i < numPoints; i++ {
		x = uniform(-1.0,1.0)
		y = uniform(-1.0,1.0)
		if inCircle(x,y) == true {
			pointsInCircle++
		}
	}
	return 4.0*float64(pointsInCircle)/float64(numPoints)
}

func main() {
	pi := 3.1415
	piE := drawPoints(100000000)
	fmt.Printf("Pi: %.4f, Pi estimate: %.4f", pi, piE)
}