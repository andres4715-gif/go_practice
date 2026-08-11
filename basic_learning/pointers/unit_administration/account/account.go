package account

import "time"

type Account struct {
	Owners map[string]OwnerInfo
}

// Constructor
func NewOwnerAccount(ownerId string, info OwnerInfo) *Account {
	return &Account{
		Owners: map[string]OwnerInfo{
			ownerId: info,
		},
	}
}

// RentSocialRoom tries to rent the social room today to the given owner.
// It returns false if the owner does not exist or if someone else already rented
// has the social room rented today (only one owner can have it at a time).
func (a *Account) RentSocialRoom(ownerId string) bool {
	// does the owner exist?
	owner, exists := a.Owners[ownerId]
	// ⬆️⬆️⬆ This is a map and the map returns the element and a bool value

	if !exists {
		return false
	}

	// is the social room available today?
	for _, info := range a.Owners {
		if info.HasRentedSocialRoomToday() {
			return false // occupied can't be rented
		}
	}

	// Free -> we assign it to this owner
	// We reassign it cos the map stores a COPY of the struct
	owner.SocialRoomRent = time.Now()
	a.Owners[ownerId] = owner
	return true
}
