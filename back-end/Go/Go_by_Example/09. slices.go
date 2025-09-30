package main

// import (
//     "fmt"
//     "slices"
// )

// func main() {

// 	// Slices are typed only by the elements they contain. 
// 	// An uninitialized slice equals to nil and has length 0.
//     var s []string
//     fmt.Println("uninit:", s, s == nil, len(s) == 0)

// 	// An empty slice is not nil, but has length 0.
// 	k := []int{}
// 	fmt.Println("int but empty:", k, k == nil, len(k) == 0)

// 	// We can create a slice with the built-in make function;
//     s = make([]string, 3)
//     fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

// 	// We can set and get just like with arrays.
//     s[0] = "a"
//     s[1] = "b"
//     s[2] = "c"
//     fmt.Println("set:", s)
//     fmt.Println("get:", s[2])

//     fmt.Println("len:", len(s))

// 	// append might allocate a completely new array behind the scenes, 
// 	// the returned slice may point to a different memory location. 
// 	// use the returned slice instead of assuming the original one was updated.
//     s = append(s, "d")
//     s = append(s, "e", "f")
//     fmt.Println("apd:", s)

// 	// We can copy a slice with the built-in copy function.
//     c := make([]string, len(s))
//     copy(c, s)
//     fmt.Println("cpy:", c)

// 	// Slices support a "slice" operator with the syntax slice[low:high].
//     l := s[2:5]
//     fmt.Println("sl1:", l)

// 	// This slices up to (but excluding) s[5].
//     l = s[:5]
//     fmt.Println("sl2:", l)

// 	// And this slices from (and including) s[2].
//     l = s[2:]
//     fmt.Println("sl3:", l)

// 	// We can declare and initialize a slice in one line as well.
//     t := []string{"g", "h", "i"}
//     fmt.Println("dcl:", t)

// 	// The slices package contains a number of useful utility functions for slices.
//     t2 := []string{"g", "h", "i"}
//     if slices.Equal(t, t2) {
//         fmt.Println("t == t2")
//     }

// 	// We can build multi-dimensional slices with make as well.
//     twoD := make([][]int, 3)
//     for i := range 3 {
//         innerLen := i + 1
//         twoD[i] = make([]int, innerLen)
//         for j := range innerLen {
//             twoD[i][j] = i + j
//         }
//     }
//     fmt.Println("2d: ", twoD)
// }

// Output:
// uninit: [] true true
// int but empty: [] false true
// emp: [  ] len: 3 cap: 3     
// set: [a b c]
// get: c
// len: 3
// apd: [a b c d e f]
// cpy: [a b c d e f]
// sl1: [c d e]
// sl2: [a b c d e]
// sl3: [c d e f]
// dcl: [g h i]
// t == t2
// 2d:  [[0] [1 2] [2 3 4]]