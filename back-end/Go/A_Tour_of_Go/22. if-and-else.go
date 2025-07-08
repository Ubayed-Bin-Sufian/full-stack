package main

// import (
// 	"fmt"
// 	"math"
// )

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	// can't use v here, though
// 	return lim
// }

// // Both calls to pow return their results before the call to fmt.Println in main begins.
// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }

// Output:
// 27 >= 20
// 9 20
