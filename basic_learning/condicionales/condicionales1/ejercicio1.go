package main

import "fmt"

func main() {
	edad := 22
	nacionalidad := "Colombia"

	if edad >= 18 {
		fmt.Println("✅ Permitido el acceso")
	} else {
		fmt.Println("❌ Negado el acceso")
	}

	if edad >= 18 && nacionalidad == "Colombia" {
		fmt.Println("✅ Es colombiano y mayor de edad -> Permitido el acceso")
	} else if edad < 18 && nacionalidad != "Colombia" {
		fmt.Println("❌ No es colombiano y menor de edad -> Negado el acceso")
	} else {
		fmt.Println("Es del pais: ", nacionalidad, " y tiene ", edad, " años")
	}
}
