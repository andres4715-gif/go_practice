package main

import (
	"fmt"
	"log"
)

func dividedBy(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("The divided is %d and the divisor value is %d ", a, b)
		return 0, err
	}
	result = a / b
	return result, nil
}

func main() {
	divided := 100
	divisor := 10

	result, err := dividedBy(divided, divisor)
	if err != nil {
		log.Fatalf("🚨 No valid operation: %s", err)
	}
	fmt.Println(result)
}
