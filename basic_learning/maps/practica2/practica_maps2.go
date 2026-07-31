package main

import "fmt"

func main() {

	// Form 1 for map declaration
	var myMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	myMap["four"] = 4
	fmt.Println(myMap["four"])

	// Forma 2 for map declaration
	myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println(myMap["one"])

	// How to edit a map value
	myMap["three"] = 89
	fmt.Println(myMap["three"])

	// How to delete a map value
	fmt.Print("----- Removing a map value -----\n")
	delete(myMap, "three") // The value of three was removed an it is not available
	fmt.Println(myMap)

	// Map iteration
	var myMap2 = make(map[string]int)
	myMap2["four"] = 4
	myMap2["five"] = 5
	myMap2["six"] = 6
	myMap2["seven"] = 7
	myMap2["eight"] = 8

	fmt.Print("----- Map to work on -----\n")
	fmt.Println(myMap2)

	// Map iteration v2
	for _, value := range myMap2 {
		fmt.Print(value)
	}
}
