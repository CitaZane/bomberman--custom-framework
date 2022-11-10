package game

import "fmt"

type Player struct {
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
		Name:     name,
		Speed:    1,
		Movement: movement,
		X:        x,
		Y:        y,
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
	}
}

// Movement functions
func (player *Player) MoveUp() {
	// player.showCoordinates()
	xFit(player)
	if player.X%64 == 0 {
		if State.Map[player.calcPlayerPosition()] == 1 || State.Map[player.calcPlayerPosition()] == 2 {
			return
		} else {
			player.Y -= player.Speed * 2
		}
	}
}
func (player *Player) MoveDown() {
	// player.showCoordinates()
	xFit(player)
	if player.X%64 == 0 {
		if State.Map[player.calcPlayerPosition()+11] == 1 || State.Map[player.calcPlayerPosition()+11] == 2 {
			return
		} else {
			player.Y += player.Speed * 2
		}
	}
}

func (player *Player) MoveRight() {
	// player.showCoordinates()
	yFit(player)
	if player.Y%64 == 0 {
		if State.Map[player.calcPlayerPosition()+1] == 1 || State.Map[player.calcPlayerPosition()+1] == 2 {
			return
		} else {
			player.X += player.Speed * 2
		}
	}
}
func (player *Player) MoveLeft() {
	yFit(player)
	if player.Y%64 == 0 {
		if State.Map[player.calcPlayerPosition()] == 1 || State.Map[player.calcPlayerPosition()] == 2 {
			return
		} else {
			player.X -= player.Speed * 2
		}
	}
}

// Calculates on which map cell player is standing. Cell is map index.
func (player *Player) calcPlayerPosition() int {
	row := player.Y / 64
	place := player.X / 64
	index := row*11 + place
	return index
}

func (player *Player) showCoordinates() {
	fmt.Printf("x: %v y: %v\n", player.X, player.Y)
}

func xFit(player *Player) {
	if player.X%64 > 32 {
		player.X++
	} else {
		player.X--
	}
}

func yFit(player *Player) {
	if player.Y%64 > 32 {
		player.Y++
	} else {
		player.Y--
	}
}
