package main

import (
	"errors"
	"fmt"
)
	
func dividir(a, b int ) (result int, err error) {
	if b == 0 {
		err = errors.New("Operation is not valid: division by zero")
		return 0, err
	}
	result = a / b
	return result, nil
}

func main() {
	result, err := dividir(100, 5)
		if err != nil {
			fmt.Println("No valid operation: ", err)
	}
	fmt.Println(result)
}
