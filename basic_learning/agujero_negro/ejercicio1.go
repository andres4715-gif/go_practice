package main

import (
	"fmt"
	"log"
)

type Carro struct {
	Modelo string
	Marca  string
	Anio   int64
	Precio int64
}

func datosCarro(modelo string, version string, marca string, anio int64, precio int64) (Carro, error) {
	if modelo == "" || version == "" || marca == "" || anio <= 0 || precio <= 0 {
		return Carro{}, fmt.Errorf("faltan datos obligatorios o hay valores negativos")
	}
	// Si todo está bien, "armamos" nuestro carrito con los datos
	nuevoCarro := Carro{
		Modelo: modelo + " " + version, // Aprovechamos para unir estos dos
		Marca:  marca,
		Anio:   anio,
		Precio: precio,
	}

	return nuevoCarro, nil
}

func main() {
	myCarro, err := datosCarro("3", "Gran turing", "Mazda", 2017, 30000)
	if err != nil {
		log.Fatalf("🚨 No valid operation: %s", err)
	}
	fmt.Printf("Carro registrado: %s %s, Precio: $%v\n", myCarro.Marca, myCarro.Modelo, myCarro.Precio)
}
