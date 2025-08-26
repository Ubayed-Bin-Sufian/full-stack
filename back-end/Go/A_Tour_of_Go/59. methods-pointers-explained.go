package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// // Abs is a function (not a method) that calculates
// // the Euclidean distance of a Vertex from the origin.
// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// // Scale takes a pointer to a Vertex and scales its coordinates
// // by multiplying them with the factor f.
// // Because we pass a pointer, the original Vertex gets modified.
// func Scale(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	Scale(&v, 10)
// 	fmt.Println(Abs(v))
// }

// Output:
// 50