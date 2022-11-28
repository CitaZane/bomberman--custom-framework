package game

type GameState struct {
	Players  []Player   `json:"players"`
	Map      []Tile     `json:"map"`
	Bombs    []Bomb     `json:"bombs"`
	PowerUps []*PowerUp `json:"power_ups"` // holds power ups, which are shown on screen
	State    State      `json:"state"`
}

func NewGame() *GameState {
	return &GameState{
		Players:  make([]Player, 0),
		Bombs:    make([]Bomb, 0),
		Map:      make([]Tile, 0),
		PowerUps: make([]*PowerUp, 0),
		State:    Lobby,
	}
}
func (g *GameState) StartGame() {
	g.Map = CreateBaseMap()
	g.State = Play
}
func (g *GameState) FinishGame() {
	g.Map = []Tile{}
	// g.Players = []Player{}
	g.Bombs = []Bomb{}
	g.PowerUps = []*PowerUp{}
	g.State = Lobby
}

func (g *GameState) FindPlayer(name string) int {
	for index, player := range g.Players {
		if player.Name == name {
			return index
		}
	}

	return -1
}

func (g *GameState) IsPlayer(name string) bool {
	playerIndex := g.FindPlayer(name)
	return playerIndex != -1
}

// check how many players are still alive.
// in case of 1 player left -> game over
func (g *GameState) CheckGameOverState() {
	playersAlive := 0
	for _, player := range g.Players {
		if player.Movement != Died {
			playersAlive += 1
		}
	}
	if playersAlive == 1 && len(g.Players) != 1 {
		g.State = GameOver
	}
}

func (g *GameState) ClearGameIfLastPlayerLeft() {
	playersAlive := 0
	for _, player := range g.Players {
		if player.Movement != Died {
			playersAlive += 1
		}
	}
	if playersAlive == 0 {
		g.FinishGame()
	}
}

// Loop through all players in game and check if somebody is in the explosion
// return slice with monster that died
func (g *GameState) CheckIfSomebodyDied(explosion *Explosion) []int {
	var monstersLostLives = []int{}
	if g.State != Play {
		return monstersLostLives
	}
	for i := 0; i < len(g.Players); i++ {
		var lostLive = g.Players[i].CheckIfIDie(explosion)
		if lostLive {
			monstersLostLives = append(monstersLostLives, i)
		}
	}
	g.CheckGameOverState()
	return monstersLostLives
}

// Loop through all active explosion in game and check if current player stepped in it
func (g *GameState) CheckIfPlayerDied(p *Player) bool {
	var lostLive = false
	if g.State != Play {
		return lostLive
	}
	for _, player := range g.Players {
		for _, explosion := range player.Explosions {
			lostLive = p.CheckIfIDie(&explosion)
			if lostLive {
				return lostLive
			}
		}
	}
	g.CheckGameOverState()
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

func (g *GameState) LetMonstersReborn(monstersLostLives []int) {
	for _, i := range monstersLostLives { //reset the movement
		g.Players[i].Invincible = false
	}
}

func (g *GameState) FindWinner() string {
	for _, player := range g.Players {
		if player.Movement != Died {
			return player.Name
		}
	}
	return ""
}

func (g *GameState) IsThereBomb(tileIndex int) bool {
	for _, player := range g.Players {
		for _, bomb := range player.Bombs {
			if bomb.Tile == tileIndex {
				return true
			}
		}
	}
	return false
}
