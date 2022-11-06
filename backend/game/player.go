package game

type Player struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Name string `json:"name"`
	Movement Movement `json:"movement"` 
	Speed int `json:"-"` //for changing how fas is movement
}

// initialization functions returns palyer with initial state and position in  11x11 field
func CreatePlayer(name string, index int ) Player{
	var x int
	var y int
	var movement Movement
	switch index {
	case 0: x=64; y=64; movement=RightStop;
	case 1: x=576; y=64; movement=LeftStop;
	case 2: x=64; y=576; movement=RightStop;
	default: x=576; y=576; movement=LeftStop
	}
	return Player{
		Name: name,
		Speed: 1,
		Movement: movement,
		X:x,
		Y:y,
	}
}

// methods for updatig monster position based on input from websocket
func (player *Player) Move(input string){
	// update movement variable
	player.Movement = translateMovement(input)

	if player.Movement == Up{
		player.MoveUp()
	}else if player.Movement == Down{
		player.MoveDown()
	}else if player.Movement == Right{
		player.MoveRight()
	}else if player.Movement == Left{
		player.MoveLeft()
	}
}

// Base movement functions
func (player *Player) MoveUp(){
	player.Y -= player.Speed
}
func (player *Player) MoveDown(){
	player.Y += player.Speed
}
func (player *Player) MoveRight(){
	player.X += player.Speed
}
func (player *Player) MoveLeft(){
	player.X -= player.Speed
}