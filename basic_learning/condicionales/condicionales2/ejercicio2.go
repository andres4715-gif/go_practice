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
    hoy := time.Now()

    // 1. Convertimos tus números enteros en una Fecha Real de Go
    // Usamos time.Month() para convertir tu int en el tipo que pide Go
    fechaVencimiento := time.Date(
        newTarjetaRegalo.year, 
        time.Month(newTarjetaRegalo.mes), 
        newTarjetaRegalo.dia, 
        23, 59, 59, 0, time.Local,
    )

    // 2. Comparamos las dos fechas directamente
    // ¿Hoy es ANTES de la fecha de vencimiento?
    if hoy.Before(fechaVencimiento) {
        return true, nil  // Sigue vigente
    } else {
        return false, nil // Ya venció
    }
}

func main() {
	myTarjetaRegalo := tarjetaRegalo{
		valor:        1000,
		dia:          17,
		mes:          3,
		year:         2026,
		negocio:      "Amazon",
		dondeRedimir: []string{"Santa Fe", "Oviedo", "La Central"},
	}

	estadoRegalo, err := validarRegalo(myTarjetaRegalo)
	if err != nil {
		fmt.Println("Hubo un error:", err)
	}	

	fmt.Println(estadoRegalo)
}
