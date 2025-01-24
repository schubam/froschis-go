package game

import "fmt"

type gameplayer interface {
	Play()
	Print() string

	drawCard(player int) (*gamepiece, error)
	flipCard(player, position int) (*gamepiece, error)
	putCard(player, position int, card *gamepiece)
	discardCard(player, card *gamepiece)
	determineStartPlayer() int
}

type Game struct {
	State *GameState
}

func NewGame(numPlayers int) gameplayer {
	return &Game{
		State: NewGameState(numPlayers),
	}
}

func (g *Game) Play() {
}

func (g *Game) Print() string {
	return fmt.Sprintf("%+v", g.State)
}

func (g *Game) flipCard(player int, position int) (*gamepiece, error) {
	c := g.State.Players[player].board[position]
	g.State.Players[player].board[position] = nil
	return c, nil
}

func (g *Game) drawCard(player int) (*gamepiece, error) {
	return &gamepiece{}, fmt.Errorf("oops")
}

func (g *Game) putCard(player, position int, card *gamepiece) {
}

func (g *Game) discardCard(player, card *gamepiece) {
}

func (g *Game) determineStartPlayer() int {
	return 0
}
