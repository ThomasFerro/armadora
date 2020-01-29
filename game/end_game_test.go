package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/gold"
)

/*
TODO:
- Territories are re-computed and the player with the more golds is the winner
- Manage the ties
*/

func TestGameEndsWhenEveryPlayerPassTheirTurn(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		event.PlayerJoined{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
		event.PlayerJoined{
			Nickname:  "Kileek",
			Character: character.Elf,
		},
		event.GoldStacksDistributed{
			GoldStacks: gold.GoldStacks,
		},
		event.PalisadesDistributed{
			Count: 35,
		},
		event.GameStarted{},
		event.TurnPassed{
			Player: 0,
		},
		event.NextPlayer{},
		event.NextPlayer{},
		event.TurnPassed{
			Player: 2,
		},
		event.NextPlayer{},
	}

	turnCommand := command.PassTurnPayload{
		Player: 1,
	}

	history = append(
		history,
		command.PassTurn(history, turnCommand)...,
	)

	eventFound := false

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.GameFinished); isOfRightEventType {
			eventFound = true
			break
		}
	}

	if !eventFound {
		t.Error("No 'GameFinished' event")
		return
	}

	currentGame := game.ReplayHistory(history)

	if currentGame.State() != game.Finished {
		t.Errorf("The game is not finished, %v instead", currentGame.State())
		return
	}
}
