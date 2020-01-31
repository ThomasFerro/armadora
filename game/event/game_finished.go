package event

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/score"
)

// GameFinished Dispatched when the game is finished
type GameFinished struct {
	Scores score.Scores
}

// EventMessage Indicate that the game is finished
func (event GameFinished) EventMessage() string {
	return fmt.Sprintf("The game is finished with scores %v.", event.Scores)
}

func (event GameFinished) String() string {
	return event.EventMessage()
}
