package main

// import "fmt"

// // 'do' takes an empty interface parameter (interface{})
// // This means 'i' can hold ANY type (int, string, bool, etc.)
// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }

// func main() {
// 	do(21)       
// 	do("hello")  
// 	do(true)    
// }

// Output:
// Twice 21 is 42
// "hello" is 5 bytes long      
// I don't know about type bool!
