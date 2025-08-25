package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// // This has a value receiver (v Vertex).
// // That means the method gets a copy of the Vertex.
// // It does not change the original vertex.
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// // This has a pointer receiver (v *Vertex).
// // That means the method can modify the original struct.
// // It scales (multiplies) both coordinates by f.
// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(10)
// 	fmt.Println(v.Abs())
// }

// Output:
// 50
