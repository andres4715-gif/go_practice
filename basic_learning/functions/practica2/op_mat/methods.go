package op_mat

func Suma(a, b int) int {
	return a + b
}

func Resta(a, b int) int {
	return a - b
}

func Multiplicacion(a, b int) int {
	return a * b
}

func Division(a, b int) int {
	return a / b
}

func SumandoVariosNumeros(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
