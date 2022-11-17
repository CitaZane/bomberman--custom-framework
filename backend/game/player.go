package game

type Player struct {
	X              int         `json:"x"`
	Y              int         `json:"y"`
	Name           string      `json:"name"`
	Movement       Movement    `json:"movement"`
	Lives          int         `json:"lives"`
	Speed          int         `json:"-"` //for changing how fast is movement
	BombsLeft      int         `json:"bombsLeft"`
	Bombs          []Bomb      `json:"bombs"`
	ExplosionRange int         `json:"-"`
	Explosions     []Explosion `json:"explosions"`
	ActivePowerUp  PowerUpType `json:"active_powerup"`
	GameMap        []int
}

type Bomb struct {
	X        int      `json:"x"`
	Y        int      `json:"y"`
	Name     string   `json:"name"`
	Movement Movement `json:"movement"`
	Speed    int      `json:"-"` //for changing how fast is movement
}

// initialization functions returns player with initial state and position in  11x11 field
func CreatePlayer(name string, index int, gameMap []int) Player {
	var x int
	var y int
	var movement Movement
	switch index {
	case 0:
		x = 64
		y = 64
		movement = RightStop
	case 1:
		x = 576
		y = 576
		movement = LeftStop
	case 2:
		x = 64
		y = 576
		movement = RightStop
	default:
		x = 576
		y = 64
		movement = LeftStop
	}
	return Player{
		Name:           name,
		Speed:          1,
		Movement:       movement,
		X:              x,
		Y:              y,
		Lives:          3,
		ExplosionRange: 1,
		BombsLeft:      1,
		Bombs:          make([]Bomb, 0),
		Explosions:     []Explosion{},
		GameMap:        gameMap,
	}
}

// methods for updatig monster position based on input from websocket
func (player *Player) Move(input string, powerUps *[]*PowerUp) bool {
	// update movement variable
	player.Movement = translateMovement(input)

	if player.Movement == Up {
		player.MoveUp()
	} else if player.Movement == Down {
		player.MoveDown()
	} else if player.Movement == Right {
		player.MoveRight()
	} else if player.Movement == Left {
		player.MoveLeft()
	} else if player.Movement == DropBomb {
		player.DropBomb()
	}

	playerX, playerY := player.GetCurrentCoordinates()

	// check if there is a powerup on player x and y
	for i, powerUp := range *powerUps {

		if powerUp.X == playerX && powerUp.Y == playerY {
			s := *powerUps

			//add powerUp to the player
			switch powerUp.Type {
			case INCREASE_BOMBS:
				player.BombsLeft++
			case INCREASE_SPEED:
				player.Speed++
			case INCREASE_FLAMES:
				player.ExplosionRange++
			}

			// remove the powerup from powerups array
			s = append(s[:i], s[i+1:]...)
			*powerUps = s
			return true
		}
	}

	return false
}

// player drops the  bomb
func (player *Player) DropBomb() {
	baseX, baseY := player.GetCurrentCoordinates()
	player.Bombs = append(player.Bombs, Bomb{X: baseX, Y: baseY})
	player.BombsLeft--
}
func (player *Player) BombExplosionComplete() {
	player.BombsLeft++
	player.Bombs = player.Bombs[1:]
}

// player create explosion
func (player *Player) MakeExplosion(gameMap []int) ([]int, Explosion) {
	var explosion, destroyedBlocks = NewExplosion(&player.Bombs[0], gameMap, player)
	player.Explosions = append(player.Explosions, explosion)
	return destroyedBlocks, explosion
}
func (player *Player) ExplosionComplete() {
	if len(player.Explosions) == 0 {
		return
	}
	player.Explosions = player.Explosions[1:]
}

// Movement functions
func (player *Player) MoveUp() {
	if player.Y < 64 {
		return
	}
	if player.X%64 != 0 {
		xFit(player)
	}
	if player.X%64 == 0 {
		if player.GameMap[player.calcPlayerPosition()-11] != 0 && player.Y%64 == 0 {
			return
		} else {
			player.Y -= player.Speed * 2
		}
	}
}
func (player *Player) MoveDown() {
	if player.X%64 != 0 {
		xFit(player)
	}
	if player.X%64 == 0 {
		if player.GameMap[player.calcPlayerPosition()+11] != 0 && player.Y%64 == 0 {
			return
		} else {
			player.Y += player.Speed * 2
		}
	}
}

func (player *Player) MoveRight() {
	if player.X > 574 {
		return
	}
	if player.Y%64 != 0 {
		yFit(player)
	} else {
		if player.GameMap[player.calcPlayerPosition()+1] != 0 && player.X%64 == 0 && player.Y%64 == 0 {
			return
		} else {
			player.X += player.Speed * 2
		}
	}
}
func (player *Player) MoveLeft() {
	if player.X < 65 {
		return
	}
	if player.Y%64 != 0 {
		yFit(player)
	} else {
		if player.GameMap[player.calcPlayerPosition()-1] != 0 && player.X%64 == 0 {
			return
		} else {
			player.X -= player.Speed * 2
		}
	}
}

// Calculates on which map cell player is standing. Cell is map index.
func (player *Player) calcPlayerPosition() int {
	xRemainder := player.X % 64
	yRemainder := player.Y % 64

	row := player.Y / 64
	place := player.X / 64
	if xRemainder > 32 {
		place++
	}
	if yRemainder > 32 {
		row++
	}

	index := row*11 + place
	return index
}

// func (player *Player) showCoordinates() {
// 	fmt.Printf("x: %v y: %v\n", player.X, player.Y)
// }

// Fit player on x-axis. (Auto move near corners)
func xFit(player *Player) {
	if player.X%64 > 32 {
		if player.Movement == "down" {
			if player.GameMap[player.calcPlayerPosition()+11] == 0 {
				player.X = player.X + 2
			}
		}
		if player.Movement == "up" {
			if player.GameMap[player.calcPlayerPosition()-11] == 0 {
				player.X = player.X + 2
			}
		}
	} else {
		if player.Movement == "down" {
			if player.GameMap[player.calcPlayerPosition()+11] == 0 {
				player.X = player.X - 2
			}
		}
		if player.Movement == "up" {
			if player.GameMap[player.calcPlayerPosition()-11] == 0 {
				player.X = player.X - 2
			}
		}
	}
}

// fit player on y-axis
func yFit(player *Player) {
	if player.Y%64 > 32 {
		if player.Movement == "right" {
			if player.GameMap[player.calcPlayerPosition()+1] == 0 {
				player.Y = player.Y + 2
			}
		}
		if player.Movement == "left" {
			if player.GameMap[player.calcPlayerPosition()-1] == 0 {
				player.Y = player.Y + 2
			}
		}
	} else {
		if player.Movement == "right" {
			if player.GameMap[player.calcPlayerPosition()+1] == 0 {
				player.Y = player.Y - 2
			}
		}
		if player.Movement == "left" {
			if player.GameMap[player.calcPlayerPosition()-1] == 0 {
				player.Y = player.Y - 2
			}
		}
	}
}

func (player *Player) GetCurrentCoordinates() (int, int) {
	var baseX = getBase(player.X)
	var baseY = getBase(player.Y)
	return baseX, baseY
}
func getBase(x int) int {
	var base = x
	var remainder = x % 64
	if remainder > 32 { //base is next tile
		base += 64 - remainder
	} else { //base is previous tile
		base -= remainder
	}
	return base
}

// bool value is true only if live lost, but monster is not dead yet
func (player *Player) CheckIfIDie(explosion *Explosion) bool {
	// for each fire in explosion, check if monster is inside it
	var lostLive = false
	for _, fire := range explosion.Fires {
		var monsterBurned = fire.IsMonsterInside(player.X, player.Y)
		// if monster is inside the fire ->
		if monsterBurned {
			if state := player.LoseLife(); state == LostLive {
				lostLive = true
			}
			break
		}
	}
	return lostLive
}

func (player *Player) LoseLife() Movement {
	player.Lives = player.Lives - 1
	if player.Lives > 0 {
		player.Movement = LostLive
	} else {
		player.Movement = Died
	}
	return player.Movement
}

// Check if player still alive
func (player *Player) IsAlive() bool {
	if player.Movement == Died || player.Movement == LostLive {
		return false
	}
	return true
}
