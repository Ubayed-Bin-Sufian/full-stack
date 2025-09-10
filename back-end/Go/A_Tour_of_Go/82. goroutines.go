package main

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	// starts a new goroutine that runs say("world") concurrently
// 	go say("world")
// 	say("hello")
// }

// Output:
// hello
// world
// world
// hello
// world
// hello
// world
// hello
// world
// hello 