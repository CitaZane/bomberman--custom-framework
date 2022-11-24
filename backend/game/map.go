package game

import (
	"math/rand"
	"time"
)

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

type Tile int

const (
	Wall Tile = 2
	Brick Tile = 1
	Empty Tile = 0
	
	Blocked Tile= 3
	Replace Tile = 9
)

func CreateBaseMap() []Tile{
	rand.Seed(time.Now().UnixNano())
	basemap := []Tile{}
	// basemap := append([]int{}, mapBase...)
	GeneratedPowerUps = nil
	breakableBricks := []int{}
	for i, tile := range mapBase {
		if tile == 0{
			basemap = append(basemap, Empty)
		}else if tile == 2 {
			basemap = append(basemap, Wall)
		}else if tile == 9 {
			if rand.Intn(10) < 6 {
				basemap = append(basemap, Brick)
				breakableBricks = append(breakableBricks, i)
			} else {
				basemap = append(basemap, Empty)
			}
		}
	}

	// generate a powerup for 6 breakable bricks
	for i := 0; i < 6; i++ {
		GeneratePowerUp(basemap, breakableBricks)
	}

	// restore breakable bricks in map array
	for i, tile := range basemap {
		if tile == Blocked {
			basemap[i] = Brick
		}
	}
	return basemap
}

func DestroyBlocks(original []Tile, indexList []int) []Tile {
	for _, v := range indexList {
		original[v] = Empty
	}
	return original
}
