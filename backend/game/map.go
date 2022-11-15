package game

import (
	"math"
	"math/rand"
)

//function for creating starting map with random placed breakable blocks.
// 0 is floor tile
// 1 is breakable block
// 2 is wall
// func CreateBaseMap() []int {
// 	var m = make([]int, 0)
// 	i := 0
// 	for i < 121 {
// 		if i < 11 || i > 109 || i%11 == 0 || i == 10 || i == 21 || i == 32 || i == 43 || i == 54 || i == 65 || i == 76 || i == 87 || i == 98 || i == 109 || i == 24 || i == 26 || i == 28 || i == 30 || i == 46 || i == 48 || i == 50 || i == 52 || i == 68 || i == 70 || i == 72 || i == 74 || i == 90 || i == 92 || i == 94 || i == 96 {
// 			m = append(m, 2)
// 		} else {
// 			if rand.Intn(10) > 5 {
// 				m = append(m, 0)
// 			} else {
// 				m = append(m, 1)
// 			}
// 		}
// 		i++
// 	}
// 	return m
// }

/* ---------------------- maybe base map could help -> ---------------------- */
// predifined walls, empty corners for monsters and first breakable walls around monsters
// all places where int ==3 should be replaced

// var mapBase = []int{
// 	2,2,2,2,2,2,2,2,2,2,2,
// 	2,0,0,1,3,3,3,1,0,0,2,
// 	2,0,2,3,2,3,2,3,2,0,2,
// 	2,1,3,3,3,3,3,3,3,1,2,
// 	2,3,2,3,2,3,2,3,2,3,2,
// 	2,3,3,3,3,3,3,3,3,3,2,
// 	2,3,2,3,2,3,2,3,2,3,2,
// 	2,1,3,3,3,3,3,3,3,1,2,
// 	2,0,2,3,2,3,2,3,2,0,2,
// 	2,0,0,1,3,3,3,1,0,0,2,
// 	2,2,2,2,2,2,2,2,2,2,2,
// }

// map withouth 3s look like this->
// var mapBase = []int{
// 	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
// 	2, 0, 0, 1, 9, 9, 9, 1, 0, 0, 2,
// 	2, 0, 2, 9, 2, 9, 2, 9, 2, 0, 2,
// 	2, 1, 9, 9, 9, 9, 9, 9, 9, 1, 2,
// 	2, 9, 2, 9, 2, 9, 2, 9, 2, 9, 2,
// 	2, 9, 9, 9, 9, 9, 9, 9, 9, 9, 2,
// 	2, 9, 2, 9, 2, 9, 2, 9, 2, 9, 2,
// 	2, 1, 9, 9, 9, 9, 9, 9, 9, 1, 2,
// 	2, 0, 2, 9, 2, 9, 2, 9, 2, 0, 2,
// 	2, 0, 0, 1, 9, 9, 9, 1, 0, 0, 2,
// 	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
// }

var mapBase = []int{
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 0, 0, 9, 9, 9, 9, 9, 0, 0, 2,
	2, 0, 2, 9, 2, 9, 2, 9, 2, 0, 2,
	2, 9, 9, 9, 9, 9, 9, 9, 9, 9, 2,
	2, 9, 2, 9, 2, 9, 2, 9, 2, 9, 2,
	2, 9, 9, 9, 9, 9, 9, 9, 9, 9, 2,
	2, 9, 2, 9, 2, 9, 2, 9, 2, 9, 2,
	2, 9, 9, 9, 9, 9, 9, 9, 9, 9, 2,
	2, 0, 2, 9, 2, 9, 2, 9, 2, 0, 2,
	2, 0, 0, 9, 9, 9, 9, 9, 0, 0, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
}

func generatePowerUp(tile int, gameState *GameState) {

	powerUpX := tile % 11 * 64
	powerUpY := math.Floor(float64(tile)/11) * 64

	amountOfPowerUps := len(gameState.PowerUps)
	var powerUpType PowerUpType

	switch amountOfPowerUps % 3 {
	case 0:
		powerUpType = INCREASE_BOMBS
	case 1:
		powerUpType = INCREASE_FLAMES
	case 2:
		powerUpType = INCREASE_SPEED
	}
	gameState.PowerUps = append(gameState.PowerUps, PowerUp{Type: powerUpType, X: powerUpX, Y: int(powerUpY)})

}

func CreateBaseMap(game *GameState) []int {
	basemap := append([]int{}, mapBase...)
	game.PowerUps = nil
	breakableBricks := []int{}
	for i, tile := range basemap {
		if tile == 9 {
			if rand.Intn(10) < 6 {
				basemap[i] = 1
				breakableBricks = append(breakableBricks, i)

				// breakableBricks++

				// fmt.Println("breakableBricks: ", breakableBricks)
				// if breakableBricks%6 == 0 {
				// 	generatePowerUp(i, game)

				// }
			} else {
				basemap[i] = 0
			}
		}
	}

	// generatePowerUp()
	for i := 0; i < 6; i++ {
		powerUpPlaced := false
		for !powerUpPlaced {
			randomPos := rand.Intn(len(breakableBricks))
			if basemap[breakableBricks[randomPos]] == 1 {
				generatePowerUp(breakableBricks[randomPos], game)
				basemap[breakableBricks[randomPos]] = 3
				if basemap[breakableBricks[randomPos]-1] == 1 {
					basemap[breakableBricks[randomPos]-1] = 3
				}
				if basemap[breakableBricks[randomPos]-2] == 1 {
					basemap[breakableBricks[randomPos]-2] = 3
				}
				if basemap[breakableBricks[randomPos]+1] == 1 {
					basemap[breakableBricks[randomPos]+1] = 3
				}
				if basemap[breakableBricks[randomPos]+2] == 1 {
					basemap[breakableBricks[randomPos]+2] = 3
				}

				if breakableBricks[randomPos] > 12 {
					if basemap[breakableBricks[randomPos]-11] == 1 {
						basemap[breakableBricks[randomPos]-11] = 3
					}
					if basemap[breakableBricks[randomPos]-10] == 1 {
						basemap[breakableBricks[randomPos]-10] = 3
					}
					if basemap[breakableBricks[randomPos]-12] == 1 {
						basemap[breakableBricks[randomPos]-12] = 3
					}
					if breakableBricks[randomPos] > 23 {
						if basemap[breakableBricks[randomPos]-22] == 1 {
							basemap[breakableBricks[randomPos]-22] = 3
						}
						if basemap[breakableBricks[randomPos]-21] == 1 {
							basemap[breakableBricks[randomPos]-21] = 3
						}
						if basemap[breakableBricks[randomPos]-23] == 1 {
							basemap[breakableBricks[randomPos]-23] = 3
						}
					}
				}

				if breakableBricks[randomPos]+12 < len(basemap) {
					if basemap[breakableBricks[randomPos]+11] == 1 {
						basemap[breakableBricks[randomPos]+11] = 3
					}
					if basemap[breakableBricks[randomPos]+10] == 1 {
						basemap[breakableBricks[randomPos]+10] = 3
					}
					if basemap[breakableBricks[randomPos]+12] == 1 {
						basemap[breakableBricks[randomPos]+12] = 3
					}
					if breakableBricks[randomPos]+23 < len(basemap) {
						if basemap[breakableBricks[randomPos]+22] == 1 {
							basemap[breakableBricks[randomPos]+22] = 3
						}
						if basemap[breakableBricks[randomPos]+21] == 1 {
							basemap[breakableBricks[randomPos]+21] = 3
						}
						if basemap[breakableBricks[randomPos]+23] == 1 {
							basemap[breakableBricks[randomPos]+23] = 3
						}
					}
				}
				powerUpPlaced = true
			}
		}
		// randomPos := rand.Intn(len(breakableBricks))
		// generatePowerUp(breakableBricks[randomPos], game)
	}
	for i, tile := range basemap {
		if tile == 3 {
			basemap[i] = 1
		}
	}
	return basemap
}

func DestroyBlocks(original []int, indexList []int) []int {
	for _, v := range indexList {
		original[v] = 0
	}
	return original
}
