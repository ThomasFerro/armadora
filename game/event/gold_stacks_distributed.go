package event

import (
	"fmt"
)

// GoldStacksDistributed The gold stacks are distributed
type GoldStacksDistributed struct {
	GoldStacks []int
}

// EventMessage Indicate that the gold stacks are distributed in the game's grid
func (event GoldStacksDistributed) EventMessage() string {
	return fmt.Sprintf("The gold stacks are distributed as follow: %v.", event.GoldStacks)
}
