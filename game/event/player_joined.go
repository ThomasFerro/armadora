package event

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/character"
)

// PlayerJoined A player has joined the game
type PlayerJoined struct {
	Nickname  string
	Character character.Character
}

// EventMessage Indicate that a player has joined the game
func (event PlayerJoined) EventMessage() string {
	return fmt.Sprintf("%v has joined the game as %v.", event.Nickname, event.Character)
}

func (event PlayerJoined) String() string {
	return event.EventMessage()
}
