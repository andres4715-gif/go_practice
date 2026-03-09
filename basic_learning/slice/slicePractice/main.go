package main

import (
	"fmt"
	"myProject/basic_learning/slice/slicePractice/practica"
)

func main() {
	myArrayConDatos := []int{1, 2, 3, 4, 5} // Para ejemplos con array con datos
	// myArrayVacio := []int{} // Para ejemplos con array sin datos
	result, err := practica.RetornandoArrayConDatos(myArrayConDatos)

	if err != nil {
		fmt.Println("Hubo un error:", err)
	} else {
		fmt.Println("El resultado con los datos agregados es: ", result)
	}

	eliminandoDato, eliminandoDatoError := practica.RetornandoArrayConDatosEliminados(myArrayConDatos)
	if eliminandoDatoError != nil {
		fmt.Println("Hubo un error:", eliminandoDato)
	} else {
		fmt.Println("El resultado sin el dato eliminado es: ", eliminandoDato)
	}
}
