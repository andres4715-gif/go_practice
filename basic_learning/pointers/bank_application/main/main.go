package main

import (
	"fmt"
	"myProject/basic_learning/pointers/bank_application/account"
)

func main() {
	fmt.Println("💰 💵 💰 .:: South Africa International BANK ::. 💰 💵 💰")
	c := account.NewAccount("John Doe", 5000)
	c.Deposit(1000)

	_, err := c.Withdraw(50000)
	if err != nil {
		fmt.Println("❌ Error:", err)
	}
}
