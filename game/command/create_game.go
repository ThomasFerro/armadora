package command

import (
	"github.com/ThomasFerro/armadora/game/event"
)

// CreateGame Create a new game
func CreateGame() event.GameCreated {
	return event.GameCreated{}
}
