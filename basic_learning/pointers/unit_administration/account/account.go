package account

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
