package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	// Create a Reader from a string. 
// 	// This allows us to read the string like a stream of data (io.Reader).
// 	r := strings.NewReader("Hello, Reader!")

// 	// Create a buffer (slice of bytes) with 8 slots.
// 	// This is where Read will put the data it reads.
// 	b := make([]byte, 8)
// 	// fmt.Println(b) // This would print the empty buffer: [0 0 0 0 0 0 0 0]

// 	for {
// 		// Calls the Read method on r
// 		// Read up to len(b) bytes from the Reader into the buffer.
// 		// r will try to fill that bucket b with up to 8 characters from the string.
// 		// n = number of bytes actually read
// 		// err = error value (io.EOF when the Reader has no more data)
// 		n, err := r.Read(b)

// 		// Print how many bytes were read, the error, and the raw buffer (numbers).
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)

// 		// Print only the portion of the buffer that was actually filled as a string.
// 		fmt.Printf("b[:n] = %q\n", b[:n])

// 		// Stop the loop if we've reached the end of the Reader.
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// Output:
// n = 8 err = <nil> b = [72 101 108 108 111 44 32 82] <- ASCII codes of the characters (H e l l o ,   R)
// b[:n] = "Hello, R"
// n = 6 err = <nil> b = [101 97 100 101 114 33 32 82] <- ASCII codes of the characters (e a d e r !   R)
// b[:n] = "eader!"
// n = 0 err = EOF b = [101 97 100 101 114 33 32 82] <- ASCII codes of the characters (e a d e r !   R)
// b[:n] = ""
// In the last slice, same slice because that’s basically "eader! R" left over in memory.
