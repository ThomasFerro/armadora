package game

import (
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/warrior"
)

// Player A player
type Player interface {
	Nickname() string
	Character() character.Character
	Warriors() warrior.Warriors
	SetWarriors(warrior.Warriors) player
}

type player struct {
	nickname  string
	character character.Character
	warriors  warrior.Warriors
}

func (p player) Nickname() string {
	return p.nickname
}

func (p player) Character() character.Character {
	return p.character
}

func (p player) Warriors() warrior.Warriors {
	return p.warriors
}

func (p player) SetWarriors(warriors warrior.Warriors) player {
	p.warriors = warriors
	return p
}

// NewPlayer Create a new player
func NewPlayer(nickname string, character character.Character) Player {
	return player{
		nickname:  nickname,
		character: character,
	}
}
