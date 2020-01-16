package game

import (
	"github.com/ThomasFerro/armadora/game/character"
)

// Player A player
type Player interface {
	Nickname() string
	Character() character.Character
}

type player struct {
	nickname  string
	character character.Character
}

func (p player) Nickname() string {
	return p.nickname
}

func (p player) Character() character.Character {
	return p.character
}

// NewPlayer Create a new player
func NewPlayer(nickname string, character character.Character) Player {
	return player{
		nickname,
		character,
	}
}
