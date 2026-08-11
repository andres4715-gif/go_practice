package main

import "fmt"

type holderInfo struct {
	Name     string
	LastName string
	InUse    bool
	Id       string
}

type Card struct {
	Number     int
	HolderName string
	InUse      bool
	Id         string
}

func main() {
	holders := []holderInfo{
		{"Andres", "Palomino", true, "1234"},
		{"Maria", "Estrada", false, "5678"},
		{"Carlos", "Rios", false, "2345"},
		{"Luis", "Macarena", false, "6543"},
		{"Juan", "Moreno", false, "65438"},
		{"Andrea", "Pol", false, "787878"},
	}

	cards := make([]*Card, 0, len(holders)) // Making a slice

	for i := 1; i <= len(holders); i++ {
		cards = append(cards, &Card{
			Number:     i,
			HolderName: holders[i-1].Name,
			InUse:      holders[i-1].InUse,
			Id:         holders[i-1].Id,
		})
		fmt.Printf("%+v\n", cards[i-1])
	}
}
