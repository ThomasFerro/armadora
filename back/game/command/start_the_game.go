package command

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/warrior"
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

	events := []event.Event{}

	warriorsToDistribute := warrior.WarriorsToDistribute(len(newGame.Players()))

	events = append(events, event.WarriorsDistributed{
		warriorsToDistribute,
	})

	return append(events, event.GameStarted{})
}
