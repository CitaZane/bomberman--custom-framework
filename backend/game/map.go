package game

import (
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
	2, 0, 0, 1, 9, 9, 9, 9, 0, 0, 2,
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

func CreateBaseMap(game *GameState) []int {
	basemap := append([]int{}, mapBase...)
	GeneratedPowerUps = nil
	breakableBricks := []int{}
	for i, tile := range basemap {
		if tile == 9 {
			if rand.Intn(10) < 6 {
				basemap[i] = 1
				breakableBricks = append(breakableBricks, i)
			} else {
				basemap[i] = 0
			}
		}
	}

	// generate a powerup for 6 breakable bricks
	for i := 0; i < 6; i++ {
		GeneratePowerUp(basemap, breakableBricks)
	}

	// restore breakable bricks in map array
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
