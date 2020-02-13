package warrior

// Warriors A player's warriors
type Warriors interface {
	OnePoint() int
	TwoPoints() int
	ThreePoints() int
	FourPoints() int
	FivePoints() int
}

type warriors struct {
	onePoint    int
	twoPoints   int
	threePoints int
	fourPoints  int
	fivePoints  int
}

func (w warriors) OnePoint() int {
	return w.onePoint
}

func (w warriors) TwoPoints() int {
	return w.twoPoints
}

func (w warriors) ThreePoints() int {
	return w.threePoints
}

func (w warriors) FourPoints() int {
	return w.fourPoints
}

func (w warriors) FivePoints() int {
	return w.fivePoints
}

func RemoveUsedWarrior(previousWarriors Warriors, strengthOfTheUsedWarrior int) Warriors {
	newWarriors := warriors{
		previousWarriors.OnePoint(),
		previousWarriors.TwoPoints(),
		previousWarriors.ThreePoints(),
		previousWarriors.FourPoints(),
		previousWarriors.FivePoints(),
	}
	switch strengthOfTheUsedWarrior {
	case 1:
		newWarriors.onePoint--
	case 2:
		newWarriors.twoPoints--
	case 3:
		newWarriors.threePoints--
	case 4:
		newWarriors.fourPoints--
	case 5:
		newWarriors.fivePoints--
	}
	return newWarriors
}

// NewWarriors Create a new warriors pool
func NewWarriors(onePoint, twoPoints, threePoints, fourPoints, fivePoints int) Warriors {
	return warriors{
		onePoint,
		twoPoints,
		threePoints,
		fourPoints,
		fivePoints,
	}
}
