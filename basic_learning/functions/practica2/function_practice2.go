package main

import (
	"fmt"
	"myProject/basic_learning/functions/practica2/op_mat"
)

func main() {
	// 2. Usamos el nombre del nuevo paquete
	totalSuma := op_mat.Suma(30, 90)
	totalRest := op_mat.Resta(90, 30)
	totalMultiplicacion := op_mat.Multiplicacion(10, 2)
	totalDivision := op_mat.Division(10, 2)
	fmt.Println("Total suma es: ", totalSuma)
	fmt.Println("Total resta es: ", totalRest)
	fmt.Println("Total multiplicacion es: ", totalMultiplicacion)
	fmt.Println("Total division es: ", totalDivision)
}