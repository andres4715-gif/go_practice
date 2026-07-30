package main

import (
	"errors"
	"fmt"
)

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) (int, error) {
	if b > a {
		myError := errors.New("💊 NOT VALID")
		return 0, myError
	}
	return a - b, nil
}

func main() {
	totalAddition := Add(10, 20)
	fmt.Printf("The addition result is: %d\n", totalAddition)
	finalSubtraction, myNewError := Subtract(10, 5)
	fmt.Print("🎊 The final values is: ", finalSubtraction, "\n")

	if myNewError != nil {
		fmt.Printf("💥 Fatal error💥: %s", myNewError)
		return
	}
	fmt.Println("--- The final result is:", finalSubtraction)
}
