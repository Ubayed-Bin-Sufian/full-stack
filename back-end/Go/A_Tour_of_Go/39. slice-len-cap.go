package main

// import "fmt"

// // This program demonstrates how to manipulate slices in Go by changing their length and capacity.
// // The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}
// 	printSlice(s)

// 	// Slice the slice to give it zero length.
// 	s = s[:0]
// 	printSlice(s)

// 	// Extend its length.
// 	s = s[:4]
// 	printSlice(s)

// 	// Drop its first two values.
// 	s = s[2:]
// 	printSlice(s)
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// Output:
// len=6 cap=6 [2 3 5 7 11 13]
// len=0 cap=6 []
// len=4 cap=6 [2 3 5 7]
// len=2 cap=4 [5 7]