package main

import (
	"fmt"
	"myProject/basic_learning/functions/usingStruct/structExample"
)

func main() {
	person1 := &structExample.Person{Name: "Andres"}
	fmt.Println("The House owner is:", person1.Name)

	person1.SayHello("Oscar")
}
