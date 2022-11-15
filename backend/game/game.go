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
// return slice with monster that died
func (g *GameState) CheckIfSomebodyDied(explosion *Explosion) []int{
	var monstersLostLives = []int{}
	for i:=0 ; i<len(g.Players); i++{
		var lostLive = g.Players[i].CheckIfIDie(explosion)
		if lostLive {
			monstersLostLives = append(monstersLostLives, i)
		}
	}
	return monstersLostLives
}

// Loop through all active explosion in game and check if current player stepped in it
func (g *GameState) CheckIfPlayerDied(p *Player) bool {
	var lostLive = false
	for _, player := range g.Players{
		for _, explosion := range player.Explosions{
			lostLive = p.CheckIfIDie(&explosion)
			if lostLive{break}
		}
	}
	return lostLive
}