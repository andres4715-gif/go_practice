package main

import (
	"fmt"
	"myProject/basic_learning/slice/slicePractice/practica"
)

func main() {
	myArray := []int{1, 2, 3, 4, 5}
	result, err := practica.RetornandoArray(myArray)

	if err != nil {
		fmt.Println("Hubo un error:", err)
	} else {
		fmt.Println("El resultado es: ", result)
	}
}
