package game

// holds game state to send it to all players
type GameState struct {
	Players []Player `json:"players"`
	Map []int `json:"map"`
	Bombs   []Bomb   `json:"bombs"`
	Explosion Explosion `json:"explosion"`
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
// Loop through all players in game and check if somebody is in the explosion
func (g *GameState) CheckIfSomebodyDie(explosion *Explosion) {
	for i:=0 ; i<len(g.Players); i++{
		g.Players[i].CheckIfIDie(explosion)
	}
	// g.Explosions = [][]Explosion{}
}
