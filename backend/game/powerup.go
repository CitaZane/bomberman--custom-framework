package game

type PowerUpType int

// iota will assign each value a number. Starting from 0
// None is 0
// INCREASED_BOMS is 1
// etc..

const (
	None PowerUpType = iota
	INCREASED_BOMBS
	INCREASED_SPEED
	INCREASED_FLAMES
)

// you have the option to also print powerup type with string value
func (p PowerUpType) String() string {
	switch p {
	case INCREASED_BOMBS:
		return "increased_bombs"
	}

	return ""
}

type PowerUp struct {
	Type PowerUpType `json:"power_up_type"`
	X    int         `json:"x"`
	Y    int         `json:"y"`
	// Active bool        `json:"active"`
}
