package game

import (
	"github.com/ThomasFerro/armadora/game/event"
)

// Game Instance of an Aramdora game
type Game interface {
	State() State
	Apply(event event.GameCreated) Game
}

type game struct {
	state State
}

// State The game's current state
func (g game) State() State {
	return g.state
}

func (g game) Apply(event event.GameCreated) Game {
	g.state = WaitingForPlayers
	return g
}

// CreateGame Create a new game
func CreateGame() event.GameCreated {
	return event.GameCreated{}
}

// ReplayHistory Replay the provided history to retrieve the game state
func ReplayHistory(history []event.Event) Game {
	var returnedGame Game
	returnedGame = game{}
	for _, nextEvent := range history {
		switch nextEvent.(type) {
		case event.GameCreated:
			gameCreatedEvent, _ := nextEvent.(event.GameCreated)
			returnedGame = returnedGame.Apply(gameCreatedEvent)
		}
	}
	return returnedGame
}
