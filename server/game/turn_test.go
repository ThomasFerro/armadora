package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/board/cell"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/exception"
	"github.com/ThomasFerro/armadora/game/gold"
	"github.com/ThomasFerro/armadora/game/palisade"
	"github.com/ThomasFerro/armadora/game/warrior"
)

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

	putWarriorEvents, err := command.PutWarrior(history, turnCommand)

	if err != nil {
		t.Errorf("The player cannot put a warrior: %v", err)
		return
	}

	history = append(
		history,
		putWarriorEvents...,
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PutPalisadesPayload{
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 0,
				Y1: 0,
				X2: 1,
				Y2: 0,
			},
		},
	}

	putPalisadesEvents, err := command.PutPalisades(history, turnCommand)

	if err != nil {
		t.Errorf("The player could not put palisades: %v", err)
		return
	}

	history = append(
		history,
		putPalisadesEvents...,
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
		event.NextPlayer{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 1,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 0,
				Y1: 0,
				X2: 1,
				Y2: 0,
			},
		},
	}

	putPalisadesEvents, err := command.PutPalisades(history, turnCommand)

	if err != nil {
		t.Errorf("The player could not put palisades: %v", err)
		return
	}

	history = append(
		history,
		putPalisadesEvents...,
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
		event.NextPlayer{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 0,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 0,
				Y1: 0,
				X2: 0,
				Y2: 0,
			},
		},
	}

	_, err := command.PutPalisades(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.NotThePlayerTurn); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put palisades when it is not his turn: %v", err)
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

	putWarriorEvents, err := command.PutWarrior(history, turnCommand)

	if err != nil {
		t.Errorf("The player cannot put a warrior: %v", err)
		return
	}

	history = append(
		history,
		putWarriorEvents...,
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

	_, err := command.PutWarrior(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.NotThePlayerTurn); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put warrior when it is not his turn: %v", err)
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

	putWarriorEvents, err := command.PutWarrior(history, turnCommand)

	if err != nil {
		t.Errorf("The player cannot put a warrior: %v", err)
		return
	}

	history = append(
		history,
		putWarriorEvents...,
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

	_, err := command.PutWarrior(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.CellAlreadyTaken); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put warrior on an already taken cell: %v", err)
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

	_, err := command.PutWarrior(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.NoMoreWarriorOfThisStrength); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put warrior that he does not have %v", err)
	}
}

func TestPutOnePalisade(t *testing.T) {
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 0,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 0,
				Y1: 0,
				X2: 1,
				Y2: 0,
			},
		},
	}

	putPalisadesEvents, err := command.PutPalisades(history, turnCommand)

	if err != nil {
		t.Errorf("The player could not put palisades: %v", err)
		return
	}

	history = append(
		history,
		putPalisadesEvents...,
	)

	palisadePutEvent := []event.PalisadePut{}

	for _, nextEvent := range history {
		if typedEvent, isOfRightEventType := nextEvent.(event.PalisadePut); isOfRightEventType {
			palisadePutEvent = append(palisadePutEvent, typedEvent)
		}
	}

	if len(palisadePutEvent) != 1 {
		t.Errorf("Should have put one palisade, found %v palisade(s) instead", len(palisadePutEvent))
		return
	}

	palisadePut := palisadePutEvent[0]

	if palisadePut.X1 != 0 || palisadePut.Y1 != 0 || palisadePut.X2 != 1 || palisadePut.Y2 != 0 {
		t.Errorf("The palisade does not match with the expected one, got this instead %v", palisadePut)
	}
}

func TestPutTwoPalisades(t *testing.T) {
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 0,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 0,
				Y1: 0,
				X2: 1,
				Y2: 0,
			},
			palisade.Palisade{
				X1: 0,
				Y1: 1,
				X2: 0,
				Y2: 2,
			},
		},
	}

	putPalisadesEvents, err := command.PutPalisades(history, turnCommand)

	if err != nil {
		t.Errorf("The player could not put palisades: %v", err)
		return
	}

	history = append(
		history,
		putPalisadesEvents...,
	)

	palisadePutEvent := []event.PalisadePut{}

	for _, nextEvent := range history {
		if typedEvent, isOfRightEventType := nextEvent.(event.PalisadePut); isOfRightEventType {
			palisadePutEvent = append(palisadePutEvent, typedEvent)
		}
	}

	if len(palisadePutEvent) != 2 {
		t.Errorf("Should have put two palisades, found %v palisade(s) instead", len(palisadePutEvent))
		return
	}

	firstPalisade := palisadePutEvent[0]

	if firstPalisade.X1 != 0 || firstPalisade.Y1 != 0 || firstPalisade.X2 != 1 || firstPalisade.Y2 != 0 {
		t.Errorf("The first palisade does not match with the expected one, got this instead %v", firstPalisade)
		return
	}

	secondPalisade := palisadePutEvent[1]

	if secondPalisade.X1 != 0 || secondPalisade.Y1 != 1 || secondPalisade.X2 != 0 || secondPalisade.Y2 != 2 {
		t.Errorf("The second palisade does not match with the expected one, got this instead %v", secondPalisade)
	}
}

func TestPalisadeInvalidIfNotBetweenToAdjacentCells(t *testing.T) {
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 0,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 1,
				Y1: 1,
				X2: 0,
				Y2: 0,
			},
		},
	}

	_, err := command.PutPalisades(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.InvalidPalisadePosition); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put palisades on an invalid position: %v", err)
	}
}

func TestUnableToPutAPalisadeOnAAlreadyTakenBorder(t *testing.T) {
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
		event.PalisadePut{
			Player: 0,
			X1:     6,
			Y1:     2,
			X2:     7,
			Y2:     2,
		},
		event.NextPlayer{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 1,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 6,
				Y1: 2,
				X2: 7,
				Y2: 2,
			},
		},
	}

	_, err := command.PutPalisades(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.BorderAlreadyTaken); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put palisades on already taken border: %v", err)
	}
}

func TestOnlyOnePalisadePutForACommandWithTwiceTheSameBorder(t *testing.T) {
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.PalisadesDistributed{
			35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 0,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 4,
				Y1: 1,
				X2: 5,
				Y2: 1,
			},
			palisade.Palisade{
				X1: 4,
				Y1: 1,
				X2: 5,
				Y2: 1,
			},
		},
	}

	putPalisadesEvents, err := command.PutPalisades(history, turnCommand)

	if err != nil {
		t.Errorf("The player could not put palisades: %v", err)
		return
	}

	history = append(
		history,
		putPalisadesEvents...,
	)

	eventFound := 0

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.PalisadePut); isOfRightEventType {
			eventFound++
		}
	}

	if eventFound != 1 {
		t.Errorf("Expected to find one 'PalisadePut' event but found %v", eventFound)
	}
}

func TestUnableToPutAPalisadeIfThereIsNoMoreLeft(t *testing.T) {
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.GameStarted{},
		event.PalisadesDistributed{
			1,
		},
		event.PalisadePut{
			Player: 0,
			X1:     2,
			Y1:     1,
			X2:     2,
			Y2:     0,
		},
		event.NextPlayer{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 1,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 6,
				Y1: 1,
				X2: 7,
				Y2: 1,
			},
		},
	}

	_, err := command.PutPalisades(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.NoMorePalisadeLeft); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put palisades when there is no more left: %v", err)
	}
}

func TestUnableToPutAPalisadeIfItBreaksGridValidity(t *testing.T) {
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
		event.GoldStacksDistributed{
			gold.GoldStacks,
		},
		event.GameStarted{},
		event.PalisadesDistributed{
			35,
		},
		event.PalisadePut{
			Player: 0,
			X1:     0,
			Y1:     0,
			X2:     1,
			Y2:     0,
		},
		event.PalisadePut{
			Player: 0,
			X1:     0,
			Y1:     1,
			X2:     1,
			Y2:     1,
		},
		event.NextPlayer{},
		event.PalisadePut{
			Player: 1,
			X1:     0,
			Y1:     2,
			X2:     1,
			Y2:     2,
		},
		event.NextPlayer{},
	}

	turnCommand := command.PutPalisadesPayload{
		Player: 0,
		Palisades: []palisade.Palisade{
			palisade.Palisade{
				X1: 0,
				Y1: 2,
				X2: 0,
				Y2: 3,
			},
		},
	}

	_, err := command.PutPalisades(history, turnCommand)
	if _, isOfRightExceptionType := err.(exception.InvalidPalisadePosition); !isOfRightExceptionType {
		t.Errorf("The player should not be able to put palisades on invalid position: %v", err)
	}
}
