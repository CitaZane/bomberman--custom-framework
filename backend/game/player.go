package game

import (
	"fmt"
)

type Player struct {
	X              int         `json:"x"`
	Y              int         `json:"y"`
	Name           string      `json:"name"`
	Movement       Movement    `json:"movement"`
	Speed          int         `json:"-"` //for changing how fas is movement
	BombsLeft      int         `json:"bombsLeft"`
	Bombs          []Bomb      `json:"bombs"`
	ExplosionRange int         `json:"-"`
	Explosions     []Explosion `json:"explosions"`
}

type Bomb struct {
	X        int      `json:"x"`
	Y        int      `json:"y"`
	Name     string   `json:"name"`
	Movement Movement `json:"movement"`
	Speed    int      `json:"-"` //for changing how fast is movement
}

// initialization functions returns player with initial state and position in  11x11 field
func CreatePlayer(name string, index int) Player {
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
		y = 64
		movement = LeftStop
	case 2:
		x = 64
		y = 576
		movement = RightStop
	default:
		x = 576
		y = 576
		movement = LeftStop
	}
	return Player{
		Name:           name,
		Speed:          1,
		Movement:       movement,
		X:              x,
		Y:              y,
		ExplosionRange: 1,
		BombsLeft:      1,
		Bombs:          make([]Bomb, 0),
		Explosions:     []Explosion{},
	}
}

// methods for updating monster position based on input from websocket
func (player *Player) Move(input string) {
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
func (player *Player) MakeExplosion(gameMap []int) []int {
	var explosion, destroyedBlocks = NewExplosion(&player.Bombs[0], gameMap, player)
	player.Explosions = append(player.Explosions, explosion)
	return destroyedBlocks
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
		if State.Map[player.calcPlayerPosition()-11] != 0 && player.Y%64 == 0 {
			// fmt.Println("'MoveUp' func blocks movement!				player.go(110)")
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
		if State.Map[player.calcPlayerPosition()+11] != 0 && player.Y%64 == 0 {
			// fmt.Println("'MoveDown' func blocks movement!				player.go(121)")
			return
		} else {
			player.Y += player.Speed * 2
		}
	}
}

func (player *Player) MoveRight() {
	if player.X > 574 {
		// fmt.Println("Right wall")
		return
	}
	if player.Y%64 != 0 {
		yFit(player)
	} else {
		if State.Map[player.calcPlayerPosition()+1] != 0 && player.X%64 == 0 && player.Y%64 == 0 {
			// fmt.Println("'MoveRight' func blocks movement!				player.go(133)")
			return
		} else {
			player.X += player.Speed * 2
		}
	}
}
func (player *Player) MoveLeft() {
	// player.showCoordinates()
	if player.X < 65 {
		return
	}
	if player.Y%64 != 0 {
		yFit(player)
	} else {
		if State.Map[player.calcPlayerPosition()-1] != 0 && player.X%64 == 0 {
			// fmt.Println("'MoveLeft' func blocks movement!				player.go(147)")
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

func (player *Player) showCoordinates() {
	fmt.Printf("x: %v y: %v\n", player.X, player.Y)
}

// Fit player on x-axis. (Auto move near corners)
func xFit(player *Player) {
	// fmt.Println(player.calcPlayerPosition(), "				player.go(174)")

	if player.X%64 > 32 {
		if player.Movement == "down" {
			if State.Map[player.calcPlayerPosition()+11] == 0 {
				player.X = player.X + 2
			}
		}
		if player.Movement == "up" {
			if State.Map[player.calcPlayerPosition()-11] == 0 {
				player.X = player.X + 2
			}
		}
	} else {
		if player.Movement == "down" {
			if State.Map[player.calcPlayerPosition()+11] == 0 {
				player.X = player.X - 2
			}
		}
		if player.Movement == "up" {
			if State.Map[player.calcPlayerPosition()-11] == 0 {
				player.X = player.X - 2
			}
		}
	}
}

// fit player on y-axis
func yFit(player *Player) {
	// fmt.Println(player.calcPlayerPosition(), "				player.go(204)")
	if player.Y%64 > 32 {
		if player.Movement == "right" {
			if State.Map[player.calcPlayerPosition()+1] == 0 {
				player.Y = player.Y + 2
			}
		}
		if player.Movement == "left" {
			if State.Map[player.calcPlayerPosition()-1] == 0 {
				player.Y = player.Y + 2
			}
		}
	} else {
		if player.Movement == "right" {
			if State.Map[player.calcPlayerPosition()+1] == 0 {
				player.Y = player.Y - 2
			}
		}
		if player.Movement == "left" {
			if State.Map[player.calcPlayerPosition()-1] == 0 {
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
