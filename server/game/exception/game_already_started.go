package exception

// GameAlreadyStarted Dispatched if a player try to join the game after it started
type GameAlreadyStarted struct{}

func (exception GameAlreadyStarted) Error() string {
	return "A player cannot join the game after it started"
}
