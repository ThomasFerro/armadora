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

func TestFirstPlayerAgainWhenTheTurnIsOver(t *testing.T) {
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
		event.NextPlayer{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 1,
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

	var nextPlayerEventFound = 0

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.NextPlayer); isOfRightEventType {
			nextPlayerEventFound++
		}
	}

	if nextPlayerEventFound != 2 {
		t.Errorf("The \"NextPlayer\" event should have been dispatched twice but was only dispatched %v times", nextPlayerEventFound)
		return
	}

	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != 0 {
		t.Errorf("The current player is invalid, should be 0 instead of %v", currentGame.CurrentPlayer())
	}
}

// TODO: Same for warrior
func TestPalisadesCanOnlyBePutByTheCurrentPlayer(t *testing.T) {
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
		event.NextPlayer{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 0,
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
	lastEvent := history[len(history)-1]

	if _, isOfRightEventType := lastEvent.(event.NotThePlayerTurn); !isOfRightEventType {
		t.Error("The \"NotThePlayerTurn\" event should have been dispatched")
		return
	}

	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != 1 {
		t.Errorf("The current player is invalid, should be 1 instead of %v", currentGame.CurrentPlayer())
	}
}

func TestWarriorCanOnlyBePutByTheCurrentPlayer(t *testing.T) {
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
		event.NextPlayer{},
	}

	turnCommand := command.PutWarriorPayload{
		Player:  0,
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
	lastEvent := history[len(history)-1]

	if _, isOfRightEventType := lastEvent.(event.NotThePlayerTurn); !isOfRightEventType {
		t.Error("The \"NotThePlayerTurn\" event should have been dispatched")
		return
	}

	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != 1 {
		t.Errorf("The current player is invalid, should be 1 instead of %v", currentGame.CurrentPlayer())
	}
}
