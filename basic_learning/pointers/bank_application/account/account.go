package account

import "fmt"

type Account struct {
	Owner string
	Balance int
}

func NewAccount(owner string, initialBalance int) *Account {
	fmt.Println("Client name:", owner)
	fmt.Println("Client Current Balance:", initialBalance)
	return &Account{
		Owner: owner,
		Balance: initialBalance,
	}
}