package main

import (
	"fmt"
	"myProject/basic_learning/functions/practica2/op_mat"
)

func main() {
	// 2. Usamos el nombre del nuevo paquete
	totalAddition := op_mat.Add(30, 90)
	totalSubtraction := op_mat.Subtraction(90, 30)
	totalTimes := op_mat.Times(10, 2)
	totalDivided := op_mat.DividedBy(10, 2)
	total := op_mat.TotalAddition(10, 20, 30, 5, 2, 2, 1)

	// Print
	fmt.Println("Total suma es: ", totalAddition)
	fmt.Println("Total resta es: ", totalSubtraction)
	fmt.Println("Total multiplicacion es: ", totalTimes)
	fmt.Println("Total division es: ", totalDivided)
	fmt.Println("Total suma varios numeros es: ", total)
}
