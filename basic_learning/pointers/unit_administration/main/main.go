package main

import (
	"fmt"
	"myProject/basic_learning/pointers/unit_administration/account"
)

func main() {
	userInfo := account.OwnerInfo{
		OwnerName:         "Andres",
		OwnerLastName:     "Palomin",
		Apartments:        []string{"454", "987"},
		Balance:           90000,
		ParkingLots:       []string{"98d", "77b", "98u", "69h"},
		CarLicensePlates:  []string{"ABC-123", "EFG-456"},
		BikeLicensePlates: []string{"HIJ-789"},
		Debt:              500,
	}

	// Calling the constructor = NewOwnerAccount()
	a := account.NewOwnerAccount("user1", userInfo)

	// Adding a new Owner from AddOwner
	a.AddOwner("user2", account.OwnerInfo{
		OwnerName:         "Liliana",
		OwnerLastName:     "Palmas",
		Apartments:        []string{"454"},
		Balance:           70000,
		ParkingLots:       []string{"98d"},
		CarLicensePlates:  []string{"ABC-123"},
		BikeLicensePlates: []string{},
		Debt:              5000,
	})

	a.AddOwner("user3", account.OwnerInfo{
		OwnerName:         "Daniel",
		OwnerLastName:     "Rosales",
		Apartments:        []string{"875"},
		Balance:           60000,
		ParkingLots:       []string{"52d"},
		CarLicensePlates:  []string{"ZFK-987"},
		BikeLicensePlates: []string{},
		Debt:              0,
	})

	// Renting the social room: only one owner can have it at a time.
	fmt.Println(" ⏰ [user1] rents social room:", a.RentSocialRoom("user1")) // true  -> he rented it first
	fmt.Println(" ⏰ [user2] rents social room:", a.RentSocialRoom("user2")) // false -> is already occupied today
	fmt.Println(" ⏰ [user3] rents social room:", a.RentSocialRoom("user3")) // false -> is already occupied today

	// Printing list of users
	a.ListOwners()
}
