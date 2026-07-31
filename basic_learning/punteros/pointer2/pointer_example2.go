package main

/*
Pointer: A variable that stores the memory address of another variable.
*/

import "fmt"

var CarBrand string = "Mazda"
var PcarBrand *string = &CarBrand

func main() {
	fmt.Println(*PcarBrand)

	// Changing the variable value
	*PcarBrand = "Kia"
	fmt.Println(*PcarBrand)

	// Verify that the original value changed
	fmt.Println("After CarBrand changes:", CarBrand == *PcarBrand)
	fmt.Println("The memory space of memory:", PcarBrand)
	fmt.Println("The memory pointer value ", *PcarBrand)
}
