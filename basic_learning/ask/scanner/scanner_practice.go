package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // This is to type a response

	apartmentID := ask(scanner, "Apartment id: ")
	name := ask(scanner, "Person name: ")
	dateText := ask(scanner, "Entry date (YYYY-MM-DD): ")

	date, err := time.Parse("2006-01-02", dateText)
	if err != nil {
		fmt.Println("❌ Invalid date: ", err)
		return
	}
	fmt.Println(apartmentID, name, date)
}

// ask prints a prompt and returns the next trimmed line from stdin.
func ask(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
