package pool

import "time"

// Card is a single pool-access card that belongs to an apartment.
// A card can be handed to one person at a time.
type Card struct {
	Number     int       // 1..4, identifies the card within its apartment
	InUse      bool      // true while handed to a person
	HolderName string    // name of the person who currently holds it ("" when free)
	EntryDate  time.Time // entry date recorded at delivery (zero value = free)
}

// IsFree reports whether the card is available to be delivered.
func (c *Card) IsFree() bool {
	return !c.InUse
}
