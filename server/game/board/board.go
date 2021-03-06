package board

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/board/cell"
	"github.com/ThomasFerro/armadora/game/palisade"
)

// Position A position on the board
type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("Position [%v,%v]", p.X, p.Y)
}

// Board A board
type Board interface {
	GoldStacks() []int
	Cell(position Position) cell.Cell
	Palisades() []palisade.Palisade
	PalisadesLeft() int
	SetPalisadesLeft(palisadeLeft int) Board
	Width() int
	Height() int
	PutWarriorInCell(position Position, player int, strength int) Board
	PutPalisade(palisadeToPut palisade.Palisade) Board
}

type board struct {
	goldStacks    []int
	grid          [][]cell.Cell
	palisadesLeft int
	palisades     []palisade.Palisade
}

func (b board) GoldStacks() []int {
	return b.goldStacks
}

func (b board) Cell(position Position) cell.Cell {
	return b.grid[position.Y][position.X]
}

func (b board) PalisadesLeft() int {
	return b.palisadesLeft
}

func (b board) SetPalisadesLeft(palisadeLeft int) Board {
	b.palisadesLeft = palisadeLeft
	return b
}

func (b board) Palisades() []palisade.Palisade {
	return b.palisades
}

func (b board) Width() int {
	return len(b.grid[0])
}

func (b board) Height() int {
	return len(b.grid)
}

func (b board) PutWarriorInCell(position Position, player, strength int) Board {
	b.grid[position.Y][position.X] = cell.NewWarrior(player, strength)
	return b
}

func (b board) PutPalisade(palisadeToPut palisade.Palisade) Board {
	b.palisades = append(b.palisades, palisadeToPut)
	b.palisadesLeft--
	return b
}

var positionsWithGoldStacks = []Position{
	Position{
		X: 3,
		Y: 0,
	},
	Position{
		X: 1,
		Y: 1,
	},
	Position{
		X: 5,
		Y: 1,
	},
	Position{
		X: 7,
		Y: 1,
	},
	Position{
		X: 0,
		Y: 3,
	},
	Position{
		X: 4,
		Y: 3,
	},
	Position{
		X: 2,
		Y: 4,
	},
	Position{
		X: 6,
		Y: 4,
	},
}

func goldCell(x, y int) bool {
	for _, positionToCheck := range positionsWithGoldStacks {
		if positionToCheck.X == x && positionToCheck.Y == y {
			return true
		}
	}
	return false
}

func initGrid(goldStacks []int) [][]cell.Cell {
	var grid [][]cell.Cell
	currentGoldStackIndex := 0

	for y := 0; y < 5; y++ {
		grid = append(grid, []cell.Cell{})
		for x := 0; x < 8; x++ {
			if goldCell(x, y) {
				grid[y] = append(grid[y], cell.NewGold(goldStacks[currentGoldStackIndex]))
				currentGoldStackIndex++
			} else {
				grid[y] = append(grid[y], cell.NewLand())
			}
		}
	}

	return grid
}

// NewBoard Create a new board with the provided gold stacks distribution
func NewBoard(goldStacks []int) Board {
	grid := initGrid(goldStacks)
	return board{
		goldStacks,
		grid,
		0,
		[]palisade.Palisade{},
	}
}
