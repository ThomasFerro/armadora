package event

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/character"
)

// CharacterAlreadyChosen The character has already been chosen by another player in the game
type CharacterAlreadyChosen struct {
	Character character.Character
}

// EventMessage Indicate that the character is unavailable
func (event CharacterAlreadyChosen) EventMessage() string {
	return fmt.Sprintf("The character %v has already been chosen by another player in the game.", event.Character)
}
