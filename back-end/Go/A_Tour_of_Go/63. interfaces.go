package main

// import (
// 	"fmt"
// 	"math"
// )

// // Abser is an interface that requires a single method Abs() returning a float64
// type Abser interface {
// 	Abs() float64
// }

// func main() {
// 	var a Abser          	  // 'a' is of interface type Abser
// 	f := MyFloat(-math.Sqrt2) // f is a value of type MyFloat
// 	v := Vertex{3, 4}         // v is a Vertex struct (plain value)

// 	a = f   // ✅ MyFloat implements Abser because it has Abs() with a value receiver
// 	a = &v  // ✅ *Vertex implements Abser because Abs() is defined on pointer receiver

// 	// Important: v (plain Vertex) does NOT implement Abser
// 	// because Abs() is only defined on *Vertex (pointer receiver)
// 	// a = v  // must use &v, not v; compile error

// 	fmt.Println(a.Abs()) // Calls Abs() method of the concrete type stored in 'a'
// }

// // MyFloat is a custom float64 type
// type MyFloat float64

// // Abs method for MyFloat; value receiver works for both MyFloat and *MyFloat
// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// // Vertex is a struct with X, Y coordinates
// type Vertex struct {
// 	X, Y float64
// }

// // Abs method for Vertex; pointer receiver means only *Vertex implements Abser
// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// Output:
// 5