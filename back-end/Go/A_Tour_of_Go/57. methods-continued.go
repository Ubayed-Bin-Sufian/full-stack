package main

// import (
// 	"fmt"
// 	"math"
// )

// type MyFloat float64

// // If you want to add methods to an existing type (like int, float64, string),
// // you must first create a new named type based on it, and then define methods on that.
// // You can only declare a method with a receiver 
// // whose type is defined in the same package as the method.
// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func main() {
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println(f.Abs())
// }

// Output:
// 1.4142135623730951