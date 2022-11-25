package game

import (
	"math"
)

type Player struct {
	X              float64        `json:"x"`
	Y              float64        `json:"y"`
	Name           string         `json:"name"`
	Movement       Movement       `json:"movement"`
	Invincible     bool           `json:"invincible"`
	Lives          int            `json:"lives"`
	Speed          float64        `json:"-"` //for changing how fast is movement
	BombsLeft      int            `json:"bombsLeft"`
	Bombs          []Bomb         `json:"bombs"`
	ExplosionRange int            `json:"-"`
	Explosions     []Explosion    `json:"explosions"`
	ActivePowerUps ActivePowerUps `json:"active_powerups"`
}

type ActivePowerUps struct {
	Bombs  int `json:"bombs"`
	Flames int `json:"flames"`
	Speed  int `json:"speed"`
}

// initialization functions returns palyer with initial state and position in  11x11 field
func CreatePlayer(name string, index int) Player {
	var x float64
	var y float64
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
		Invincible:     false,
		X:              x,
		Y:              y,
		Lives:          3,
		ExplosionRange: 1,
		BombsLeft:      1,
		Bombs:          make([]Bomb, 0),
		Explosions:     []Explosion{},
		ActivePowerUps: ActivePowerUps{},
	}
}

// methods for updatig monster position based on input from websocket
func (player *Player) Move(input string, delta float64, gameState *GameState) {
	// update movement variable
	player.Movement = translateMovement(input)

	if player.Movement == Up {
		player.MoveUp(delta, gameState)
	} else if player.Movement == Down {
		player.MoveDown(delta, gameState)
	} else if player.Movement == Right {
		player.MoveRight(delta, gameState)
	} else if player.Movement == Left {
		player.MoveLeft(delta, gameState)
	} else if player.Movement == DropBomb {
		player.DropBomb()
	}

}

// automate movement to the middle of the screen
// for the walk of fame
func (player *Player) AutoMove(input string) bool {

	if player.X < 320 {
		player.Movement = Right
		if player.X+player.Speed > 320 {
			player.X = 320
		} else {
			player.X += player.Speed
		}
	} else if player.X > 320 {
		player.Movement = Left
		if player.X+player.Speed < 320 {
			player.X = 320
		} else {
			player.X -= player.Speed
		}
	} else if player.Y < 320 {
		player.Movement = Down
		if player.Y+player.Speed > 320 {
			player.Y = 320
		} else {
			player.Y += player.Speed
		}
	} else if player.Y > 320 {
		player.Movement = Up
		if player.Y+player.Speed < 320 {
			player.Y = 320
		} else {
			player.Y -= player.Speed
		}

	}

	if player.X == 320 && player.Y == 320 {
		player.Movement = DownStop
		return true
	}
	return false
}

func (player *Player) PickedUpPowerUp(powerUps *[]*PowerUp) bool {
	playerX, playerY := player.GetCurrentCoordinates()

	// check if there is a powerup on player x and y
	for i, powerUp := range *powerUps {

		if powerUp.X == playerX && powerUp.Y == playerY {
			s := *powerUps

			//add powerUp to the player
			switch powerUp.Type {
			case INCREASE_BOMBS:
				player.BombsLeft++
				player.ActivePowerUps.Bombs++
			case INCREASE_SPEED:
				player.Speed += 0.5
				player.ActivePowerUps.Speed++
			case INCREASE_FLAMES:
				player.ExplosionRange++
				player.ActivePowerUps.Flames++

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
	tile := player.calcPlayerPosition()
	player.Bombs = append(player.Bombs, Bomb{X: baseX, Y: baseY, Tile: tile})
	player.BombsLeft--
}
func (player *Player) BombExplosionComplete() {
	player.BombsLeft++
	player.Bombs = player.Bombs[1:]
}

// player create explosion
func (player *Player) MakeExplosion(gameMap []Tile) ([]int, Explosion) {
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

// Movement functions with delta
func (player *Player) MoveUp(delta float64, gameState *GameState) {
	if !player.isHorizontalyAligned() {
		xFit(player, delta, gameState)
		return
	}
	var nextTileIndex = player.calcPlayerPosition() - 11
	var nextTile = gameState.Map[nextTileIndex]
	var bombOnNextTile = gameState.IsThereBomb(nextTileIndex)

	if nextTile == Empty && player.Movement == Up && !bombOnNextTile {
		player.Y -= player.Speed * delta
		return
	}
	var distanceToTile = player.verticalDistanceToTile()
	if distanceToTile > player.Speed*delta {
		player.Y -= player.Speed * delta
	} else {
		player.Y -= distanceToTile
	}
}

func (player *Player) MoveDown(delta float64, gameState *GameState) {
	if !player.isHorizontalyAligned() {
		xFit(player, delta, gameState)
		return
	}
	var nextTileIndex = player.calcPlayerPosition() + 11
	var nextTile = gameState.Map[nextTileIndex]
	var bombOnNextTile = gameState.IsThereBomb(nextTileIndex)

	if nextTile == Empty && player.Movement == Down && !bombOnNextTile {
		player.Y += player.Speed * delta
		return
	}

	if player.isVerticalyAligned() {
		return
	}
	var distanceToTile = 64 - player.verticalDistanceToTile()
	if distanceToTile > player.Speed*delta {
		player.Y += player.Speed * delta
	} else {
		player.Y += distanceToTile
	}
}

func (player *Player) MoveLeft(delta float64, gameState *GameState) {
	if !player.isVerticalyAligned() {
		yFit(player, delta, gameState)
		return
	}
	var nextTileIndex = player.calcPlayerPosition() - 1
	var nextTile = gameState.Map[nextTileIndex]
	var bombOnNextTile = gameState.IsThereBomb(nextTileIndex)

	if nextTile == Empty && player.Movement == Left && !bombOnNextTile {
		player.X -= player.Speed * delta
		return
	}
	var distanceToTile = player.horizontalDistanceToTile()
	if distanceToTile > player.Speed*delta {
		player.X -= player.Speed * delta
	} else {
		player.X -= distanceToTile
	}
}

func (player *Player) MoveRight(delta float64, gameState *GameState) {
	if !player.isVerticalyAligned() {
		yFit(player, delta, gameState)
		return
	}
	var nextTileIndex = player.calcPlayerPosition() + 1
	var nextTile = gameState.Map[nextTileIndex]
	var bombOnNextTile = gameState.IsThereBomb(nextTileIndex)

	if nextTile == Empty && player.Movement == Right && !bombOnNextTile {
		player.X += player.Speed * delta
		return
	}

	if player.isHorizontalyAligned() {
		return
	}
	var distanceToTile = 64 - player.horizontalDistanceToTile()
	if distanceToTile > player.Speed*delta {
		player.X += player.Speed * delta
	} else {
		player.X += distanceToTile
	}
}

// if player is near a corner xFit function will help move left or right to line up perfectly
func xFit(player *Player, delta float64, gameState *GameState) {
	var playerPosition = player.calcPlayerPosition()
	if player.horizontalDistanceToTile() > 32 {
		if player.Movement == Down {
			if gameState.Map[playerPosition+11] == Empty {
				player.MoveRight(delta, gameState)
			}
		}
		if player.Movement == Up {
			if gameState.Map[playerPosition-11] == Empty {
				player.MoveRight(delta, gameState)
			}
		}
	} else {
		if player.Movement == Down {
			if gameState.Map[playerPosition+11] == Empty {
				player.MoveLeft(delta, gameState)
			}
		}
		if player.Movement == Up {
			if gameState.Map[playerPosition-11] == Empty {
				player.MoveLeft(delta, gameState)
			}
		}
	}
}

func yFit(player *Player, delta float64, gameState *GameState) {
	var playerPosition = player.calcPlayerPosition()
	if player.verticalDistanceToTile() > 32 {
		if player.Movement == Right {
			if gameState.Map[playerPosition+1] == Empty {
				player.MoveDown(delta, gameState)
			}
		}
		if player.Movement == Left {
			if gameState.Map[playerPosition-1] == Empty {
				player.MoveDown(delta, gameState)
			}
		}
	} else {
		if player.Movement == Right {
			if gameState.Map[playerPosition+1] == Empty {
				player.MoveUp(delta, gameState)
			}
		}
		if player.Movement == Left {
			if gameState.Map[playerPosition-1] == Empty {
				player.MoveUp(delta, gameState)
			}
		}
	}
}

func (player *Player) isHorizontalyAligned() bool {
	return player.horizontalDistanceToTile() == 0
}
func (player *Player) isVerticalyAligned() bool {
	return player.verticalDistanceToTile() == 0
}

func (player *Player) horizontalDistanceToTile() float64 {
	return math.Mod(player.X, 64)
}
func (player *Player) verticalDistanceToTile() float64 {
	return math.Mod(player.Y, 64)
}

// Calculates on which map cell player is standing. Cell is map index.
func (player *Player) calcPlayerPosition() int {
	xRemainder := math.Mod(player.X, 64)
	yRemainder := math.Mod(player.Y, 64)
	column := player.Y / 64
	row := player.X / 64
	if xRemainder > 32 {
		row++
	}
	if yRemainder > 32 {
		column++
	}
	index := int(math.Floor(column)*11 + math.Floor(row))
	return index
}

func getBase(x float64) int {
	var base = x
	var remainder = math.Mod(x, 64)
	if remainder > 32 { //base is next tile
		base += 64 - remainder
	} else { //base is previous tile
		base -= remainder
	}
	return int(math.Round(base))
}

func (player *Player) GetCurrentCoordinates() (int, int) {
	var baseX = getBase(player.X)
	var baseY = getBase(player.Y)
	return baseX, baseY
}

// bool value is true only if live lost, but monster is not dead yet
func (player *Player) CheckIfIDie(explosion *Explosion) bool {
	if player.Invincible {
		return false
	}
	// for each fire in explosion, check if monster is inside it
	for _, fire := range explosion.Fires {
		var monsterBurned = fire.IsMonsterInside(player.X, player.Y)
		if monsterBurned {
			player.LoseLife()
			return player.Invincible
		}
	}
	return player.Invincible
}
func (player *Player) LoseLife() bool {
	player.Lives = player.Lives - 1
	if player.Lives > 0 {
		player.Invincible = true
	} else {
		player.Movement = Died
	}
	return player.Invincible
}

// Check if player still alive
func (player *Player) IsAlive() bool {
	return player.Movement != Died
}
