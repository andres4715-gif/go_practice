package pool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// helper: fecha valida para las pruebas
func sampleDate() time.Time {
	return time.Date(2026, time.August, 10, 0, 0, 0, 0, time.UTC)
}

func TestNewApartment(t *testing.T) {
	// apartamento valido: debe tener 4 tarjetas libres
	a, err := NewApartment("101")
	require.NoError(t, err)
	require.Len(t, a.Cards, CardsPerApartment)
	for i, card := range a.Cards {
		assert.Equal(t, i+1, card.Number) // numeradas 1..4
		assert.True(t, card.IsFree())     // todas libres al inicio
	}

	// id vacio: debe fallar
	_, err = NewApartment("")
	require.Error(t, err)
}

// helper: pool con un apartamento listo
func newPoolWithApartment(t *testing.T, id string) *Pool {
	t.Helper()
	a, err := NewApartment(id)
	require.NoError(t, err)
	p := NewPool()
	p.AddApartment(a)
	return p
}

func TestDeliverCard_HappyPath(t *testing.T) {
	p := newPoolWithApartment(t, "101")
	date := sampleDate()

	card, err := p.DeliverCard("101", "Ana", date)
	require.NoError(t, err)
	assert.True(t, card.InUse)
	assert.Equal(t, "Ana", card.HolderName)
	assert.Equal(t, date, card.EntryDate)
}

func TestDeliverCard_Invalid(t *testing.T) {
	casos := []struct {
		nombre      string
		apartmentID string
		name        string
		date        time.Time
	}{
		{"apartamento inexistente", "999", "Ana", sampleDate()},
		{"sin nombre", "101", "", sampleDate()},
		{"fecha cero", "101", "Ana", time.Time{}},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			p := newPoolWithApartment(t, "101")
			_, err := p.DeliverCard(c.apartmentID, c.name, c.date)
			require.Error(t, err)
		})
	}
}

func TestDeliverCard_MaxFourPeople(t *testing.T) {
	p := newPoolWithApartment(t, "101")
	date := sampleDate()

	// las 4 primeras entregas deben funcionar
	for i := 0; i < CardsPerApartment; i++ {
		_, err := p.DeliverCard("101", "persona", date)
		require.NoError(t, err)
	}

	// la 5a persona no puede entrar: no quedan tarjetas
	_, err := p.DeliverCard("101", "persona", date)
	require.Error(t, err)
}

func TestReturnCard(t *testing.T) {
	p := newPoolWithApartment(t, "101")
	date := sampleDate()

	// ocupamos las 4 tarjetas
	for i := 0; i < CardsPerApartment; i++ {
		_, err := p.DeliverCard("101", "persona", date)
		require.NoError(t, err)
	}

	// devolvemos la tarjeta 1 -> queda libre otra vez
	require.NoError(t, p.ReturnCard("101", 1))

	// ahora una nueva persona si puede entrar
	_, err := p.DeliverCard("101", "nueva", date)
	require.NoError(t, err)
}

func TestReturnCard_Invalid(t *testing.T) {
	p := newPoolWithApartment(t, "101")

	// tarjeta que ya esta libre
	require.Error(t, p.ReturnCard("101", 1))
	// apartamento inexistente
	require.Error(t, p.ReturnCard("999", 1))
	// numero de tarjeta inexistente
	require.Error(t, p.ReturnCard("101", 99))
}
