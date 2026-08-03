package main

import (
	"fmt"
)

func main() {
	// 1. fmt.Print - Prints without a line break
	fmt.Print("Hello ")
	fmt.Print("Colombia")
	fmt.Print("\n") // manual line break

	// 2. fmt.Println - Prints with a line break
	fmt.Println("Print with a line break")
	fmt.Println("Second line")

	// 3. fmt.Printf - Prints with format like C
	myName := "Juan"
	MyAge := 25
	fmt.Printf("My name is %s and I am %d years old\n", myName, MyAge)

	// 4. More format examples
	initialCost := 15500.75
	fmt.Printf("The cost is: $%.2f dollars\n", initialCost)
	fmt.Printf("The binary is: %b\n", 42)        // Binary
	fmt.Printf("The Hexadecimal: %x\n", 255)     // Hexadecimal
	fmt.Printf("The Percentage: %.1f%%\n", 85.5) // Percentage

	// 5. fmt.Sprintf - String format without printing
	message := fmt.Sprintf("Use: %s, City: %s", "Ana", "Medellín")
	fmt.Println(message)

	message2 := fmt.Sprintf("My name is: %s, I am: %d, old and my id is %s", "Andres", 44, "9999999")
	fmt.Println(message2)

	var name string = "Andres"
	var age int = 44
	var id string = "9898777"
	fmt.Printf("My name is: %s, I am %d old and mi id is %s", name, age, id)
	fmt.Printf(" \n")
	var a, b, c int = 1, 2, 3
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
