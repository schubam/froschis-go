package main

import (
	"fmt"

	"github.com/schubam/froschis-go/internal/game"
)

func main() {
	game := game.NewGame(2)

	fmt.Println(game.Print())
	game.Play()
}
