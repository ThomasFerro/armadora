package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/gold"
	"github.com/ThomasFerro/armadora/game/warrior"
)

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

func TestScoresComputedWhenTheGameEnds(t *testing.T) {
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
		event.WarriorsDistributed{
			WarriorsDistributed: warrior.NewWarriors(10, 10, 10, 10, 10),
		},
		event.GameStarted{},
		event.WarriorPut{
			Player:   0,
			Strength: 5,
			Position: board.Position{
				X: 4,
				Y: 4,
			},
		},
		event.NextPlayer{},
		event.WarriorPut{
			Player:   1,
			Strength: 5,
			Position: board.Position{
				X: 4,
				Y: 0,
			},
		},
		event.NextPlayer{},
		event.WarriorPut{
			Player:   2,
			Strength: 5,
			Position: board.Position{
				X: 6,
				Y: 2,
			},
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 0,
			X1:     2,
			Y1:     0,
			X2:     3,
			Y2:     0,
		},
		event.PalisadePut{
			Player: 0,
			X1:     2,
			Y1:     1,
			X2:     3,
			Y2:     1,
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 1,
			X1:     4,
			Y1:     0,
			X2:     5,
			Y2:     0,
		},
		event.PalisadePut{
			Player: 1,
			X1:     4,
			Y1:     1,
			X2:     5,
			Y2:     1,
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 2,
			X1:     3,
			Y1:     1,
			X2:     3,
			Y2:     2,
		},
		event.PalisadePut{
			Player: 2,
			X1:     4,
			Y1:     1,
			X2:     4,
			Y2:     2,
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 0,
			X1:     4,
			Y1:     2,
			X2:     4,
			Y2:     3,
		},
		event.PalisadePut{
			Player: 0,
			X1:     5,
			Y1:     2,
			X2:     5,
			Y2:     3,
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 1,
			X1:     3,
			Y1:     3,
			X2:     4,
			Y2:     3,
		},
		event.PalisadePut{
			Player: 1,
			X1:     3,
			Y1:     4,
			X2:     4,
			Y2:     4,
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 2,
			X1:     5,
			Y1:     0,
			X2:     6,
			Y2:     0,
		},
		event.PalisadePut{
			Player: 2,
			X1:     5,
			Y1:     1,
			X2:     6,
			Y2:     1,
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 0,
			X1:     5,
			Y1:     2,
			X2:     6,
			Y2:     2,
		},
		event.PalisadePut{
			Player: 0,
			X1:     5,
			Y1:     3,
			X2:     6,
			Y2:     3,
		},
		event.NextPlayer{},

		event.PalisadePut{
			Player: 1,
			X1:     5,
			Y1:     4,
			X2:     6,
			Y2:     4,
		},
		event.NextPlayer{},
		event.NextPlayer{},

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

	passTurnEvents, err := command.PassTurn(history, turnCommand)

	if err != nil {
		t.Errorf("The player cannot pass his turn: %v", err)
		return
	}

	history = append(
		history,
		passTurnEvents...,
	)

	gameScores := game.ReplayHistory(history).Scores()

	if gameScores[1] == nil || gameScores[1].Player() != 2 || gameScores[1].TotalGold() != 12 {
		t.Errorf("The game's winner is invalid, expected the player 2 with 12 gold but got this %v", gameScores[1])
		return
	}

	if gameScores[2] == nil || gameScores[2].Player() != 0 || gameScores[2].TotalGold() != 6 {
		t.Errorf("The game's second best player is invalid, expected the player 0 with 6 gold but got this %v", gameScores[2])
		return
	}

	if gameScores[3] == nil || gameScores[3].Player() != 1 || gameScores[3].TotalGold() != 3 {
		t.Errorf("The game's worst player is invalid, expected the player 1 with 3 gold but got this %v", gameScores[3])
	}
}
