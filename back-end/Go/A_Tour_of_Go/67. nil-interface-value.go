package main

// import "fmt"

// // I is an interface with a single method M
// type I interface {
// 	M()
// }

// func main() {
// 	// Declare a variable i of type I (interface)
// 	// At this point, i is a "nil interface"
// 	// meaning: (value = nil, type = nil)
// 	var i I

// 	// Prints the contents of the interface pair (value, type)
// 	// Expect output: (<nil>, <nil>)
// 	describe(i)

// 	// Try to call M() on i
// 	// Problem: since the type inside i is nil,
// 	// Go does not know which concrete method M() to call
// 	// This causes a runtime panic
// 	i.M()
// }

// // describe shows what is inside the interface value:
// // the concrete value and the concrete type
// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// Output:
// (<nil>, <nil>)
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x8ac059]

// goroutine 1 [running]:
// main.main()
//         H:/Courses/Frontend/full-stack/back-end/Go/A_Tour_of_Go/67. nil-interface-value.go:24 +0x19
// exit status 2
