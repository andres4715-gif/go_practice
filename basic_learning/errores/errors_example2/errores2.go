package main

import (
	"errors"
	"fmt"
	"log"
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
	result, err := dividir(100, 0)
		if err != nil {
			log.Fatalf("🚨 No valid operation: %s", err)
	}
	fmt.Println(result)
}
