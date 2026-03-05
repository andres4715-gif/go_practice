package main

import (
"fmt"
) 

func main() {
	// 1. fmt.Print - Imprime sin salto de línea
	fmt.Print("Hola ")
	fmt.Print("Colombia")
	fmt.Print("\n") // Salto de línea manual

	// 2. fmt.Println - Imprime con salto de línea automático
	fmt.Println("Este es un salto automático")
	fmt.Println("Segunda línea")

	// 3. fmt.Printf - Imprime con formato (como printf en C)
	nombre := "Juan"
	edad := 25
	fmt.Printf("Mi nombre es %s y tengo %d años\n", nombre, edad)

	// 4. Más ejemplos de formato
	precio := 15500.75
	fmt.Printf("El precio es: $%.2f pesos\n", precio) // 2 decimales
	fmt.Printf("En binario: %b\n", 42)                // Binario
	fmt.Printf("En hexadecimal: %x\n", 255)           // Hexadecimal
	fmt.Printf("Porcentaje: %.1f%%\n", 85.5)          // Porcentaje

	// 5. fmt.Sprintf - Formatear a string sin imprimir
	mensaje := fmt.Sprintf("Usuario: %s, Ciudad: %s", "Ana", "Medellín")
	fmt.Println(mensaje)

	mensaje2 := fmt.Sprintf("Mi nombre es: %s, tengo: %d, y mi cedula es %s", "Andres", 44, "71376060")
	fmt.Println(mensaje2)

	var name string = "Andres"
	var age int = 44
	var cedula string = "71376060"
	fmt.Printf("Mi nombre es: %s, tengo: %d, y mi cedula es %s", name, age, cedula)
	fmt.Printf(" \n")
	var a, b, c int = 1, 2, 3
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
