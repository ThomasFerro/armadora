package command

import (
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/event"
)

// JoinGamePayload Containing every information about the player who wants to join
type JoinGamePayload struct {
	Nickname  string
	Character character.Character
}

// JoinGame Join a game
func JoinGame(history []event.Event, joinGamePayload JoinGamePayload) []event.Event {
	// TODO: Check if there is already 4 players
	// TODO: Check if the character is already chosen
	return []event.Event{
		event.PlayerJoined{
			Nickname:  joinGamePayload.Nickname,
			Character: joinGamePayload.Character,
		},
	}
}
