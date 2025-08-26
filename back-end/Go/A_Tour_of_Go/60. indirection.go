package main

// import "fmt"

// type Vertex struct {
// 	X, Y float64
// }

// // Method with pointer receiver
// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// // Normal function with pointer parameter
// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}

// 	// Even though v is a value, Go lets us call the pointer receiver method.
// 	// Internally, this is rewritten as (&v).Scale(2).
// 	v.Scale(2)  // method call on a value
// 	// fmt.Println("(&v).Scale(2) =", v)

// 	// For a function, Go does NOT auto-take the address.
// 	// We must explicitly pass &v.
// 	ScaleFunc(&v, 10)
// 	// fmt.Println("ScaleFunc(&v, 10) =", v)

// 	// Now p is already a *Vertex (pointer).
// 	p := &Vertex{4, 3}

// 	// Calling method with pointer receiver on a pointer works directly.
// 	p.Scale(3)  // method call on a pointer
// 	// fmt.Println("p.Scale(3) =", p)

// 	// For functions, since p is already a pointer, this works as is.
// 	ScaleFunc(p, 8)

// 	fmt.Println(v, p)
// }

// Output:
// {60 80} &{96 72}