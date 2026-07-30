package main

import (
	"fmt"
	"myProject/basic_learning/structs/example1/user"
)

func main() {
	Human1 := user.Human{Name: "Andres", Age: 44, City: "Medellin"}
	Human2 := user.Human{Name: "Liliana", Age: 45, City: "Itagui"}

	fmt.Println("---------- Human 1 ----------")
	fmt.Println(Human1.Name)
	fmt.Println(Human1.Age)
	fmt.Println(Human1.City)

	fmt.Println("---------- Human 2 ----------")
	fmt.Println(Human2.Name)
	fmt.Println(Human2.Age)
	fmt.Println(Human2.City)
}
