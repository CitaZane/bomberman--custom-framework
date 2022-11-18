package game

type GameState struct {
	Players   []Player   `json:"players"`
	Map       []int      `json:"map"`
	Bombs     []Bomb     `json:"bombs"`
	Explosion Explosion  `json:"explosion"`
	PowerUps  []*PowerUp `json:"power_ups"` // holds power ups, which are shown on screen
	// created bool
}

func NewGame() *GameState {
	return &GameState{
		Players:   make([]Player, 0),
		Bombs:     make([]Bomb, 0),
		Map:       make([]int, 0),
		PowerUps:  make([]*PowerUp, 0),
		Explosion: Explosion{},
	}
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
func (g *GameState) CheckIfSomebodyDied(explosion *Explosion) []int {
	var monstersLostLives = []int{}
	for i := 0; i < len(g.Players); i++ {
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
	for _, player := range g.Players {
		for _, explosion := range player.Explosions {
			lostLive = p.CheckIfIDie(&explosion)
			if lostLive {
				return lostLive
			}
		}
	}
	return lostLive
}
// check if destroyed block index match with powerup block index
func (g *GameState) RevealPowerUps(destroyedBlocks []int) {
	for _, blockIndex := range destroyedBlocks {
		for _, powerUp := range GeneratedPowerUps {
			if blockIndex == powerUp.Tile {
				g.PowerUps = append(g.PowerUps, powerUp)
			}
		}
	}
}

func(g *GameState)LetMonstersReborn(monstersLostLives []int){
	for _, i := range monstersLostLives { //reset the movement
		g.Players[i].Movement = RightStop
	}
}