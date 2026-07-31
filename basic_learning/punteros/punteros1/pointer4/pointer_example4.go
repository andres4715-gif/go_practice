package main

/*
Pointer: A variable that stores the memory address of another variable
*/

import "fmt"

func GettingData(pNumber *int) int {
	return *pNumber
}

func Savings(pMySavings *int) {
	*pMySavings = *pMySavings - 1000
	fmt.Println("💥 The new value is: ", *pMySavings)
}

func main() {
	initialData := 2000
	MyData := GettingData(&initialData) // MyData = 2000

	Savings(&MyData)

	fmt.Println("🚀 The int value is: ", MyData)
}
