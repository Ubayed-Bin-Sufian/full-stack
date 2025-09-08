package main

// import "fmt"

// // Index is a generic function that searches for x inside slice s.
// func Index[T comparable](s []T, x T) int {
// 	// Loop over slice s. Each element v is of type T.
// 	for i, v := range s {
// 		// Since T is constrained by "comparable",
// 		// we are allowed to use == and != operators on values of type T.
// 		if v == x {
// 			// If we find a match, return the index i.
// 			return i
// 		}
// 	}
// 	// If loop ends without finding x, return -1.
// 	return -1
// }

// func main() {
// 	// Example 1: Using Index with integers
// 	si := []int{10, 20, 15, -10}
// 	fmt.Println(Index(si, 15)) // Output: 2 (15 is at index 2)

// 	// Example 2: Using Index with strings
// 	ss := []string{"foo", "bar", "baz"}
// 	fmt.Println(Index(ss, "hello")) // Output: -1 ("hello" not in slice)
// }

// Output:
// 2
// -1
