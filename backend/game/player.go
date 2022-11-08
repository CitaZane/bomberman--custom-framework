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
	// fmt.Println(player.calcPlayerPosition())
	player.calcPlayerPosition()
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
	player.showCoordinates()
	if player.Y == 64 {
		return
	}
	// if (player.X-2)%64 == 0 || (player.X-1)%64 == 0 || player.X%64 == 0 || (player.X+1)%64 == 0 || (player.X+2)%64 == 0 {
	if xFit(player) {
		if State.Map[player.calcPlayerPosition()] == 1 || State.Map[player.calcPlayerPosition()] == 2 {
			return
		}
		player.Y -= player.Speed
	}
}
func (player *Player) MoveDown() {
	player.showCoordinates()
	// if (player.X-2)%64 == 0 || (player.X-1)%64 == 0 || player.X%64 == 0 || (player.X+1)%64 == 0 || (player.X+2)%64 == 0 {
	if xFit(player) {

		if State.Map[player.calcPlayerPosition()+11] == 1 || State.Map[player.calcPlayerPosition()+11] == 2 {
			return
		}
		player.Y += player.Speed
	}
}

func xFit(player *Player) bool {
	i := 5
	for i > -5 {
		if !((player.X-i)%64 == 0) {
			i--
			continue
		} else {
			return true
		}
	}
	return false
}
func yFit(player *Player) bool {
	i := 5
	for i > -5 {
		if !((player.Y-i)%64 == 0) {
			i--
			continue
		} else {
			return true
		}
	}
	return false
}

func (player *Player) MoveRight() {
	player.showCoordinates()
	// if (player.Y-2)%64 == 0 || (player.Y-1)%64 == 0 || player.Y%64 == 0 || (player.Y+1)%64 == 0 || (player.Y+2)%64 == 0 {
	if yFit(player) {
		if State.Map[player.calcPlayerPosition()+1] == 1 || State.Map[player.calcPlayerPosition()+1] == 2 {
			return
		}
		player.X += player.Speed
	}

	// if player.X < 576 {
	// 	player.X += player.Speed
	// }
}
func (player *Player) MoveLeft() {
	player.showCoordinates()
	if player.X == 64 {
		return
	}
	// fmt.Println("player pos", player.calcPlayerPosition())
	// fmt.Println("index,val", player.calcPlayerPosition()-1, State.Map[player.calcPlayerPosition()-1])
	// fmt.Println("MAP", State.Map)

	// if (player.Y-2)%64 == 0 || (player.Y-1)%64 == 0 || player.Y%64 == 0 || (player.Y+1)%64 == 0 || (player.Y+2)%64 == 0 {
	if yFit(player) {
		if State.Map[player.calcPlayerPosition()] == 1 || State.Map[player.calcPlayerPosition()] == 2 {
			return
		}
		player.X -= player.Speed
	}

	// if player.X > 64 {
	// 	player.X -= player.Speed
	// }
}

// Calculates on which map cell player is standing. Cell is map index.
func (player *Player) calcPlayerPosition() int {
	// fmt.Println("player", player)
	row := player.Y / 64
	// fmt.Println("row", row)
	place := player.X / 64
	// fmt.Println("place", place)
	// fmt.Println("row, place: ", row, place)
	index := row*11 + place
	// fmt.Println("index: ", index)
	return index
}

// Check is there obstacles on the way.
// func (player *Player) isBlockOnTheWay(gs *GameState) bool {

// 	// fmt.Println("state map", gs.Map)
// 	// fmt.Println("player x,y: ", player.X, player.Y)
// 	// fmt.Println("map index, value: ", player.X/64+12, gs.Map[player.X/64+12])
// 	if gs.Map[player.X/64+12] == 1 || gs.Map[player.X/64+12] == 2 {
// 		return true
// 	}
// 	return false
// }

func (player *Player) showCoordinates() {
	fmt.Printf("x: %v y: %v\n", player.X, player.Y)
}
