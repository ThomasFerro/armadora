package cell

import "fmt"

// Cell A board's cell
type Cell interface{}

// Warrior A cell with a warrior in it
type Warrior interface {
	Strength() int
	Player() int
}

type warrior struct {
	strength int
	player   int
}

func (w warrior) Strength() int {
	return w.strength
}

func (w warrior) Player() int {
	return w.player
}

func (w warrior) String() string {
	return fmt.Sprintf("Warrior of player %v with strength %v", w.player, w.strength)
}

func NewWarrior(player, strength int) Warrior {
	return warrior{
		strength,
		player,
	}
}

type Gold interface {
	Stack() int
}

type gold struct {
	stack int
}

func (g gold) Stack() int {
	return g.stack
}

func (g gold) String() string {
	return fmt.Sprintf("Gold with stack %v", g.stack)
}

func NewGold(stack int) Gold {
	return gold{
		stack,
	}
}

type Land interface{}

type land struct{}

func (l land) String() string {
	return "Land"
}

func NewLand() Land {
	return land{}
}
