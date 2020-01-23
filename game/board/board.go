package board

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/board/cell"
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
	Width() int
	Height() int
	PutWarriorInCell(position Position, player int, strength int) Board
}

type board struct {
	goldStacks []int
	grid       [][]cell.Cell
}

func (b board) GoldStacks() []int {
	return b.goldStacks
}

func (b board) Cell(position Position) cell.Cell {
	return b.grid[position.X][position.Y]
}

func (b board) Width() int {
	return len(b.grid)
}

func (b board) Height() int {
	return len(b.grid[0])
}

func (b board) PutWarriorInCell(position Position, player, strength int) Board {
	b.grid[position.X][position.Y] = cell.NewWarrior(player, strength)
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

	for x := 0; x < 8; x++ {
		grid = append(grid, []cell.Cell{})
		for y := 0; y < 5; y++ {
			if goldCell(x, y) {
				grid[x] = append(grid[x], cell.NewGold(goldStacks[currentGoldStackIndex]))
				currentGoldStackIndex++
			} else {
				grid[x] = append(grid[x], cell.NewLand())
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
	}
}
