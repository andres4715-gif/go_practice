package account

import "fmt"

// Pointer receiver too, so it mutates the same account as Deposit
func (c *Account) Withdraw(amount int) (int, error) {
	if amount > c.Balance {
		fmt.Println("❌ Operation not allowed - Current balance", c.Balance)
		return 0, fmt.Errorf("😡 Insufficient balance")
	}

	c.Balance -= amount
	fmt.Println("Success Withdraw operation, new Balance:", c.Balance)
	return c.Balance, nil
}
