package main

// import "fmt"

// func sum(s []int, c chan int) {
//     sum := 0
//     for _, v := range s {
//         sum += v
//     }
//     c <- sum                  // send the sum into channel c (blocks until received)
// }

// func main() {
//     s := []int{7, 2, 8, -9, 4, 0}

//     // create a channel that carries ints
// 	// channels are typed
//     c := make(chan int)

//     // run sum() in two goroutines (concurrent workers)
//     go sum(s[:len(s)/2], c)
//     go sum(s[len(s)/2:], c)

//     // receive results from the channel twice
//     // (main will block here until goroutines send their results)
//     x, y := <-c, <-c

//     fmt.Println(x, y, x+y)    // order of x,y may vary (goroutines are concurrent)
// }

// Output:
// -5 17 12