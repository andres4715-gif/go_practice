package account

import "fmt"

// Pointer receiver: modifies the balance of the ORIGINAL account, not a copy
func (c *Account) Deposit(amount int) int {
	c.Balance += amount
	fmt.Println("Success Operation, new balance:", c.Balance)
	return c.Balance
}
