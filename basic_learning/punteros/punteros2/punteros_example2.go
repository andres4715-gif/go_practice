package main

/*
Puntero: Variable que almacena la direccion de memoria de otra variable. 
*/

import "fmt"

var marca string = "Mazda"
var marcaPuntero *string = &marca

func main() {
	fmt.Println(*marcaPuntero)

	// Cambiando el valor de la variable: 
	*marcaPuntero = "Kia"
	fmt.Println(*marcaPuntero)

	// Validando si le valor original tambien cambio: 
	fmt.Println("Despues el valor de la marca es:", marca == *marcaPuntero)
	fmt.Println("Espacio en memoria del puntero es:", marcaPuntero)
	fmt.Println("El valor del puntero en memoria es:", *marcaPuntero)
}
