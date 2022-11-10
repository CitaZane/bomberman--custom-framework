package game

// holds game state to send it to all players
type GameState struct {
	Players  []Player  `json:"players"`
	Map      []int     `json:"map"`
	PowerUps []PowerUp `json:"power_ups"`
	// Explosions[][]Explosion `json:"explosions"`
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
