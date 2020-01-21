package command

import (
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/event"
)

// PutWarriorPayload Data about the warrior to put on the board
type PutWarriorPayload struct {
	Warrior  int
	Position board.Position
}

func PutWarrior(history []event.Event, payload PutWarriorPayload) []event.Event {
	return []event.Event{
		event.NextPlayer{},
	}
}
