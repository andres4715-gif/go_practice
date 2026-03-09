package main

import "fmt"

func main() {

	// Forma 1 de declarar un map
	var myMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	myMap["four"] = 4
	fmt.Println(myMap["four"])

	// Forma 2 de declarar un map
	myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println(myMap["one"])

	// Como se edita un valor de un map 
	myMap["three"] = 89
	fmt.Println(myMap["three"])

	// Como se elimina un valor de un map
	fmt.Print("----- Eliminando un valor de un map -----\n")
	delete(myMap, "three") // El valor de three fue eliminado y ya no se puede acceder
	fmt.Println(myMap)

	// iterar sobre un mapa: 
	var myMap2 = make(map[string]int)
	myMap2["four"] = 4
	myMap2["five"] = 5
	myMap2["six"] = 6
	myMap2["seven"] = 7
	myMap2["eight"] = 8

	fmt.Print("----- Map a trabajar -----\n")
	fmt.Println(myMap2)

	// Como se itera en el map
	for _, value := range myMap2 {
		fmt.Print(value)
	}
}
