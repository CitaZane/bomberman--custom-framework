package websocket

import (
	g "bomberman-dom/game"
	"time"
)

type Message struct {
	Type      string       `json:"type"`
	Creator   string       `json:"creator"`
	Body      string       `json:"body"`
	GameState *g.GameState `json:"gameState"`
	Delta     float64      `json:"delta"`
}

func (m Message) ExplosionComplete(pool *Pool) {
	time.Sleep(900 * time.Millisecond)
	m.Type = "EXPLOSION_COMPLETED"
	pool.Broadcast <- m
}

func (m Message) BombExploded(pool *Pool) {
	time.Sleep(3000 * time.Millisecond)
	m.Type = "BOMB_EXPLODED"
	pool.Broadcast <- m
}

func (m Message) MonstersReborn(pool *Pool, gameState *g.GameState, monstersLostLives []int) {
	if len(monstersLostLives) == 0 {
		return
	}
	time.Sleep(3000 * time.Millisecond)
	m.Type = "PLAYER_REBORN"
	gameState.LetMonstersReborn(monstersLostLives)
	pool.Broadcast <- m
}

func (m Message) UpdateMap(pool *Pool, gameState *g.GameState, destroyedBlocks []int) {
	if len(destroyedBlocks) == 0 {
		return
	}
	time.Sleep(450 * time.Millisecond)
	m.Type = "MAP_UPDATE"
	gameState.Map = g.DestroyBlocks(gameState.Map, destroyedBlocks)
	gameState.RevealPowerUps(destroyedBlocks)

	pool.Broadcast <- m
}

// Game over screen updates game map in spiral
// turning all tiles to walls
// in case of open power up on screen -> remove it
func (m Message)ActivateGameOverScreen(pool *Pool, gameState *g.GameState){
	spiralLoop := formSpiral(gameState.Map) //calcuklate spiral, ints corresponds to index of original map
	for t, index :=range spiralLoop{
		miliseconds := t *45 //make delay different for each tile
		i:=index 
		go func(){
			time.Sleep( time.Duration(miliseconds)* time.Millisecond)
			gameState.Map[i] = 2 //turn into wall
			// remove power up if detected on tile
			for idx,powerUp := range gameState.PowerUps{
				if powerUp.Tile == i{
					gameState.PowerUps = append(gameState.PowerUps[:idx], gameState.PowerUps[idx +1:]...)
				}
			}
			m.Type = "MAP_UPDATE"
			pool.Broadcast <- m
		}()
	}
}


// Auto guide winner to the middle
// spawn movement after sleep time
func (m Message)AutoGuideWinner(pool *Pool, winner string){
	time.Sleep(4* time.Millisecond)

	m.Creator = winner
	m.Type = "PLAYER_AUTO_MOVE"
	pool.Broadcast <- m

}

/* -------------------------------------------------------------------------- */
/*                                   helper                                   */
/* -------------------------------------------------------------------------- */
// calculate indexes in spirale
// example -> from [0,1,2,3,4,5,6,7,8] in grid 3x3
// result  ->      [0,1,2,5,8,7,6,3,4]
func formSpiral(base []g.Tile) []int{
	var(
		rows = 11;
		col = 11;
		left = 0;
		top = 0;
		bottom = rows-1;
		right = col -1;
	) 
	var result = []int{}
	for {
		// calculate top row
		if (left>right) {break};
		for i:= left; i<=right; i++ {
			result = append(result, i + top * col)
		}
		top +=1;
		// calculate right column
		if(top >bottom){break};
		for i:= top; i<=bottom; i++{
			result = append(result, right + i * col)
		}
		right -=1

		// calculate bottom row
		if(left> right){break} ;
		for i:= right; i>=left; i--{
			result = append(result, i + bottom * col)
		}
		bottom -= 1;
		// claculate left column
		if (top>bottom){ break};
		for i:= bottom; i>=top; i--{
			result = append(result, left + i * col)
		}
		left +=1
	}	
	return result
}