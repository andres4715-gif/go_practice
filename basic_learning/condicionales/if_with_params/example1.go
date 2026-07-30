package main

import (
	"fmt"
)

func main() {
	fmt.Println("😎 Using if with a short declaration")
	if total, err := divide(100, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide result is", total)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
