package main

import (
	"fmt"
	"log"
)

type carro struct {
	Modelo string
	Marca  string
	Anio   int64
	Precio int64
}

type casa struct {
	ciudad       string
	departamento string
	precio       int64
	tipo         string
}

func datosCarro(modelo string, version string, marca string, anio int64, precio int64) (carro, error) {
	if modelo == "" || version == "" || marca == "" || anio <= 0 || precio <= 0 {
		return carro{}, fmt.Errorf("faltan datos obligatorios o hay valores negativos")
	}
	// Si todo está bien, "armamos" nuestro carrito con los datos
	nuevoCarro := carro{
		Modelo: modelo + " " + version, // Aprovechamos para unir estos dos
		Marca:  marca,
		Anio:   anio,
		Precio: precio,
	}

	return nuevoCarro, nil
}

func datosCasa(ciudad string, departamento string, precio int64, tipo string) (string, casa, error) {
	if ciudad == "" || precio <= 0 || tipo == "" || tipo != "casa" && tipo != "apartamento" {
		return ciudad, casa{}, fmt.Errorf("Faltan datos mandatorios o estan erroneos")
	}

	myCasa := casa{
		precio:       precio,
		ciudad:       ciudad,
		departamento: departamento,
		tipo:         tipo,
	}
	return ciudad, myCasa, nil
}

func main() {
	myCarro, err := datosCarro("3", "Gran turing", "Mazda", 2017, 30000)
	if err != nil {
		log.Fatalf("🚨 No valid operation: %s", err)
	}
	fmt.Printf("Carro registrado: %s %s, Precio: $%v\n", myCarro.Marca, myCarro.Modelo, myCarro.Precio)

	_, myNuevaCasaPepito, err := datosCasa("Caldas", "Antioquia", 1000, "apartamento") // <--- APlicando el concepto del agujero negro
	if err != nil {
		log.Fatalf("No es posible registrar la casa %s", err)
	}
	fmt.Printf("%s registrada en %s departamento de %s con precio de $%d \n", myNuevaCasaPepito.tipo, myNuevaCasaPepito.ciudad, myNuevaCasaPepito.departamento, myNuevaCasaPepito.precio)
}
