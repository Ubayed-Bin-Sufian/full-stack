package main

// import "fmt"

// func main() {

// 	// By default an array is zero-valued
//     var a [5]int
//     fmt.Println("emp:", a)

// 	// We can set a value at an index using the array[index] = value syntax,
// 	//  and get a value with array[index]
//     a[4] = 100
//     fmt.Println("set:", a)
//     fmt.Println("get:", a[4])

// 	// len returns the length of an array
//     fmt.Println("len:", len(a))

// 	// declare and initialize an array in one line
//     b := [5]int{1, 2, 3, 4, 5}
//     fmt.Println("dcl:", b)

// 	// count the number of elements
//     b = [...]int{1, 2, 3, 4, 5}
//     fmt.Println("dcl:", b)

// 	// index with :, the elements in between will be zeroed
//     b = [...]int{100, 3: 400, 500}
//     fmt.Println("idx:", b)

// 	// compose types to build multi-dimensional data structures
//     var twoD [2][3]int
//     for i := range 2 {
//         for j := range 3 {
//             twoD[i][j] = i + j
//         }
//     }
//     fmt.Println("2d: ", twoD)

// 	// create and initialize multi-dimensional arrays at once too
//     twoD = [2][3]int{
//         {1, 2, 3},
//         {1, 2, 3},
//     }
//     fmt.Println("2d: ", twoD)
// }

// Output:
// emp: [0 0 0 0 0]
// set: [0 0 0 0 100]    
// get: 100
// len: 5
// dcl: [1 2 3 4 5]      
// dcl: [1 2 3 4 5]      
// idx: [100 0 0 400 500]
// 2d:  [[0 1 2] [1 2 3]]
// 2d:  [[1 2 3] [1 2 3]]