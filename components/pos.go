package components

type Pos struct {
	X, Y int
}

func NewPosI(x, y int) Pos {
	return Pos{x, y}
}
