package main

// import (
// 	"fmt"
// 	"math"
// )

// // I is an interface that requires one method: M()
// // Any type that has an M() method is said to "implement" I.
// type I interface {
// 	M()
// }

// // T is a struct type that stores a string.
// type T struct {
// 	S string
// }

// // Method M() for type *T (pointer to T).
// // This means only *T (not T itself) implements the interface I.
// // When called, it prints the value of S.
// func (t *T) M() {
// 	fmt.Println(t.S)
// }

// // F is a new type based on float64.
// type F float64

// // Method M() for type F (value receiver).
// // This means both F and *F can implement the interface I.
// // When called, it prints the numeric value.
// func (f F) M() {
// 	fmt.Println(f)
// }

// func main() {
// 	// Declare a variable i of type I (the interface).
// 	// At this point, i is "empty" — no concrete value or type yet.
// 	var i I

// 	// Assign i to hold a *T value.
// 	// Under the hood, i is now a tuple:
// 	//   (value = &T{"Hello"}, type = *T)
// 	i = &T{"Hello"}
// 	describe(i) // shows the concrete value and type inside i
// 	i.M()       // calls the *T version of M(), prints "Hello"

// 	// Assign i to hold a value of type F.
// 	// Under the hood, i is now:
// 	//   (value = 3.14159..., type = F)
// 	i = F(math.Pi)
// 	describe(i) // shows the concrete value and type inside i
// 	i.M()       // calls the F version of M(), prints "3.14159..."
// }

// // describe prints out the "tuple" stored in the interface variable.
// // %v prints the value, %T prints the concrete type.
// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// Output:
// (&{Hello}, *main.T)
// Hello
// (3.141592653589793, main.F)
// 3.141592653589793