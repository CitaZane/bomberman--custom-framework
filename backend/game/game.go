package game

// holds game state to send it to all players
type GameState struct {
	Players []Player `json:"players"`
	// created bool
}

func (g *GameState) FindPlayer(name string) int {
	for index, player := range g.Players {
		if player.Name == name {
			return index
		}
	}

	return -1
}