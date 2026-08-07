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
	Debt              int
}

func (a *Account) AddOwner(ownerId string, info OwnerInfo) {
	a.Owners[ownerId] = info
}

// ListOwners prints a summary of every owner in the account
func (a *Account) ListOwners() {
	for id, info := range a.Owners {
		fmt.Printf("- [%s] %s %s | Balance: %d | Debt: %d | HasDebt: %t\n",
			id, info.OwnerName, info.OwnerLastName, info.Balance, info.Debt, info.HasDebt())
	}
}

// HasDebt reports whether the given owner has any pending debt
func (info OwnerInfo) HasDebt() bool {
	return info.Debt > 0
}
