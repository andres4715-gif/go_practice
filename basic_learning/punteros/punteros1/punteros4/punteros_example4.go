package main

/*
Puntero: Variable que almacena la direccion de memoria de otra variable.
*/

import "fmt"

func gettingData(pNumber *int) int {
	return *pNumber
}

func savings(pMySavings *int) {
	*pMySavings = *pMySavings - 1000
	fmt.Println("💥 The new value is: ", *pMySavings)
} 

func main() {
	initialData := 2000
	MyData := gettingData(&initialData)

	savings(&MyData)

	fmt.Println("🚀 The int value is: ", MyData) 
}
