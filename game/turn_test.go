package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/palisade"
)

/*
TODO:
- Put a warrior on the board:
	- Put on an empty land
	- Unable to put on a cell already taken (gold or another warrior)
- Put a palisade:
	- Put one palisade
	- Put two palisades
	- Unable to put a palisade on a already taken border
	- Unable to put a palisade if it breaks grid validity (territory with less than 4 cells)
- An action can only be make by the current player
- After an action has been make, the current player change:
	- End of a turn, current player = 0
*/

func TestChangeTheCurrentPlayerWhenPuttingWarrior(t *testing.T) {
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
		event.GameStarted{},
	}

	turnCommand := command.PutWarriorPayload{
		Warrior: 1,
		Position: board.Position{
			X: 0,
			Y: 0,
		},
	}

	history = append(
		history,
		command.PutWarrior(history, turnCommand)...,
	)

	var nextPlayerEventFound = false

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.NextPlayer); isOfRightEventType {
			nextPlayerEventFound = true
			break
		}
	}

	if !nextPlayerEventFound {
		t.Error("No \"NextPlayer\" event has been dispatched")
		return
	}

	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != 1 {
		t.Errorf("The current player is invalid, should be 1 instead of %v", currentGame.CurrentPlayer())
	}
}

func TestChangeTheCurrentPlayerWhenPuttingPalisade(t *testing.T) {
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
		event.GameStarted{},
	}

	turnCommand := command.PutPalisadesPayload{
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X: 0,
				Y: 0,
			},
		},
	}

	history = append(
		history,
		command.PutPalisades(history, turnCommand)...,
	)

	var nextPlayerEventFound = false

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.NextPlayer); isOfRightEventType {
			nextPlayerEventFound = true
			break
		}
	}

	if !nextPlayerEventFound {
		t.Error("No \"NextPlayer\" event has been dispatched")
		return
	}

	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != 1 {
		t.Errorf("The current player is invalid, should be 1 instead of %v", currentGame.CurrentPlayer())
	}
}
