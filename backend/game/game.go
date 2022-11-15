package game

// holds game state to send it to all players
type GameState struct {
	Players   []Player  `json:"players"`
	Map       []int     `json:"map"`
	Bombs     []Bomb    `json:"bombs"`
	Explosion Explosion `json:"explosion"`
	PowerUps  []PowerUp `json:"power_ups"`
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

// func (g *GameState) ClearExplosions() {
// 	g.Explosions = [][]Explosion{}
// }
