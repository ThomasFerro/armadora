package exception

// GameAlreadyFull No more room for player in the game
type GameAlreadyFull struct{}

// Error Indicate that the game is already full
func (exception GameAlreadyFull) Error() string {
	return "The game is already full, no more player can join it."
}
