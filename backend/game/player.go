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
}

type Bomb struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// initialization functions returns palyer with initial state and position in  11x11 field
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
	}
}

// methods for updatig monster position based on input from websocket
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

// Base movement functions
func (player *Player) MoveUp() {
	player.Y -= player.Speed
}
func (player *Player) MoveDown() {
	player.Y += player.Speed
}
func (player *Player) MoveRight() {
	player.X += player.Speed
}
func (player *Player) MoveLeft() {
	player.X -= player.Speed
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