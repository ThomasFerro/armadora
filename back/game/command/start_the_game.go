package command

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
)

// StartTheGame Start the game
func StartTheGame(history []event.Event) []event.Event {
	newGame := game.ReplayHistory(history)

	if len(newGame.Players()) < 2 {
		return []event.Event{
			event.NotEnoughPlayers{
				NumberOfPlayers: len(newGame.Players()),
			},
		}
	}

	return []event.Event{
		event.GameStarted{},
	}
}
