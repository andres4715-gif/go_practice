package main

import (
	"fmt"
)

func contains(siblings []string, value string) bool {
	for _, element := range siblings {
		if element == value {
			return true
		}
	}
	return false
}

func main() {
	myData := contains([]string{"Andres", "Daniel", "Liliana"}, "Carlos") // Must print false, cos "Carlos" does not exist in slice siblings
	fmt.Println(myData)
}
