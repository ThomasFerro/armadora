package game

import (
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/warrior"
)

// Player A player
type Player interface {
	Nickname() string
	Character() character.Character
	TurnPassed() bool
	PassTurn() Player
	Warriors() warrior.Warriors
	SetWarriors(warrior.Warriors) Player
}

type player struct {
	nickname   string
	character  character.Character
	turnPassed bool
	warriors   warrior.Warriors
}

func (p player) Nickname() string {
	return p.nickname
}

func (p player) Character() character.Character {
	return p.character
}

func (p player) TurnPassed() bool {
	return p.turnPassed
}

func (p player) PassTurn() Player {
	p.turnPassed = true
	return p
}

func (p player) Warriors() warrior.Warriors {
	return p.warriors
}

func (p player) SetWarriors(warriors warrior.Warriors) Player {
	p.warriors = warriors
	return p
}

// NewPlayer Create a new player
func NewPlayer(nickname string, character character.Character) Player {
	return player{
		nickname:   nickname,
		character:  character,
		turnPassed: false,
	}
}
