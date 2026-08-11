package main

import (
	"fmt"
	"log"
)

type car struct {
	Model string
	Brand string
	Year  int64
	Price int64
}

type house struct {
	City       string
	Department string
	Price      int64
	Type       string
}

func carDataConstructor(model string, version string, brand string, year int64, price int64) (car, error) {
	if model == "" || version == "" || brand == "" || year <= 0 || price <= 0 {
		return car{}, fmt.Errorf("Missing car data required data or negative data")
	}
	// If everything is ok "make" our car with data
	NewCar := car{
		Model: model + " " + version, // We take the opportunity to combine these two (model and version)
		Brand: brand,
		Year:  year,
		Price: price,
	}

	return NewCar, nil
}

func houseDataConstructor(city string, department string, price int64, myType string) (house, error) {
	if city == "" || department == "" || price <= 0 || myType == "" || myType != "house" && myType != "apartment" {
		return house{}, fmt.Errorf("Missing house mandatory data or some errors")
	}

	NewHouse := house{
		City:       city,
		Department: department,
		Price:      price,
		Type:       myType,
	}
	return NewHouse, nil
}

// Check in this function if you don't use the variable is possible to use _
func main() {
	_, err := carDataConstructor("3", "Gran turing", "Mazda", 2017, 30000)
	if err != nil {
		log.Fatalf("🚨 No valid operation: %s", err)
	}
	fmt.Println("👷🏻‍♂️ The car is ready")
	// fmt.Printf("Register car: %s %s, Price: $%v\n", myCar.Brand, myCar.Model, myCar.Price)

	myNewHouse, err := houseDataConstructor("Caldas", "Antioquia", 1000, "apartment")
	if err != nil {
		log.Fatalf("Is not possible to register %s", err)
	}
	fmt.Printf("%s registered in %s department of %s with price $%d \n", myNewHouse.Type, myNewHouse.City, myNewHouse.Department, myNewHouse.Price)
}
