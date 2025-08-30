package main

// import "fmt"

// // Define an interface 'I'.
// // Any type that has a method called M() qualifies as an 'I'.
// type I interface {
// 	M()
// }

// // Define a struct with one field 'S' of type string.
// type T struct {
// 	S string
// }

// // Give type 'T' a method called M().
// // Since 'T' now has M(), it automatically satisfies interface 'I'.
// // (No need to say "implements I" like in Java/C#.)
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// func main() {
// 	// Declare a variable of type 'I' (the interface).
// 	// Assign it a value of type 'T'.
// 	// Go checks: Does 'T' have M()? Yes → so T implements I.
// 	var i I = T{"hello"}

// 	// Call the method via the interface.
// 	// At runtime, it executes T's version of M().
// 	i.M() 
// }

// Output: hello