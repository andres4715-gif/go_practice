package main

import "myProject/basic_learning/pointers/unit_administration/account"

func main() {
	info := account.OwnerInfo{
		OwnerName:         "Andres",
		OwnerLastName:     "Palomino",
		Apartments:        []string{"454", "987"},
		Balance:           90000,
		ParkingLots:       []string{"98d", "77b", "98u", "69h"},
		CarLicensePlates:  []string{"ABC-123", "EFG-456"},
		BikeLicensePlates: []string{"HIJ-789"},
	}

	// Calling the constructor = NewOwnerAccount()
	a := account.NewOwnerAccount("user1", info)

	// Adding a new Owner from AddOwner
	a.AddOwner("user2", account.OwnerInfo{
		OwnerName:         "Liliana",
		OwnerLastName:     "Palmas",
		Apartments:        []string{"454"},
		Balance:           70000,
		ParkingLots:       []string{"98d"},
		CarLicensePlates:  []string{"ABC-123"},
		BikeLicensePlates: []string{},
	})

	a.AddOwner("user3", account.OwnerInfo{
		OwnerName:         "Daniel",
		OwnerLastName:     "Rosales",
		Apartments:        []string{"875"},
		Balance:           700000,
		ParkingLots:       []string{"52d"},
		CarLicensePlates:  []string{"ZFK-987"},
		BikeLicensePlates: []string{},
	})

	// Printing list of users
	a.ListOwners()
}
