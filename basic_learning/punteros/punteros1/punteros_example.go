package main

/*
Puntero: Variable que almacena la direccion de memoria de otra variable.
*/

import "fmt"

var myVar string = "Original message"
var pMyVar *string = &myVar

func main() {
	fmt.Println("The value of myVar is:", myVar)
	fmt.Println("The address of myVar is:", &myVar)
	fmt.Println("The value of pMyVar is:", pMyVar)
	fmt.Println("The value of *pMyVar is:", *pMyVar) // Este es el valor de la variable myVar el cual debe ser igual al valor original
	fmt.Println("The value of *pMyVar is:", &pMyVar) // Se le asigna otro valor de memoria

	fmt.Println()	
	fmt.Println("---- Antes de la modificacion ----")
	fmt.Println("The original value of myVar is:", myVar)
	fmt.Println("The new value of *pMyVar is:", *pMyVar)

	fmt.Println()
	fmt.Println("---- Despues de la modificacion ----")
	*pMyVar = "Adding new meesage"
	fmt.Println("The value of myVar is:", myVar)	
	fmt.Println("The new value of pMyVar is:", *pMyVar)	
}
