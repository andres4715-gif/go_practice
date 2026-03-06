package main

import "fmt"


func main() {

	fmt.Println("---------- For basico -----------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------- For con break  -----------")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("---------- For con continue -----------")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("---------- While -----------")
	i := 0
	for i < 6 {
		fmt.Println(i)
		i++
	}

	fmt.Println("---------- Do while -----------")
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 6 {
			break
		}
	}

	fmt.Println("---------- Recoriendo un map ----------- -----------")
	myMap2 := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	} 
	for k, v := range myMap2 {
		fmt.Println("El numero para el key: ", k, " es: ", v)
	}
}