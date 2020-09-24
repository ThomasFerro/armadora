package helpers

import (
	"fmt"
)

func CheckFinishedGame(partyId string) error {
	game, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	if string(game.State) != "Finished" {
		return fmt.Errorf(
			"Expected the game to be finished, got %v instead",
			game.State,
		)
	}
	// TODO: Check the scores
	return nil
}
