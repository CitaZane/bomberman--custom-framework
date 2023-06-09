package game

import (
	uuid "github.com/satori/go.uuid"
)

type Explosion struct {
	Fires []Fire `json:"fires"`
	Id    string `json:"id"`
}

type ExplosionManager struct {
	Range        int             //explosion range
	CurrentRange int             //holds current range in calculation
	FireAlive    map[string]bool //turn value to false if explosion blockd
	Base         map[string]int  //explosion base coordinates
	TypeMap      map[string]int  //map direction to type value
	EndValue     int             //end incrimentor
	BaseMap      []Tile
	Directions   []string
}

func setupManager(exlosionRange int, x, y int, baseMap []Tile) ExplosionManager {
	return ExplosionManager{
		Range:        exlosionRange,
		FireAlive:    map[string]bool{"UP": true, "DOWN": true, "LEFT": true, "RIGHT": true},
		Base:         map[string]int{"x": x, "y": y},
		TypeMap:      map[string]int{"UP": 1, "DOWN": 3, "LEFT": 4, "RIGHT": 2},
		EndValue:     4,
		BaseMap:      baseMap,
		CurrentRange: 0,
		Directions:   []string{"UP", "DOWN", "LEFT", "RIGHT"},
	}
}

// stopsserching for fire in certain direction
//
//	func (manager *ExplosionManager) turnOffFire(direction string) {
//		manager.FireAlive[direction] = false
//	}
func (manager *ExplosionManager) incrementRange() {
	manager.CurrentRange += 1
}

/* ------------------------- explosion type diagram ------------------------- */
// 		 __
// 		|5|
//  ____|1|___
// |8 4 0 2 6|
//     |3|
//     |7|

// Map 2=wall ; 1=Bush;  0=ground

// calculate new explosion based on bobm coordinates, base map and players explosion range
// return []Explosion tiles that makes 1 explosion and
// []indexes for bushes destroyed int the explosion
func NewExplosion(bomb *Bomb, m []Tile, player *Player) (Explosion, []int) {
	var id = "explosion-" + player.Name + "-" + uuid.NewV4().String()
	explosion := Explosion{Id: id} //hold end explosion
	destroyedBlocks := []int{}     //hold index of destroyed blocks

	// add base in place of bomb
	var base = Fire{X: bomb.X, Y: bomb.Y, Type: 0}
	explosion.Fires = append(explosion.Fires, base)

	// create explosion manager
	manager := setupManager(player.ExplosionRange, bomb.X, bomb.Y, m)

	//calculate and add explosions based on range
	for i := 1; i <= manager.Range; i++ {
		manager.incrementRange()
		for _, direction := range manager.Directions {
			// check explosion for all sides
			if !manager.FireAlive[direction] {
				continue
			}
			if fire, destroyed, ok := manager.configFire(direction); ok {
				explosion.Fires = append(explosion.Fires, fire)
				if destroyed != -1 {
					destroyedBlocks = append(destroyedBlocks, destroyed)
				}
			}

		}
	}
	return explosion, destroyedBlocks
}

// Configure explosin x and y coordinates based on directiona and range
func (manager *ExplosionManager) configFire(direction string) (Fire, int, bool) {
	var x = manager.Base["x"]
	var y = manager.Base["y"]
	switch direction {
	case "UP":
		y -= 64 * manager.CurrentRange
	case "DOWN":
		y += 64 * manager.CurrentRange
	case "LEFT":
		x -= 64 * manager.CurrentRange
	case "RIGHT":
		x += 64 * manager.CurrentRange
	}
	return manager.findFire(direction, x, y)
}

// return false if explosion not found
// if destroyed != -1 then bush burned
func (manager *ExplosionManager) findFire(direction string, x, y int) (Fire, int, bool) {
	fire := Fire{X: x, Y: y, Type: manager.TypeMap[direction]}
	blockDestroyed := -1
	// check if coordinates are not out of game board
	if x < 0 || x > 11*64 || y < 0 || y > 11*64 {
		manager.FireAlive[direction] = false
		return fire, blockDestroyed, false
	}
	// check if block is destroyable
	var mapIndex = findMapIndex(x, y)
	if manager.BaseMap[mapIndex] == Wall {
		manager.FireAlive[direction] = false
		return fire, blockDestroyed, false
	}
	var mapIndexNext = findmapNextIndex(x, y, direction)
	// make changes to type if end of the fire
	if manager.BaseMap[mapIndex] == Brick || manager.Range == manager.CurrentRange || manager.BaseMap[mapIndexNext] == 2 {
		manager.FireAlive[direction] = false
		fire.Type = manager.TypeMap[direction] + manager.EndValue
	}
	// bush destroyed -> set index
	if manager.BaseMap[mapIndex] == Brick {
		blockDestroyed = mapIndex
	}
	return fire, blockDestroyed, true
}

func findMapIndex(x, y int) int {
	return ((x / 64) + (y/64)*11)
}
func findmapNextIndex(x, y int, direction string) int {
	switch direction {
	case "UP":
		return ((x / 64) + ((y-64)/64)*11)
	case "DOWN":
		return ((x / 64) + ((y+64)/64)*11)
	case "LEFT":
		return (((x - 64) / 64) + (y/64)*11)
	case "RIGHT":
		return (((x + 64) / 64) + (y/64)*11)
	}
	return 0
}
