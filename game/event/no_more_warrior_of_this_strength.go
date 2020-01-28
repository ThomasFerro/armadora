package event

import (
	"fmt"
)

// NoMoreWarriorOfThisStrength Distpach when no more warrior of the selected strength are available for the player
type NoMoreWarriorOfThisStrength struct {
	Strength int
}

// EventMessage Indicate that no more warrior of the selected strength are available
func (event NoMoreWarriorOfThisStrength) EventMessage() string {
	return fmt.Sprintf("No more warrior of the %v strength are available.", event.Strength)
}

func (event NoMoreWarriorOfThisStrength) String() string {
	return event.EventMessage()
}
