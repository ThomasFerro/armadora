package helpers

import (
	"fmt"

	"github.com/ThomasFerro/armadora/infra"
)

func PassTurns(partyId string, currentPlayer int) error {
	for i := 0; i < 4; i++ {
		currentPlayerToPass := (currentPlayer + i) % 4
		err := passTurn(partyId, currentPlayerToPass)
		if err != nil {
			return err
		}
	}
	return checkGameStateAfterPassingTurns(partyId)
}

func passTurn(partyId string, player int) error {
	passTurnCommand := infra.Command{
		CommandType: "PassTurn",
		Player:      player,
	}

	return PostACommand(partyId, passTurnCommand, "Pass turn")
}

func checkGameStateAfterPassingTurns(partyId string) error {
	game, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	for index, player := range game.Players {
		if !player.TurnPassed {
			return fmt.Errorf("Expected player %v's turn to be passed", index)
		}
	}

	return nil
}
