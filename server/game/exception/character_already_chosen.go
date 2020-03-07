package exception

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game/character"
)

// CharacterAlreadyChosen The character has already been chosen by another player in the game
type CharacterAlreadyChosen struct {
	Character character.Character
}

// Error Indicate that the character is unavailable
func (event CharacterAlreadyChosen) Error() string {
	return fmt.Sprintf("The character %v has already been chosen by another player in the game.", event.Character)
}
