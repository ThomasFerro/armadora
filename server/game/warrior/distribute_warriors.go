package warrior

// WarriorsToDistribute Returns the warriors to distribute to each player
func WarriorsToDistribute(numberOfPlayers int) Warriors {
	if numberOfPlayers == 2 {
		return NewWarriors(11, 2, 3, 1, 1)
	}

	if numberOfPlayers == 3 {
		return NewWarriors(7, 2, 1, 1, 0)
	}
	return NewWarriors(5, 1, 1, 1, 0)
}
