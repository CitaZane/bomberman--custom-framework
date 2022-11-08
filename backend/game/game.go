package game

type GameState struct {
	Players []Player `json:"players"`
	Map     []int    `json:"map"`
	// created bool
}

// holds game state to send it to all players
var State = GameState{}

func (g *GameState) FindPlayer(name string) int {
	for index, player := range g.Players {
		if player.Name == name {
			return index
		}
	}

	return -1
}

// func (g *GameState) GetMap() []int {
// 	return g.Map
// }
