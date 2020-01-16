package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
)

/*
TODO:
- Create a game
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
	gameCreatedEvent := game.CreateGame()

	if len(gameCreatedEvent.EventMessage()) == 0 {
		t.Error("The game has not been created")
	}
}

func TestGameCreateInWaitingForPlayersState(t *testing.T) {
	gameCreatedEvent := game.CreateGame()

	newGame := game.ReplayHistory([]event.Event{
		gameCreatedEvent,
	})

	if newGame.State() != game.WaitingForPlayers {
		t.Error("The new game is not in 'WaitingForPlayers state")
	}
}
