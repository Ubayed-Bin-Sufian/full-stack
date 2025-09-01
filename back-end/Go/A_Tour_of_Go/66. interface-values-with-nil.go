package main

// import "fmt"

// // Define an interface with a single method M
// type I interface {
// 	M()
// }

// // A simple struct with one string field
// type T struct {
// 	S string
// }

// // Method M defined on *T (pointer receiver)
// func (t *T) M() {
// 	if t == nil {
// 		// If the receiver is nil, handle gracefully instead of crashing
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	// Otherwise, print the string field
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I // i is an interface value, initially nil (no type, no value)

// 	var t *T  // t is a nil pointer to T
// 	i = t     // i now holds (type = *T, value = nil). Non-nil interface!
// 	describe(i) // prints: (<nil>, *main.T)
// 	i.M()       // safe: method sees nil receiver, prints "<nil>"

// 	// Assign a non-nil *T value to the interface
// 	i = &T{"hello"}
// 	describe(i) // prints: (&{hello}, *main.T)
// 	i.M()       // method prints "hello"
// }

// // Helper to show what's inside the interface
// func describe(i I) {
// 	// First %v shows the value, second %T shows the concrete type
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// Output:
// (<nil>, *main.T)
// <nil>
// (&{hello}, *main.T)
// hello
