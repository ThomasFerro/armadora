package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/board/cell"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/gold"
	"github.com/ThomasFerro/armadora/game/palisade"
	"github.com/ThomasFerro/armadora/game/warrior"
)

/*
TODO:
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
		event.WarriorsDistributed{
			WarriorsDistributed: warrior.NewWarriors(1, 0, 0, 0, 0),
		},
		event.GoldStacksDistributed{
			gold.GoldStacks,
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

func TestDecrementWarriorsCountWhenPuttingWarrior(t *testing.T) {
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
		event.WarriorsDistributed{
			WarriorsDistributed: warrior.NewWarriors(1, 0, 0, 0, 0),
		},
		event.GoldStacksDistributed{
			gold.GoldStacks,
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

	currentGame := game.ReplayHistory(history)
	firstPlayer := currentGame.Players()[0]

	if firstPlayer.Warriors().OnePoint() != 0 {
		t.Errorf("The warrior was not removed %v", firstPlayer)
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

func TestPutAWarriorOnAnEmptyLand(t *testing.T) {
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
		event.WarriorsDistributed{
			WarriorsDistributed: warrior.NewWarriors(0, 0, 1, 0, 0),
		},
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.GameStarted{},
	}

	turnCommand := command.PutWarriorPayload{
		Player:  0,
		Warrior: 3,
		Position: board.Position{
			X: 0,
			Y: 0,
		},
	}

	history = append(
		history,
		command.PutWarrior(history, turnCommand)...,
	)

	warriorPutEventFound := false

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.WarriorPut); !isOfRightEventType {
			warriorPutEventFound = true
			break
		}
	}

	if !warriorPutEventFound {
		t.Error("The \"WarriorPut\" event should have been dispatched")
		return
	}

	currentGame := game.ReplayHistory(history)

	cellToCheck := currentGame.Board().Cell(board.Position{
		X: 0,
		Y: 0,
	})

	warriorCell, isOfRightCellType := cellToCheck.(cell.Warrior)

	if !isOfRightCellType {
		t.Errorf("There is no warrior in the cell, %v instead", cellToCheck)
		return
	}

	if warriorCell.Player() != 0 || warriorCell.Strength() != 3 {
		t.Errorf("The warrior on the cell does not match the expected one, found %v", warriorCell)
	}
}

func TestUnableToPutOnACellAlreadyTaken(t *testing.T) {
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
		event.WarriorsDistributed{
			WarriorsDistributed: warrior.NewWarriors(1, 0, 1, 0, 0),
		},
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.GameStarted{},
		event.WarriorPut{
			Player:   0,
			Strength: 1,
			Position: board.Position{
				X: 6,
				Y: 2,
			},
		},
		event.NextPlayer{},
	}

	turnCommand := command.PutWarriorPayload{
		Player:  1,
		Warrior: 3,
		Position: board.Position{
			X: 6,
			Y: 2,
		},
	}

	history = append(
		history,
		command.PutWarrior(history, turnCommand)...,
	)

	eventFound := false
	var cellAlreadyTakenEvent event.CellAlreadyTaken

	for _, nextEvent := range history {
		if typedEvent, isOfRightEventType := nextEvent.(event.CellAlreadyTaken); isOfRightEventType {
			eventFound = true
			cellAlreadyTakenEvent = typedEvent
			break
		}
	}

	if !eventFound {
		t.Error("No 'CellAlreadyTaken' found")
		return
	}

	if cellAlreadyTakenEvent.Position.X != 6 ||
		cellAlreadyTakenEvent.Position.Y != 2 {
		t.Errorf("'CellAlreadyTaken' event found with invalid payload: %v", cellAlreadyTakenEvent)
		return
	}

	currentGame := game.ReplayHistory(history)

	cellToCheck := currentGame.Board().Cell(board.Position{
		X: 6,
		Y: 2,
	})

	warriorCell, isOfRightCellType := cellToCheck.(cell.Warrior)

	if !isOfRightCellType {
		t.Errorf("There is no warrior in the cell, %v instead", cellToCheck)
		return
	}

	if warriorCell.Player() != 0 || warriorCell.Strength() != 1 {
		t.Errorf("The warrior on the cell does not match the expected one, found %v", warriorCell)
	}
}

func TestCanOnlyPutWarriorThatThePlayerHaveLeft(t *testing.T) {
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
		event.WarriorsDistributed{
			WarriorsDistributed: warrior.NewWarriors(0, 0, 0, 0, 1),
		},
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.GameStarted{},
		event.WarriorPut{
			Player:   0,
			Strength: 5,
			Position: board.Position{
				X: 0,
				Y: 0,
			},
		},
		event.NextPlayer{},
		event.WarriorPut{
			Player:   1,
			Strength: 5,
			Position: board.Position{
				X: 1,
				Y: 0,
			},
		},
		event.NextPlayer{},
	}

	turnCommand := command.PutWarriorPayload{
		Player:  0,
		Warrior: 5,
		Position: board.Position{
			X: 2,
			Y: 1,
		},
	}

	history = append(
		history,
		command.PutWarrior(history, turnCommand)...,
	)

	eventFound := false
	var noMoreWarriorOfThisStrengthEvent event.NoMoreWarriorOfThisStrength

	for _, nextEvent := range history {
		if typedEvent, isOfRightEventType := nextEvent.(event.NoMoreWarriorOfThisStrength); isOfRightEventType {
			eventFound = true
			noMoreWarriorOfThisStrengthEvent = typedEvent
			break
		}
	}

	if !eventFound {
		t.Error("No 'NoMoreWarriorOfThisStrength' event found")
		return
	}

	if noMoreWarriorOfThisStrengthEvent.Strength != 5 {
		t.Errorf("The strength of the event does not match with the expected one, found %v", noMoreWarriorOfThisStrengthEvent.Strength)
		return
	}

	currentGame := game.ReplayHistory(history)

	cellToCheck := currentGame.Board().Cell(board.Position{
		X: 2,
		Y: 1,
	})

	_, isOfRightCellType := cellToCheck.(cell.Land)

	if !isOfRightCellType {
		t.Errorf("The cell should be empty, %v instead", cellToCheck)
	}
}
