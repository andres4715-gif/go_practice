package main

import (
	"fmt"
	"time"
)

type tarjetaRegalo struct {
	valor        int
	dia          int
	mes          int
	year         int
	negocio      string
	dondeRedimir []string
}

func validarRegalo(newTarjetaRegalo tarjetaRegalo) (bool, error) {
	if newTarjetaRegalo.mes < 1 || newTarjetaRegalo.mes > 12 {
		return false, fmt.Errorf("mes inválido: %d", newTarjetaRegalo.mes)
	}

	hoy := time.Now()
	fechaVencimiento := time.Date(
		newTarjetaRegalo.year,
		time.Month(newTarjetaRegalo.mes),
		newTarjetaRegalo.dia,
		23, 59, 59, 0, time.Local,
	)

	if hoy.Before(fechaVencimiento) {
		return true, nil
	}
	return false, nil
}

func main() {
	myTarjetaRegalo := tarjetaRegalo{
		valor:        1000,
		dia:          17,
		mes:          2,
		year:         2030,
		negocio:      "Amazon",
		dondeRedimir: []string{"Santa Fe", "Oviedo", "La Central"},
	}

	estadoRegalo, err := validarRegalo(myTarjetaRegalo)
	if err != nil {
		fmt.Println("Hubo un error:", err)
	}

	fmt.Println("Puede reclamar el regalo? : ", estadoRegalo)
}
