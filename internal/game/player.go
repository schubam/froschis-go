package game

type gamepiece struct {
	card
	isFaceUp bool
}

type player struct {
	id        int
	boardSize int
	board     []*gamepiece
}

func NewPlayers(num int, deck deck) []player {
	ps := make([]player, num)

	for i := 0; i < num; i++ {
		drawn, _ := deck.draw(8)
		gps := make([]*gamepiece, 8)
		for j, c := range drawn {
			gps[j] = &gamepiece{card: c, isFaceUp: false}
		}

		ps[i] = player{
			id:        i,
			boardSize: 8,
			board:     gps,
		}
	}
	return ps
}
