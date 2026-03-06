package main

import "fmt"

func cambiarValor(puntero* int) {
	*puntero = 100
}

func main () {
	myVar := 10
	fmt.Println("El valor de myVar es: ", myVar)
	cambiarValor(&myVar)
	fmt.Println("El nuevo valor de myVar es: ", myVar)
}