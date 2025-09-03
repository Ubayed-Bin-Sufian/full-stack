package main

// import "fmt"

// func main() {
// 	// i is of type empty interface, so it can hold any value.
// 	var i interface{} = "hello"

// 	// Type assertion: we tell Go that i is a string.
// 	s := i.(string)
// 	fmt.Println(s) // Output: hello

// 	// Safe type assertion with "comma ok" form.
// 	// s will be "hello", ok will be true because i really is a string.
// 	s, ok := i.(string)
// 	fmt.Println(s, ok) // Output: hello true

// 	// Another safe type assertion, but this time asking for float64.
// 	// i actually holds a string, not a float64.
// 	// So f will be the zero value of float64 (0), and ok will be false.
// 	f, ok := i.(float64)
// 	fmt.Println(f, ok) // Output: 0 false

// 	// Unsafe type assertion: telling Go i is a float64 without checking.
// 	// Since i holds a string, this will cause a runtime panic and crash.
// 	f = i.(float64) // panic: interface conversion: interface {} is string, not float64
// 	fmt.Println(f)
// }

// Output
// hello
// hello true
// 0 false
// panic: interface conversion: interface {} is string, not float64