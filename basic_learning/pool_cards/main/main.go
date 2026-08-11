package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"myProject/basic_learning/pool_cards/pool"
)

// This example simulates the pool entrance: each apartment owns 4 access cards,
// so at most 4 people can be inside at once. A card is delivered to a person
// (recording a name and an entry date) and returned when they leave.
func main() {
	p := pool.NewPool()

	// We start with one apartment that already has its 4 cards.
	apartment, err := pool.NewApartment("101")
	if err != nil {
		fmt.Println("could not create apartment:", err)
		return
	}
	p.AddApartment(apartment)

	// We use a Scanner (not fmt.Scan) so names with spaces are read fully.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Pool card entrance. Commands: deliver | return | list | quit")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			return // stdin closed
		}
		command := strings.TrimSpace(scanner.Text())

		switch command {
		case "deliver":
			apartmentID := ask(scanner, "Apartment id: ")
			name := ask(scanner, "Person name: ")
			dateText := ask(scanner, "Entry date (YYYY-MM-DD): ")

			date, err := time.Parse("2006-01-02", dateText)
			if err != nil {
				fmt.Println("invalid date, use YYYY-MM-DD")
				continue
			}

			card, err := p.DeliverCard(apartmentID, name, date)
			if err != nil {
				fmt.Println("could not deliver card:", err)
				continue
			}
			fmt.Printf("Card %d delivered to %s\n", card.Number, card.HolderName)

		case "return":
			apartmentID := ask(scanner, "Apartment id: ")
			numberText := ask(scanner, "Card number: ")

			number, err := strconv.Atoi(numberText)
			if err != nil {
				fmt.Println("invalid card number")
				continue
			}

			if err := p.ReturnCard(apartmentID, number); err != nil {
				fmt.Println("could not return card:", err)
				continue
			}
			fmt.Printf("Card %d returned\n", number)

		case "list":
			p.ListStatus()

		case "quit":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("unknown command. Use: deliver | return | list | quit")
		}
	}
}

// ask prints a prompt and returns the next trimmed line from stdin.
func ask(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
