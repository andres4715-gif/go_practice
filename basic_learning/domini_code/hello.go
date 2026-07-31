package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type employee struct {
	name string
	age  int
	id   string
}

func main() {
	// Declare consts
	const (
		PI = 3.1416
		RA = 89.988
	)

	// Expression constants
	const data = PI + RA

	// Print:
	fmt.Println("Hello World 2")
	fmt.Println("The number of PI: ", PI)
	fmt.Println("Total of data: ", data)

	// Assign a variable to another type from string to float 
	var age string = "44.4"
	var finalAge, _ = strconv.ParseFloat(age, 64)
	fmt.Println("The final age is: ", finalAge)
	fmt.Println("The typeOf is: ", reflect.TypeOf(finalAge))

	// Assigning one variable to another of a different type when they are numbers:
	var ageFloat float64 = 44.4
	var finalAge2 = int(ageFloat)
	fmt.Println("The final data is: ", ageFloat)
	fmt.Println("The typeOf data is: ", reflect.TypeOf(finalAge2))

	// Logical and relational operators
	var a int = 10
	var b int = 20

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	// comparison operators
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	/*
		Key difference: fixed size vs dynamic.
		Array ([3]string): fixed size, part of the type. [3]string and [4]string are different types, you cannot change the size.
		Slice ([]string): dynamic size, you can use append to add elements and it grows. Internally, a slice is an array + reference (pointer, len, cap).
		In practice you almost always use slices; arrays are rarely used.
	*/

	// Array
	fmt.Println("------------ Arrays ------------")
	var myArray [3]string
	myArray = [3]string{"Angular", "react", "vue"}
	fmt.Println(myArray[0])
	fmt.Println(myArray)
	myArray[0] = "AngularJS"
	fmt.Println(myArray)

	// Slice
	fmt.Println("------------ Slice 1 ------------")
	mySlice := []string{"Angular", "react", "vue", "svelte"}
	fmt.Println(mySlice[0])
	fmt.Println(mySlice)

	// Add a new element to the slice
	fmt.Println("------------ slice  2 ------------")
	mySlice2 := []string{}
	mySlice2 = append(mySlice2, "Angular_V2", "svelte_V2", "react_V2", "vue_V2", "sonar_v3")
	fmt.Println("The new slice is: ", mySlice2)
	fmt.Println("The length of the new slice is: ", len(mySlice2))

	// bring only a part of the elements that are in the slice:
	fmt.Println("------------ slice  3 ------------")
	fmt.Println("✅ mySlice: ", mySlice)
	frameworklist := mySlice[1:3]
	fmt.Println("🧵 The fragmented list is: ", frameworklist)

	fmt.Println("------------ slice  4 ------------")
	mySlice3 := []string{"Carro", "Casa", "Moto"}
	fmt.Println("✅ mySlice3: ", mySlice3)
	frameworklist2 := mySlice3[1:]
	fmt.Println("🧵 The fragmented list is: ", frameworklist2)

	// Know the number of elements in the array or slice
	fmt.Println("------------ len() ------------")
	longitud := len(mySlice)
	fmt.Printf("La longitud del slice es: %d\n", longitud)

	// struct
	fmt.Println("------------ struct ------------")
	var employeeUsage employee
	employeeUsage.name = "Andres"
	employeeUsage.age = 44
	employeeUsage.id = "71376060"

	fmt.Println("------------ Employee 1 ------------")
	fmt.Println("The name of the employee is:", employeeUsage.name)
	fmt.Println("The age of the employee is:", employeeUsage.age)
	fmt.Println("The id of the employee is:", employeeUsage.id)

	fmt.Println("------------ Employee 2 ------------")
	var employeeUsage2 employee = employee{"Liliana", 67, "67876654"}
	fmt.Println("The name of the employee2 is:", employeeUsage2.name)
	fmt.Println("The age of the employee2 is:", employeeUsage2.age)
	fmt.Println("The id of the employee2 is:", employeeUsage2.id)

	// for
	fmt.Println("------------ for ------------")
	fmt.Println("------------ for example 1 ------------")
	for i := 3; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------ for example 2 ------------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("------------ for example 3 ------------")
	// for key, value := range collection {...}
	for _, value := range mySlice2 { // To omit the key, the underscore _ is used
		fmt.Println("Printing a collection: ", value)
	}
}
