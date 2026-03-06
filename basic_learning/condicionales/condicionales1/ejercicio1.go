package main

import "fmt"

func main() {
	edad := 22
	nacionalidad := "Colombia"

	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	} else {
		fmt.Println("Es menor de edad")
	}

	if edad >= 18 && nacionalidad == "Colombia" {
		fmt.Println("Es colombiano y mayor de edad")
	} else if edad < 18 && nacionalidad != "Colombia" {
		fmt.Println("No es colombiano y menor de edad")
	} else {
		fmt.Println("Es del pais: ", nacionalidad, " y tiene ", edad, " años")
	}
}
