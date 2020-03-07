package exception

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/board"
)

// CellAlreadyTaken The cell is already taken
type CellAlreadyTaken struct {
	Position board.Position
}

// Error Indicate that the cell is already taken
func (exception CellAlreadyTaken) Error() string {
	return fmt.Sprintf("The cell %v is already taken.", exception.Position)
}
