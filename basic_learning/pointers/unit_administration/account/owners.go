package account

import "fmt"

type OwnerInfo struct {
	OwnerName         string
	OwnerLastName     string
	Apartments        []string
	Balance           int
	ParkingLots       []string
	CarLicensePlates  []string
	BikeLicensePlates []string
}

func (a *Account) AddOwner(ownerId string, info OwnerInfo) {
	a.Owners[ownerId] = info
}

// ListOwners prints a summary of every owner in the account
func (a *Account) ListOwners() {
	for id, info := range a.Owners {
		fmt.Printf("- [%s] %s %s | Balance: %d\n",
			id, info.OwnerName, info.OwnerLastName, info.Balance)
	}
}
