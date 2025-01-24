package game

type GamePhase int

const (
	Setup GamePhase = iota
	PlayerTurn
	EndTurn
	EndGame
)

type GameState struct {
	Phase       GamePhase
	CurrentTurn int
	Deck        deck
	DiscardPile []card
	Players     []player
}

func NewGameState(numPlayers int) *GameState {
	deck := NewDeck()

	return &GameState{
		Phase:       Setup,
		CurrentTurn: 1,
		Deck:        deck,
		DiscardPile: []card{},
		Players:     NewPlayers(numPlayers, deck),
	}
}
