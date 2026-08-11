package pool

import (
	"fmt"
	"time"
)

// Pool manages the apartments that have access to the pool, keyed by apartment id.
type Pool struct {
	Apartments map[string]*Apartment
}

// NewPool returns an empty pool ready to receive apartments.
func NewPool() *Pool {
	return &Pool{Apartments: make(map[string]*Apartment)}
}

// AddApartment registers an apartment in the pool (stored by its id).
func (p *Pool) AddApartment(a *Apartment) {
	p.Apartments[a.ID] = a
}

// DeliverCard hands a free card of the apartment to a person at the pool entrance,
// recording their name and entry date. It returns an error when the apartment does
// not exist, when name/date are missing, or when all 4 cards are already in use
// (so at most 4 people can be inside at the same time).
func (p *Pool) DeliverCard(apartmentID, name string, date time.Time) (*Card, error) {
	apartment, exists := p.Apartments[apartmentID]
	if !exists {
		return nil, fmt.Errorf("apartment %q does not exist", apartmentID)
	}

	if name == "" || date.IsZero() {
		return nil, fmt.Errorf("name and entry date are required")
	}

	// Look for the first available card.
	for _, card := range apartment.Cards {
		if card.IsFree() {
			card.InUse = true
			card.HolderName = name
			card.EntryDate = date
			return card, nil
		}
	}

	return nil, fmt.Errorf("apartment %q has no available cards (max %d people)", apartmentID, CardsPerApartment)
}

// ReturnCard frees a card when the person leaves, so another person can use it.
// It returns an error when the apartment/card does not exist or the card is already free.
func (p *Pool) ReturnCard(apartmentID string, cardNumber int) error {
	apartment, exists := p.Apartments[apartmentID]
	if !exists {
		return fmt.Errorf("apartment %q does not exist", apartmentID)
	}

	for _, card := range apartment.Cards {
		if card.Number == cardNumber {
			if card.IsFree() {
				return fmt.Errorf("card %d of apartment %q is already free", cardNumber, apartmentID)
			}
			// Reset the card to its free state.
			card.InUse = false
			card.HolderName = ""
			card.EntryDate = time.Time{}
			return nil
		}
	}

	return fmt.Errorf("apartment %q has no card number %d", apartmentID, cardNumber)
}

// ListStatus prints every apartment and the state of its 4 cards.
func (p *Pool) ListStatus() {
	for _, apartment := range p.Apartments {
		fmt.Printf("Apartment %s:\n", apartment.ID)
		for _, card := range apartment.Cards {
			if card.IsFree() {
				fmt.Printf("  Card %d -> free\n", card.Number)
			} else {
				fmt.Printf("  Card %d -> %s (entry: %s)\n",
					card.Number, card.HolderName, card.EntryDate.Format("2006-01-02"))
			}
		}
	}
}
