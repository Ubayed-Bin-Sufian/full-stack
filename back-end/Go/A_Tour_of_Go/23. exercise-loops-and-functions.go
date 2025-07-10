package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	count := 0

	for count < 10 {
		if z < x {
			z -= (z*z - x) / (2 * z) // Newton's method for square root approximation
			// fmt.Println(z)  // TO check how the value comes nearest to square root
		}
		count++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1))
}
