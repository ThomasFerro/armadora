package command

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/exception"
	"github.com/ThomasFerro/armadora/game/gold"
	"github.com/ThomasFerro/armadora/game/warrior"
)

// StartTheGame Start the game
func StartTheGame(history []event.Event) ([]event.Event, error) {
	newGame := game.ReplayHistory(history)

	if len(newGame.Players()) < 2 {
		return nil, exception.NotEnoughPlayers{
			NumberOfPlayers: len(newGame.Players()),
		}
	}

	events := []event.Event{}

	warriorsToDistribute := warrior.WarriorsToDistribute(len(newGame.Players()))

	events = append(events, event.WarriorsDistributed{
		warriorsToDistribute,
	})

	goldStacksToDistribute := gold.GoldToDistribute()

	events = append(events, event.GoldStacksDistributed{
		goldStacksToDistribute,
	})

	events = append(events, event.PalisadesDistributed{
		35,
	})

	return append(events, event.GameStarted{}), nil
}
