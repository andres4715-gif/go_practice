package recorrer

import "fmt"

func RecorrerSlice(data []int) {
	for _, value := range data {
		fmt.Println(value)
	}
}
