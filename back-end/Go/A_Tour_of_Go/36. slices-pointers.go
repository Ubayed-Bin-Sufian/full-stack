package main

// import "fmt"

// // A slice does not store any data, it just describes a section of an underlying array.
// // Changing the elements of a slice modifies the corresponding elements of its underlying array.
// func main() {
// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names)

// 	a := names[0:2]
// 	b := names[1:3] // [Paul George]
// 	fmt.Println(a, b)

// 	b[0] = "XXX"
// 	fmt.Println(a, b)
// 	fmt.Println(names)
// }

// Output:
// [John Paul George Ringo]
// [John Paul] [Paul George]
// [John XXX] [XXX George]
// [John XXX George Ringo]
