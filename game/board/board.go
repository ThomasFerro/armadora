package board

// Position A position on the board
type Position struct {
	X int
	Y int
}

// Board A board
type Board interface {
	GoldStacks() []int
}

type board struct {
	goldStacks []int
}

func (b board) GoldStacks() []int {
	return b.goldStacks
}

// NewBoard Create a new board with the provided gold stacks distribution
func NewBoard(goldStacks []int) Board {
	return board{
		goldStacks,
	}
}
