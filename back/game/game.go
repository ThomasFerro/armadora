package game

import (
	"github.com/ThomasFerro/armadora/game/event"
)

// Game Instance of an Aramdora game
type Game interface {
	State() State
	Players() []Player
	ApplyGameCreated(event event.GameCreated) Game
	ApplyPlayerJoined(event event.PlayerJoined) Game
}

type game struct {
	state   State
	players []Player
}

// State The game's current state
func (g game) State() State {
	return g.state
}

// Player The game's Players
func (g game) Players() []Player {
	return g.players
}

func (g game) ApplyGameCreated(event event.GameCreated) Game {
	g.state = WaitingForPlayers
	return g
}

func (g game) ApplyPlayerJoined(event event.PlayerJoined) Game {
	g.players = append(g.players, NewPlayer(event.Nickname, event.Character))
	return g
}

// ReplayHistory Replay the provided history to retrieve the game state
func ReplayHistory(history []event.Event) Game {
	var returnedGame Game
	returnedGame = game{}
	for _, nextEvent := range history {
		switch nextEvent.(type) {
		case event.GameCreated:
			gameCreatedEvent, _ := nextEvent.(event.GameCreated)
			returnedGame = returnedGame.ApplyGameCreated(gameCreatedEvent)
		case event.PlayerJoined:
			playerJoinedEvent, _ := nextEvent.(event.PlayerJoined)
			returnedGame = returnedGame.ApplyPlayerJoined(playerJoinedEvent)
		}
	}
	return returnedGame
}
