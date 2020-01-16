package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
)

/*
TODO:
- Join a game with a nickname and a race
- Start the game
- One player cannot start a game alone
- A fifth player cannot join a game
- Players cannot join a game once it started
- A new player cannot select a race already chosen
- Distribute the wariors based on the number of players
- Distribute the gold
*/

func TestCreateAGame(t *testing.T) {
	gameCreatedEvent := command.CreateGame()

	if len(gameCreatedEvent.EventMessage()) == 0 {
		t.Error("The game has not been created")
	}
}

func TestGameCreateInWaitingForPlayersState(t *testing.T) {
	gameCreatedEvent := command.CreateGame()

	newGame := game.ReplayHistory([]event.Event{
		gameCreatedEvent,
	})

	if newGame.State() != game.WaitingForPlayers {
		t.Error("The new game is not in 'WaitingForPlayers state")
	}
}

func TestJoinAGame(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
	}

	joinGameCommandPayload := command.JoinGamePayload{
		Nickname:  "README.md",
		Character: character.Goblin,
	}

	history = append(history, command.JoinGame(history, joinGameCommandPayload)...)

	newGame := game.ReplayHistory(history)

	if len(newGame.Players()) != 1 {
		t.Error("The player did not join the game")
		return
	}

	newPlayer := newGame.Players()[0]

	if newPlayer.Nickname() != "README.md" || newPlayer.Character() != character.Goblin {
		t.Error("The player's information are not set correctly")
	}
}
