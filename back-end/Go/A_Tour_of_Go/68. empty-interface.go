package main

// import "fmt"

// func main() {
// 	// Declare a variable of type empty interface.
// 	// Since empty interface can hold any type, right now it's nil.
// 	var i interface{}
// 	describe(i)

// 	// Assign an int value to i. Now i holds an integer (42).
// 	i = 42
// 	describe(i)

// 	// Assign a string value to i. Now i holds a string ("hello").
// 	i = "hello"
// 	describe(i)
// }

// // describe prints the value (%v) and its type (%T).
// // Because parameter i is of type interface{}, it can accept any value.
// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// Output:
// (<nil>, <nil>)
// (42, int)      
// (hello, string)