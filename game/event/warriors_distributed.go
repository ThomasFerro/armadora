package event

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/warrior"
)

// WarriorsDistributed The warriors are distributed
type WarriorsDistributed struct {
	WarriorsDistributed warrior.Warriors
}

// EventMessage Indicate that the warriors are distributed to the players
func (event WarriorsDistributed) EventMessage() string {
	return fmt.Sprintf("The following warriors are distributed: %v.", event.WarriorsDistributed)
}

func (event WarriorsDistributed) String() string {
	return event.EventMessage()
}
