package command

import (
	"github.com/ThomasFerro/armadora/game/event"
)

// StartTheGame Start the game
func StartTheGame(history []event.Event) []event.Event {
	return []event.Event{
		event.GameStarted{},
	}
}
