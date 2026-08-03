package main

/*
Pointer: A Variable that stores the memory address of another variable
*/

import "fmt"

var myVar string = "Original message"
var pMyVar *string = &myVar

func main() {
	fmt.Println("\nThe value of myVar is:", myVar)
	fmt.Println("The address of myVar is:", &myVar)
	fmt.Println("The value of pMyVar is:", pMyVar)
	fmt.Println("The value of *pMyVar is:", *pMyVar) // This is the value of myVar, witch should be equal to the original value
	fmt.Println("The value of &pMyVar is:", &pMyVar) // Another memory value is assigned

	fmt.Println("\n---- Before Modification ----")
	fmt.Println("The original value of myVar is:", myVar)
	fmt.Println("The new value of *pMyVar is:", *pMyVar)

	fmt.Println()
	fmt.Println("---- After Modification ----")
	*pMyVar = "Adding new message"
	fmt.Println("The value of myVar is:", myVar)
	fmt.Println("The new value of pMyVar is:", *pMyVar)
}
