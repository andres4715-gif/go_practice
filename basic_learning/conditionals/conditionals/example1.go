package main

import (
	"fmt"
	"time"
)

type GiftCard struct {
	price    int
	day      int
	month    int
	year     int
	business string
	redeem   []string
}

func verifyGift(newGiftCard GiftCard) (bool, error) {
	if newGiftCard.month < 1 || newGiftCard.month > 12 {
		return false, fmt.Errorf("Invalid month: %d", newGiftCard.month)
	}

	today := time.Now()
	expirationDay := time.Date(
		newGiftCard.year,
		time.Month(newGiftCard.month),
		newGiftCard.day,
		23, 59, 59, 0, time.Local,
	)

	if today.Before(expirationDay) {
		return true, nil
	}
	return false, nil
}

func main() {
	myGiftCard := GiftCard{
		price:    1000,
		day:      17,
		month:    2,
		year:     2030,
		business: "Amazon",
		redeem:   []string{"Santa Fe", "Oviedo", "La Central"},
	}

	myGitStatus, err := verifyGift(myGiftCard)
	if err != nil {
		fmt.Println("Getting some Errors:", err)
	}

	fmt.Println("Give the gift? : ", myGitStatus)
}
