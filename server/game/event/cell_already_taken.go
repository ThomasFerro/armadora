package event

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/board"
)

// CellAlreadyTaken The cell is already taken
type CellAlreadyTaken struct {
	Position board.Position
}

// EventMessage Indicate that the cell is already taken
func (event CellAlreadyTaken) EventMessage() string {
	return fmt.Sprintf("The cell %v is already taken.", event.Position)
}

func (event CellAlreadyTaken) String() string {
	return event.EventMessage()
}
