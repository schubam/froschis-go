package game

import (
	"time"
	"math/rand"
)

type Player struct {
	id       int
	Numcards int
}

type card string

const (
	ONE   card = "1"
	TWO   card = "2"
	THREE card = "3"
	FOUR  card = "4"
	FIVE  card = "5"
	SIX   card = "6"
	SEVEN card = "7"
	EIGHT card = "8"
	// Dead cards
	GRIEBSCH card = "griebsch"
	BOTTLE   card = "bottle"
	SHOE     card = "shoe"
	TOMATO   card = "tomato"
	// Froschis
	BALLERINA card = "ballerina"
	GUITAR    card = "guitar"
	WAVE      card = "wave"
	HOLIDAY   card = "holiday"
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

func (d *deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
