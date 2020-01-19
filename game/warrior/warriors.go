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
