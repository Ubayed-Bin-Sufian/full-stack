package main

// import "fmt"

// // fibonacci generates numbers in the Fibonacci sequence
// // and sends them on channel c. It stops when it receives
// // a signal from the quit channel.
// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		// case 1: send the current Fibonacci number (x) to channel c
// 		case c <- x:
// 			// move to the next Fibonacci number
// 			x, y = y, x+y
// 		// case 2: if quit channel receives a value, print "quit" and return
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	// c is used to get Fibonacci numbers
// 	c := make(chan int)
// 	// quit is used to signal stopping
// 	quit := make(chan int)

// 	// Start a goroutine to consume Fibonacci numbers
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			// Receive from channel c and print it
// 			// This will block until a value is available
// 			fmt.Println(<-c)
// 		}
// 		// After receiving 10 numbers, send a signal on quit
// 		quit <- 0
// 	}()

// 	// Run the Fibonacci generator.
// 	// It will keep sending values on c until quit is signaled.
// 	fibonacci(c, quit)
// }

// Output:
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
// quit