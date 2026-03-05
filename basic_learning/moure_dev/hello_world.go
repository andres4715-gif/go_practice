package main

import "fmt"
import "reflect"

func main() {
	// Different ways to declare variables
	var country string
	var age int = 66
	var tall float64 = 1.90
	var currentYear = 2026
	var bornIn = 1981
	myString := "Andres" // Automatic variable declaration
	country = "Colombia"
	var isMarried bool = true
	const PI float64 = 3.1416
	var myMap = make(map[string]int)
	myMap["key"] = 1
	myMap["key2"] = 20
	myMap["key3"] = 3
	myMap["key4"] = 4
	myMap["key5"] = 5
	var myArray = make([]int, 6)
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50
	myArray[5] = 60

	type employee struct {
		name string
		age  int
		id string
	}

	fmt.Println("The name is: ", myString)
	fmt.Println("The age is: ", age)
	fmt.Println("Hello, World! with context")
	fmt.Println("The country is: ", country)
	fmt.Println("The age is ", currentYear-bornIn)
	fmt.Println("The type for age is: ", reflect.TypeOf(age))
	fmt.Println("The type for country is: ", reflect.TypeOf(country))
	fmt.Println("The tall is: ", tall)
	fmt.Println("The type for tall is: ", reflect.TypeOf(tall))
	fmt.Println("The type for PI is: ", reflect.TypeOf(PI))
	fmt.Printf("The value for PI: %.4f\n", (PI))
	fmt.Println("The value for map position 2 is: ", myMap["key2"])

	if !isMarried {
		fmt.Println("The person is married")
	} else {
		fmt.Println("The person is not married")
	}

	if PI >= 2.10 {
		fmt.Println("The PI is greater than 2.10")
	} else {
		fmt.Println("The PI is not greater than 2.10")
	}

	for index, value := range myArray {
		fmt.Println("The index is ", index, " the value is: ", value)
	}

	if !myFunction("Andres", 47) {
		fmt.Println("The function returned true")
	}

	var myEmployee employee = employee{"Liliana", 44, "123456"}	
	fmt.Println("The employee name is: ", myEmployee.name)
	fmt.Println("The employee age is: ", myEmployee.age)
	fmt.Println("The employee id is: ", myEmployee.id)
}

func myFunction(name string, age int) bool {
	fmt.Println("The name is:", name+" with age: ", age)
	if name == "Andres" && age == 44 {
		return true
	} else {
		return false
	}
}
