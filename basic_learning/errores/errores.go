package main

import "fmt"
import "os"
	


func main() {
    contenido, err := leyendoArchivo("prueba.txt")
	if err != nil {
		fmt.Println("Error leyendo el archivo: ",  err)
		return
	}
	fmt.Println("El contenido del archivo es: ", string(contenido)) // Se debe convertir a string por que es un byte
}


func leyendoArchivo(archivoName string) ([]byte, error) {
	texto, err := os.ReadFile(archivoName)
	if err != nil {
		fmt.Println("Error leyendo el archivo: ",  err)
	return nil, err
	}
	return texto, nil
}