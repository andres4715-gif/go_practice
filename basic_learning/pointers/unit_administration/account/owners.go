package account

import (
	"fmt"
	"time"
)

type OwnerInfo struct {
	OwnerName         string
	OwnerLastName     string
	Apartments        []string
	Balance           int
	ParkingLots       []string
	CarLicensePlates  []string
	BikeLicensePlates []string
	Debt              int
	SocialRoomRent    time.Time // date the owner rented the social room (zero value = not rented)
}

func (a *Account) AddOwner(ownerId string, info OwnerInfo) {
	a.Owners[ownerId] = info
}

// ListOwners prints a summary of every owner in the account
func (a *Account) ListOwners() {
	flag := 1
	for userId, userInfo := range a.Owners {
		fmt.Printf("%d - [%s] %s %s | Initial Balance: %d | Debt: %d | HasDebt: %t | Final Balance: %d | Social room Rented: %t\n",
			flag, userId, userInfo.OwnerName, userInfo.OwnerLastName, userInfo.Balance,
			userInfo.Debt, userInfo.HasDebt(), userInfo.finalBalance(), userInfo.HasRentedSocialRoomToday())
		flag = flag + 1
	}
}

// HasDebt reports whether the given owner has any pending debt
func (info OwnerInfo) HasDebt() bool {
	return info.Debt > 0
}

func (info OwnerInfo) finalBalance() int {
	newBalance := 0
	if info.HasDebt() == true {
		newBalance = info.Balance - info.Debt
		return newBalance
	}
	return info.Balance
}

// HasRentedSocialRoomToday reports whether the owner has the social room rented for today
func (info OwnerInfo) HasRentedSocialRoomToday() bool {
	now := time.Now()
	rent := info.SocialRoomRent
	return rent.Year() == now.Year() &&
		rent.Month() == now.Month() &&
		rent.Day() == now.Day() // if match all items it returns true
}
