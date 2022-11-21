package game

import "math"

type Fire struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	Type int `json:"type"`
}

// check if monster is in fire
// returns true -> monster burned and should die
// use error margin to make monster detection range smaller
func (fire *Fire) IsMonsterInside(x, y float64) bool {
	var errorMargin = 5

	var monsterLeft = int(math.Round(x)) + errorMargin
	var monsterRight = int(math.Round(x)) + 64 - errorMargin
	var monsterUp = int(math.Round(y)) + errorMargin
	var monsterDown = int(math.Round(y)) + 64 - errorMargin

	if monsterLeft >= fire.X && monsterLeft <= fire.X+64 || monsterRight >= fire.X && monsterRight <= fire.X+64 {
		if monsterUp >= fire.Y && monsterUp <= fire.Y+64 || monsterDown >= fire.Y && monsterDown <= fire.Y+64 {
			return true
		}
	}
	return false
}
