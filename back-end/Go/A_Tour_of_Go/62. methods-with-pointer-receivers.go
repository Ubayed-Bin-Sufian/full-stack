package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// // Scale has a pointer receiver (*Vertex)
// // Reason 1: It modifies the value of the receiver (changes X and Y).
// // Reason 2: Using a pointer avoids copying the struct on each call,
// //           which is more efficient if the struct is large.
// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// // Abs also has a pointer receiver (*Vertex)
// // Even though Abs does not modify the receiver,
// // we keep it consistent with Scale (all methods on Vertex use pointer receivers).
// // This avoids mixing value and pointer receivers, which can lead to confusion.
// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	// v is initialized as a pointer to Vertex.
// 	v := &Vertex{3, 4}

// 	// Before scaling, show values and the distance from origin (Abs).
// 	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

// 	// Scale modifies the actual values of v because it has a pointer receiver.
// 	v.Scale(5)

// 	// After scaling, the values of v are updated, and Abs reflects the new distance.
// 	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
// }

// Output:
// Before scaling: &{X:3 Y:4}, Abs: 5
// After scaling: &{X:15 Y:20}, Abs: 25
