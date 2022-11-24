package game

type Bomb struct {
	X        int      `json:"x"`
	Y        int      `json:"y"`
	Tile int `json:"-"`
	Name     string   `json:"name"`
	Movement Movement `json:"movement"`
	Speed    int      `json:"-"` //for changing how fast is movement
}