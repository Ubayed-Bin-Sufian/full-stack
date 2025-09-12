package main

// import (
// 	"fmt"
// )

// // fibonacci sends 'n' Fibonacci numbers into channel c
// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x          // send the current value of x into the channel
// 		x, y = y, x+y   // update to the next Fibonacci numbers
// 	}
// 	close(c)            // close channel to signal no more values will be sent
// }

// func main() {
// 	c := make(chan int, 10)          // make a buffered channel of capacity 10

// 	go fibonacci(cap(c), c)          // start a goroutine: send 'cap(c)' Fibonacci numbers
// 	// range automatically receives from c until it’s closed
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// output:
// 0
// 1 
// 1 
// 2 
// 3 
// 5 
// 8 
// 13
// 21
// 34
