package traverse

import "fmt"

func TraverseSlice(data []int) {
	for _, value := range data {
		fmt.Println(value)
	}
}
