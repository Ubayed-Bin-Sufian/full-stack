package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 2)  // buffer size = 2
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)  // receives 1 → now buffer has [2]
// 	fmt.Println(<-ch)  // receives 2 → now buffer is empty

// 	// Next line will try to overfill
//     ch <- 3  // works fine, since the buffer has space again
//     fmt.Println("This will not print until there's room in the buffer")
// }

// Output:
// 1
// 2
