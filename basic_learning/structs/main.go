package main

import (
	"fmt"
	"myProject/basic_learning/structs/ejercicio1/persona"
)

func main() {
	p1 := persona.Humano{Nombre: "Andres", Edad: 44, Ciudad:"Medellin"}
	p2 := persona.Humano{Nombre: "Lilina", Edad: 45, Ciudad: "Itagui"}

    fmt.Println("---------- Persona 1 ----------")
	fmt.Println(p1.Nombre)
	fmt.Println(p1.Edad)
	fmt.Println(p1.Ciudad)

	fmt.Println("---------- Persona 2 ----------")
	fmt.Println(p2.Nombre)
	fmt.Println(p2.Edad)
	fmt.Println(p2.Ciudad)

}