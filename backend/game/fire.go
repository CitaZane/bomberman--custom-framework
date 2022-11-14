package game

type Fire struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	Type int `json:"type"`
}

// check if monster is in fire
// returns true -> monster burned and should die
// use error margin to make monster detection range smaller
func (fire *Fire) IsMonsterInside(x, y int) bool {
	var errorMargin = 5

	var monsterLeft = x + errorMargin
	var monsterRight = x - errorMargin
	var monsterUp = y + errorMargin
	var monsterDown = y - errorMargin

	if monsterLeft >= fire.X && monsterRight <= fire.X+64 && monsterUp >= fire.Y && monsterDown <= fire.Y+64 {
		return true
	}
	return false
}