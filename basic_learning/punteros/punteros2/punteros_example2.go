package main

/*
Puntero: Variable que almacena la direccion de memoria de otra variable. 
*/

import "fmt"

var marca string = "Mazda"
var pMarca *string = &marca

func main() {
	fmt.Println(*pMarca)

	// Cambiando el valor de la variable: 
	*pMarca = "Kia"
	fmt.Println(*pMarca)

	// Validando si le valor original tambien cambio: 
	fmt.Println("Despues el valor de la marca es:", marca == *pMarca)
	fmt.Println("Espacio en memoria del puntero es:", pMarca)
	fmt.Println("El valor del puntero en memoria es:", *pMarca)
}
