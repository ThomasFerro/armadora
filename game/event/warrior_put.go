package event

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/board"
)

// WarriorPut A warrior has been put on the board
type WarriorPut struct {
	Player   int
	Strength int
	Position board.Position
}

// EventMessage Indicate that a warrior has been put on the board
func (event WarriorPut) EventMessage() string {
	return fmt.Sprintf("%v has put a %v warrior on the board at the position %v.", event.Player, event.Strength, event.Position)
}

func (event WarriorPut) String() string {
	return event.EventMessage()
}
