package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // SafeCounter is safe to use concurrently (by multiple goroutines at once).
// // It wraps a map with a mutex so we don’t get race conditions.
// type SafeCounter struct {
// 	mu sync.Mutex       // mutex to protect concurrent access
// 	v  map[string]int   // shared map holding counters
// }

// // Inc increments the counter for the given key.
// func (c *SafeCounter) Inc(key string) {
// 	c.mu.Lock()         // 🚪 lock before accessing the map
// 	// Only one goroutine can run the next line at a time
// 	c.v[key]++
// 	c.mu.Unlock()       // 🚪 unlock so others can access
// }

// // Value returns the current value of the counter for the given key.
// func (c *SafeCounter) Value(key string) int {
// 	c.mu.Lock()         // 🚪 lock before reading from the map
// 	// defer makes sure the unlock always runs,
// 	// even if the function exits early
// 	defer c.mu.Unlock()
// 	return c.v[key]     // safe read
// }

// func main() {
// 	// Create a SafeCounter with an initialized map
// 	c := SafeCounter{v: make(map[string]int)}

// 	// Start 1000 goroutines, each incrementing "somekey" once
// 	for i := 0; i < 1000; i++ {
// 		go c.Inc("somekey")
// 	}

// 	// Give goroutines time to finish their work
// 	time.Sleep(time.Second)

// 	// Safely get and print the final value of "somekey"
// 	fmt.Println(c.Value("somekey"))
// }

// Output:
// 1000
