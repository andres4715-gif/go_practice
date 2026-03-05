package main

import (
	"fmt"
	"log"
)
	
func dividir(a, b int ) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("The dividend is %d and the divisor value is %d ", a, b)
		return 0, err
	}
	result = a / b
	return result, nil
}

func main() {
	dividend := 100
	divisor := 10

	result, err := dividir(dividend, divisor)
		if err != nil {
			log.Fatalf("🚨 No valid operation: %s", err)
	}
	fmt.Println(result)
}
