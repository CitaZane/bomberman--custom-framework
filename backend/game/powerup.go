package game

import (
	"math/rand"
)

type PowerUpType string

const (
	None            PowerUpType = "None"
	INCREASE_BOMBS  PowerUpType = "increase_bombs"
	INCREASE_SPEED  PowerUpType = "increase_speed"
	INCREASE_FLAMES PowerUpType = "increase_flames"
)

type PowerUp struct {
	Type PowerUpType `json:"type"`
	X    int         `json:"x"`
	Y    int         `json:"y"`
	Tile int
}

// holds all the power ups that were generated during map creation
var GeneratedPowerUps = make([]*PowerUp, 0)

func newPowerUp(tile int) *PowerUp {
	var powerUpType PowerUpType
	amountOfPowerUps := len(GeneratedPowerUps)

	switch amountOfPowerUps % 3 {
	case 0:
		powerUpType = INCREASE_BOMBS
	case 1:
		powerUpType = INCREASE_FLAMES
	case 2:
		powerUpType = INCREASE_SPEED
	}

	x := tile % 11 * 64
	y := tile / 11 * 64

	return &PowerUp{
		Type: powerUpType,
		X:    x,
		Y:    y,
		Tile: tile,
	}

}

// generate a power up for a random breakable brick
// each power up will take up a space of blocks:
// 2 up, 2 right, 2 down, 2 left
// 1 each diagonal
// replace spots taken by a power up with number 3

func GeneratePowerUp(basemap []int, breakableBricks []int) {
	powerUpPlaced := false
	for !powerUpPlaced {
		randomPos := rand.Intn(len(breakableBricks))
		if basemap[breakableBricks[randomPos]] == 1 {
			// create a power up
			powerUp := newPowerUp(breakableBricks[randomPos])

			// save the power up
			GeneratedPowerUps = append(GeneratedPowerUps, powerUp)

			// change the basemap tile number to 3
			basemap[breakableBricks[randomPos]] = 3

			// left 1 tile
			if basemap[breakableBricks[randomPos]-1] == 1 {
				basemap[breakableBricks[randomPos]-1] = 3
			}

			// left 2 tiles
			if basemap[breakableBricks[randomPos]-2] == 1 {
				basemap[breakableBricks[randomPos]-2] = 3
			}

			// right 1 tile
			if basemap[breakableBricks[randomPos]+1] == 1 {
				basemap[breakableBricks[randomPos]+1] = 3
			}

			// right 2 tiles
			if basemap[breakableBricks[randomPos]+2] == 1 {
				basemap[breakableBricks[randomPos]+2] = 3
			}

			if breakableBricks[randomPos] > 12 {
				// up 1 tile
				if basemap[breakableBricks[randomPos]-11] == 1 {
					basemap[breakableBricks[randomPos]-11] = 3
				}

				// diagonal up right
				if basemap[breakableBricks[randomPos]-10] == 1 {
					basemap[breakableBricks[randomPos]-10] = 3
				}

				// diagonal up left
				if basemap[breakableBricks[randomPos]-12] == 1 {
					basemap[breakableBricks[randomPos]-12] = 3
				}

				if breakableBricks[randomPos] > 22 {
					// up 2 tiles
					if basemap[breakableBricks[randomPos]-22] == 1 {
						basemap[breakableBricks[randomPos]-22] = 3
					}
				}
			}

			if breakableBricks[randomPos]+12 < len(basemap) {
				// down 1 tile
				if basemap[breakableBricks[randomPos]+11] == 1 {
					basemap[breakableBricks[randomPos]+11] = 3
				}

				// diagonal down left 1 tile
				if basemap[breakableBricks[randomPos]+10] == 1 {
					basemap[breakableBricks[randomPos]+10] = 3
				}

				// diagonal right 1 tile
				if basemap[breakableBricks[randomPos]+12] == 1 {
					basemap[breakableBricks[randomPos]+12] = 3
				}

				if breakableBricks[randomPos]+22 < len(basemap) {
					// down 2 tile
					if basemap[breakableBricks[randomPos]+22] == 1 {
						basemap[breakableBricks[randomPos]+22] = 3
					}
				}
			}
			powerUpPlaced = true
		}
	}

}
