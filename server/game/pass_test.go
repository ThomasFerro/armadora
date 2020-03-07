package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/exception"
	"github.com/ThomasFerro/armadora/game/gold"
)

func TestAPlayerCanPassHisTurn(t *testing.T) {
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
			GoldStacks: gold.GoldStacks,
		},
		event.PalisadesDistributed{
			Count: 35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PassTurnPayload{
		Player: 0,
	}

	passTurnEvents, err := command.PassTurn(history, turnCommand)

	if err != nil {
		t.Errorf("The player cannot pass his turn: %v", err)
		return
	}

	history = append(
		history,
		passTurnEvents...,
	)

	eventFound := false
	var turnPassedEvent event.TurnPassed

	for _, nextEvent := range history {
		if typedEvent, isOfRightEventType := nextEvent.(event.TurnPassed); isOfRightEventType {
			eventFound = true
			turnPassedEvent = typedEvent
			break
		}
	}

	if !eventFound {
		t.Error("No 'TurnPassed' event")
		return
	}

	if turnPassedEvent.Player != 0 {
		t.Errorf("Event payload invalid, expected the player 0 but found %v", turnPassedEvent)
		return
	}

	currentGame := game.ReplayHistory(history)

	if !currentGame.Players()[0].TurnPassed() {
		t.Errorf("The player's turn is not passed. Player: %v", currentGame.Players()[0])
	}
}

func TestChangeTheCurrentPlayerWhenPassingTurn(t *testing.T) {
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
			GoldStacks: gold.GoldStacks,
		},
		event.PalisadesDistributed{
			Count: 35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PassTurnPayload{
		Player: 0,
	}

	passTurnEvents, err := command.PassTurn(history, turnCommand)

	if err != nil {
		t.Errorf("The player cannot pass his turn: %v", err)
		return
	}

	history = append(
		history,
		passTurnEvents...,
	)

	eventFound := false

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.NextPlayer); isOfRightEventType {
			eventFound = true
			break
		}
	}

	if !eventFound {
		t.Error("No 'NextPlayer' event")
		return
	}

	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != 1 {
		t.Errorf("The current player is invalid. Expected 1 but got %v", currentGame.CurrentPlayer())
	}
}

func TestCannotPassWhenItIsNotThePlayerTurn(t *testing.T) {
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
			GoldStacks: gold.GoldStacks,
		},
		event.PalisadesDistributed{
			Count: 35,
		},
		event.GameStarted{},
	}

	turnCommand := command.PassTurnPayload{
		Player: 1,
	}

	_, err := command.PassTurn(history, turnCommand)

	if _, isOfRightExceptionType := err.(exception.NotThePlayerTurn); !isOfRightExceptionType {
		t.Errorf("The player should not be able to pass when it is not his turn: %v", err)
	}
}

func TestIgnorePlayersWhoPassedTheyTurn(t *testing.T) {
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

	currentGame := game.ReplayHistory(
		append(history, event.NextPlayer{}),
	)

	if currentGame.CurrentPlayer() != 1 {
		t.Errorf("The current player is invalid. Expected 1 but got %v", currentGame.CurrentPlayer())
	}
}
