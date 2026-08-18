package main

import "fmt"

// func main() {
// 	x := 45
// 	z := &x // memory address
// 	fmt.Println("🚚 The memory address is: ", z)
// 	fmt.Println("🚚 The memory address is: ", *z) // use * to see the pointer value

// 	// What about changing the z value
// 	*z = 5656
// 	fmt.Println("😎 The new z value is: ", *z)
// }

/*
Output:
 ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓
🚚 The memory address is:  0x49cc2e5da0e0
🚚 The memory address is:  45
😎 The new z value is:  5656
*/

// Practicing pointers:
func main() {
	a := 34
	b := &a

	fmt.Println(a)  // Printing the value
	fmt.Println(b)  // Printing the memory address
	fmt.Println(*b) // Printing the memory address value
}
