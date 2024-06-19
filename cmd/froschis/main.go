package main

import (
	"fmt"

	"github.com/schubam/froschis-go/internal/game"
)

func main() {
	deck := game.NewDeck()
	fmt.Println(deck)
}
