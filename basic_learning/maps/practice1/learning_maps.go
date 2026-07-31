package main

import "fmt"

func main() {
	myMap := map[string] int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// Adding a new key and value within the map
	myMap["four"] = 4

	fmt.Println("🚀 Data of myMap: ", myMap) // using a map, order doesn't matter
	fmt.Println(myMap["one"])
	fmt.Println(myMap["two"])
	fmt.Println(myMap["three"])
	fmt.Println(myMap["four"])

	// Delete a key and value
	delete(myMap, "one") // To do it just use the key
	fmt.Println("👷🏻‍♂️ New data after remove: ", myMap)
}
