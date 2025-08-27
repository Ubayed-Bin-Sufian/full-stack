package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// // Method with a value receiver
// // Can be called on both a value (v) and a pointer (&v).
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// // Normal function that takes a value parameter
// // Must be called with exactly a value of type Vertex (no auto conversion).
// func AbsFunc(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}

// 	// Method call with value: OK
// 	fmt.Println(v.Abs())

// 	// Function call with value: OK
// 	fmt.Println(AbsFunc(v))

// 	p := &Vertex{4, 3}

// 	// Method call with pointer: OK (Go automatically does (*p).Abs())
// 	fmt.Println(p.Abs())

// 	// Function call requires a value: must dereference pointer manually
// 	fmt.Println(AbsFunc(*p))
// }

// Output:
// 5
// 5
// 5
// 5
