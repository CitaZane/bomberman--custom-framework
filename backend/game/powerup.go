package game

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
	// Active bool        `json:"active"`
}
