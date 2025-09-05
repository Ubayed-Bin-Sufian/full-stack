package main

// import (
// 	"fmt"
// 	"time"
// )

// // MyError is our own custom error type
// type MyError struct {
// 	When time.Time
// 	What string
// }

// // Error implements the built-in 'error' interface.
// // Any type with this method can act like an error.
// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s", 
// 		e.When,
// 		e.What)
// }

// // run is a function that *returns* an error value.
// // In this case, we are forcing it to always return our custom error.
// func run() error {
// 	return &MyError{       // return a pointer to MyError
// 		time.Now(),        // capture the current time
// 		"it didn't work",  // set the error message
// 	}
// }

// func main() {
// 	if err := run(); err != nil {
// 		// Since run always returns an error, this block runs.
// 		// Print the error. fmt.Println sees it's an 'error' and
// 		// calls the Error() method we defined above.
// 		fmt.Println(err)
// 	}
// }

// Output:
// at 2025-09-05 15:53:14.3567878 +0600 +06 m=+0.000706801, it didn't work
