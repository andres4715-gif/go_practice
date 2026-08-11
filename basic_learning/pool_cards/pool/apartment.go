package pool

import "fmt"

// CardsPerApartment is the number of pool cards every apartment owns.
// It also caps how many people from that apartment can be inside at once.
const CardsPerApartment = 4

// Apartment owns a fixed set of pool-access cards.
type Apartment struct {
	ID    string
	Cards []*Card
}

// NewApartment builds an apartment with CardsPerApartment free cards (numbered 1..4).
// It returns an error when the id is empty.
func NewApartment(id string) (*Apartment, error) {
	if id == "" {
		return nil, fmt.Errorf("apartment id is required")
	}

	// We use pointers so the cards can be updated in place (no map-copy gotcha).
	cards := make([]*Card, 0, CardsPerApartment)
	for i := 1; i <= CardsPerApartment; i++ {
		cards = append(cards, &Card{Number: i})
	}

	return &Apartment{ID: id, Cards: cards}, nil
}
