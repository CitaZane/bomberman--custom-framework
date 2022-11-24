package game

type State string

const (
	Lobby      State="lobby" //game still in lobby mode
	Play       State = "play"         //state while game is in action
	GameOver   State = "game-over"    //trigger game over screen
	WalkOfFame State = "walk-of-fame" //final walk
	Finish     State = "finish"       //game ended
)