package game

import (
	"math/rand"
	"time"
)

type DrawError struct {
	Message string
}

func (e *DrawError) Error() string {
	return e.Message
}

var (
	ErrNotEnoughCards = &DrawError{Message: "not enough cards to draw"}
	ErrInvalidCount   = &DrawError{Message: "number of cards to draw must be greater than zero"}
)

type deck struct {
	cards []card
}

func NewDeck() deck {
	d := deck{cards: []card{
		ONE, ONE, ONE, ONE, ONE, ONE,
		TWO, TWO, TWO, TWO, TWO, TWO,
		THREE, THREE, THREE, THREE, THREE, THREE,
		FOUR, FOUR, FOUR, FOUR, FOUR, FOUR,
		FIVE, FIVE, FIVE, FIVE, FIVE, FIVE,
		SIX, SIX, SIX, SIX, SIX, SIX,
		SEVEN, SEVEN, SEVEN, SEVEN, SEVEN, SEVEN,
		EIGHT, EIGHT, EIGHT, EIGHT, EIGHT, EIGHT,
		GRIEBSCH,
		BOTTLE,
		SHOE,
		TOMATO,
		BALLERINA,
		GUITAR,
		WAVE,
		HOLIDAY,
	}}
	d.shuffle()
	return d
}

func (d *deck) draw(i int) ([]card, error) {
	var drawn []card
	if i <= 0 {
		return nil, ErrInvalidCount
	}

	if len(d.cards) < i {
		return nil, ErrNotEnoughCards
	}

	drawn = d.cards[:i]
	d.cards = d.cards[i:]
	return drawn, nil
}

func (d *deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
