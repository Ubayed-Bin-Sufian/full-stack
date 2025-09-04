package main

// import "fmt"

// // Person is a struct type with two fields: Name and Age
// type Person struct {
// 	Name string
// 	Age  int
// }

// // String method makes Person implement the fmt.Stringer interface.
// // Whenever you print a Person with fmt (like fmt.Println),
// // this method decides what string gets shown.
// func (p Person) String() string {
// 	// fmt.Sprintf formats the string with placeholders %v
// 	// %v means "default format" for the value.
// 	// Here we’re returning a string like: "Arthur Dent (42 years)"
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// func main() {
// 	// Create two Person values using struct literals
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}

// 	// When printing, fmt.Println checks if Person has a String() method.
// 	// Since it does, the custom string is printed instead of {Name Age}.
// 	fmt.Println(a, z)
// }

// Output:
// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)