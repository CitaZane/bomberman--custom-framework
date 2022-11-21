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
