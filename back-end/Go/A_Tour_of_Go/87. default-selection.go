package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	start := time.Now() // record the start time

// 	// tick: sends a signal on the channel every 100ms
// 	tick := time.Tick(100 * time.Millisecond)

// 	// boom: sends a signal once after 500ms
// 	boom := time.After(500 * time.Millisecond)

// 	// helper function to compute elapsed time (rounded to ms)
// 	elapsed := func() time.Duration {
// 		return time.Since(start).Round(time.Millisecond)
// 	}

// 	for {
// 		select {
// 		// case 1: every 100ms, tick will be ready
// 		case <-tick:
// 			fmt.Printf("[%6s] tick.\n", elapsed())

// 		// case 2: after 500ms, boom will trigger once, then return
// 		case <-boom:
// 			fmt.Printf("[%6s] BOOM!\n", elapsed())
// 			return // exit the program

// 		// case 3 (default): if neither tick nor boom is ready,
// 		// print a dot and sleep for 50ms to avoid busy looping
// 		default:
// 			fmt.Printf("[%6s]     .\n", elapsed())
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// }

// Output:
// [    0s]     .
// [  51ms]     .
// [ 102ms] tick.
// [ 102ms]     .
// [ 152ms]     .
// [ 203ms] tick.
// [ 204ms]     .
// [ 254ms]     .
// [ 314ms] tick.
// [ 315ms]     .
// [ 367ms]     .
// [ 417ms] tick.
// [ 418ms]     .
// [ 469ms]     .
// [ 520ms] tick.
// [ 520ms] BOOM!