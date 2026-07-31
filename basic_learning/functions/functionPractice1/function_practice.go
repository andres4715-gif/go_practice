package main

import (
	"errors"
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) (int, error) {
	if b > a {
		myError := errors.New("💊 NOT VALID")
		return 0, myError
	}
	return a - b, nil
}

func main() {
	totalAddition := Add(10, 20)
	fmt.Printf("The addition result is: %d\n", totalAddition)

	result, err := Subtract(15, 5)
	if err != nil {
		fmt.Printf("💥 Fatal error💥: %s", err)
	}
	fmt.Print("🎊 The final values is: ", result, "\n")

	fmt.Println("--- The final result is:", result)
}
