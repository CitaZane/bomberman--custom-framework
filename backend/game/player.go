package game

import (
	"math"
)

type Player struct {
	X              float64     `json:"x"`
	Y              float64     `json:"y"`
	Name           string      `json:"name"`
	Movement       Movement    `json:"movement"`
	Invincible     bool        `json:"invincible"`
	Lives          int         `json:"lives"`
	Speed          float64     `json:"-"` //for changing how fast is movement
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

// initialization functions returns palyer with initial state and position in  11x11 field
func CreatePlayer(name string, index int, gameMap []int) Player {
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
		GameMap:        gameMap,
	}
}

// methods for updatig monster position based on input from websocket
func (player *Player) Move(input string, delta float64) {
	// update movement variable
	player.Movement = translateMovement(input)

	if player.Movement == Up {
		player.MoveUp(delta)
	} else if player.Movement == Down {
		player.MoveDown(delta)
	} else if player.Movement == Right {
		player.MoveRight(delta)
	} else if player.Movement == Left {
		player.MoveLeft(delta)
	} else if player.Movement == DropBomb {
		player.DropBomb()
	}

}

// automate movement to the middle of the screen
// for the walk of fame
func (player *Player) AutoMove(input string) bool {

	if player.X < 320 {
		player.Movement = Right
		if player.X + player.Speed > 320{
			player.X = 320
		}else{
			player.X += player.Speed
		}
	} else if player.X > 320 {
		player.Movement = Left
		if player.X + player.Speed < 320{
			player.X = 320
		}else{
			player.X -= player.Speed
		}
	} else if player.Y < 320 {
		player.Movement = Down
		if player.Y + player.Speed > 320{
			player.Y = 320
		}else{
			player.Y += player.Speed
		}
	} else if player.Y > 320 {
		player.Movement = Up
		if player.Y + player.Speed < 320{
			player.Y = 320
		}else{
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
			case INCREASE_SPEED:
				player.Speed += 0.5
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

//Movement functions with delta
func (player *Player) MoveUp(delta float64) {
	//Checks if the charecter lined up on a tile horizontally before allowing to move vertically
	if math.Mod(player.X, 64) != 0 {
		//if player is near a corner xFit function will help move left or right to line up perfectly
		xFit(player, delta)
		return
	}
	//player.Movement == "up" check will fail if method is called from xFit function
	if player.GameMap[player.calcPlayerPosition()-11] == 0 && player.Movement == "up" {
		player.Y -= player.Speed * delta
	} else {
		//checks the distance until the obsticle and sets that as the maximum amount allowed to move (fixes running over walls)
		if math.Mod(player.Y, 64) > player.Speed*delta {
			player.Y -= player.Speed * delta
		} else {
			player.Y -= math.Mod(player.Y, 64)
		}
	}
}

func (player *Player) MoveDown(delta float64) {
	if math.Mod(player.X, 64) != 0 {
		xFit(player, delta)
		return
	}
	if player.GameMap[player.calcPlayerPosition()+11] == 0 && player.Movement == "down" {
		player.Y += player.Speed * delta
	} else {
		if math.Mod(player.Y, 64) == 0 {
			return
		}
		if 64-math.Mod(player.Y, 64) > player.Speed*delta {
			player.Y += player.Speed * delta
		} else {
			player.Y += (64 - math.Mod(player.Y, 64))
		}
	}
}

func (player *Player) MoveLeft(delta float64) {
	if math.Mod(player.Y, 64) != 0 {
		yFit(player, delta)
		return
	}
	if player.GameMap[player.calcPlayerPosition()-1] == 0 && player.Movement == "left" {
		player.X -= player.Speed * delta
	} else {
		if math.Mod(player.X, 64) > player.Speed*delta {
			player.X -= player.Speed * delta
		} else {
			player.X -= math.Mod(player.X, 64)
		}
	}
}

func (player *Player) MoveRight(delta float64) {
	if math.Mod(player.Y, 64) != 0 {
		yFit(player, delta)
		return
	}
	if player.GameMap[player.calcPlayerPosition()+1] == 0 && player.Movement == "right" {
		player.X += player.Speed * delta
	} else {
		if math.Mod(player.X, 64) == 0 {
			return
		}
		if 64-math.Mod(player.X, 64) > player.Speed*delta {
			player.X += player.Speed * delta
		} else {
			player.X += 64 - math.Mod(player.X, 64)
		}
	}
}

//helps going around the corners
func xFit(player *Player, delta float64) {
	if math.Mod(player.X, 64) > 32 {
		if player.Movement == "down" {
			if player.GameMap[player.calcPlayerPosition()+11] == 0 {
				player.MoveRight(delta)
			}
		}
		if player.Movement == "up" {
			if player.GameMap[player.calcPlayerPosition()-11] == 0 {
				player.MoveRight(delta)
			}
		}
	} else {
		if player.Movement == "down" {
			if player.GameMap[player.calcPlayerPosition()+11] == 0 {
				player.MoveLeft(delta)
			}
		}
		if player.Movement == "up" {
			if player.GameMap[player.calcPlayerPosition()-11] == 0 {
				player.MoveLeft(delta)
			}
		}
	}
}

func yFit(player *Player, delta float64) {
	if math.Mod(player.Y, 64) > 32 {
		if player.Movement == "right" {
			if player.GameMap[player.calcPlayerPosition()+1] == 0 {
				player.MoveDown(delta)
			}
		}
		if player.Movement == "left" {
			if player.GameMap[player.calcPlayerPosition()-1] == 0 {
				player.MoveDown(delta)
			}
		}
	} else {
		if player.Movement == "right" {
			if player.GameMap[player.calcPlayerPosition()+1] == 0 {
				player.MoveUp(delta)
			}
		}
		if player.Movement == "left" {
			if player.GameMap[player.calcPlayerPosition()-1] == 0 {
				player.MoveUp(delta)
			}
		}
	}
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
	if player.Invincible {return false}
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
