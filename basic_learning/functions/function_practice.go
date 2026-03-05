package main

import (
	"errors"
	"fmt"
)

func Suma(a, b int) int {
	return a + b
}

func Resta(a, b int) (int, error) {
	if b > a {
		myError := errors.New("💊 NOT VALID")
		return 0, myError
	}
	return a - b, nil
}

func main() {
	totalSuma := Suma(10, 20)
	fmt.Printf("The final Data is: %d\n", totalSuma)
	final, myNewError:= Resta(10, 5)
	fmt.Print("---", myNewError, "\n")
	fmt.Print("🎊 The final values is: ", final, "\n")

	if myNewError != nil {
		fmt.Printf("💥 Fatal error💥 : %s", myNewError)
		return
	}
	fmt.Println("--- The final values is:", final)
}
