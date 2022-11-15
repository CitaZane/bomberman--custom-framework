package game

// movement enum
// Movement enup used to provide possible monster actions
type Movement string

const (
	Up    Movement = "up"
	Down  Movement = "down"
	Left  Movement = "left"
	Right Movement = "right"

	UpStop    Movement = "up-stop"
	DownStop  Movement = "down-stop"
	LeftStop  Movement = "left-stop"
	RightStop Movement = "right-stop"

	DropBomb  Movement = "drop-bomb"
	LostLive Movement  = "lost-live"
	Died Movement      = "died"
)

// Translate monster movement from string to one of registerd movements
func translateMovement(movement string) Movement {
	switch movement {
	case "DOWN":
		return Down
	case "UP":
		return Up
	case "LEFT":
		return Left
	case "RIGHT":
		return Right
	case "DOWN-STOP":
		return DownStop
	case "UP-STOP":
		return UpStop
	case "LEFT-STOP":
		return LeftStop
	case "DROP-BOMB":
		return DropBomb
	default:
		return RightStop
	}
}
