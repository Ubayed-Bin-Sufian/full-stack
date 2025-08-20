package main

// import "fmt"

// func main() {
// 	m := make(map[string]int)

// 	m["Answer"] = 42
// 	fmt.Println("The value:", m["Answer"])

// 	m["Answer"] = 48
// 	fmt.Println("The value:", m["Answer"])

// 	delete(m, "Answer")
// 	fmt.Println("The value:", m["Answer"])

// 	// When you read m[key] with two results, Go returns:
// 	// v: the value (or the zero value if missing)
// 	// ok: a bool that’s true if the key exists, false otherwise
// 	// Since we deleted the key, ok is false and v is 0.
// 	v, ok := m["Answer"]
// 	fmt.Println("The value:", v, "Present?", ok)
// }

// Ouput
// The value: 42
// The value: 48
// The value: 0
// The value: 0 Present? false