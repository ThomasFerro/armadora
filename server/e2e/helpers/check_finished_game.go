package helpers

import (
	"fmt"

	"github.com/ThomasFerro/armadora/infra/dto"
)

func CheckFinishedGame(partyId string) error {
	gameState, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	if string(gameState.State) != "Finished" {
		return fmt.Errorf(
			"Expected the game to be finished, got %v instead",
			gameState.State,
		)
	}
	return checkScores(gameState)
}

func checkScores(gameState *dto.GameDto) error {
	expectedScore := dto.ScoreDto{
		Player: 2,
		GoldStacks: []int{
			40,
		},
		TotalGold: 40,
	}
	if len(gameState.Scores) != 1 {
		return fmt.Errorf(
			"Expected to have only one score but got %v instead",
			gameState.Scores,
		)
	}
	actualScore := gameState.Scores[1]

	if actualScore.Player != expectedScore.Player ||
		actualScore.TotalGold != expectedScore.TotalGold ||
		len(actualScore.GoldStacks) != len(expectedScore.GoldStacks) ||
		actualScore.GoldStacks[0] != expectedScore.GoldStacks[0] {
		return fmt.Errorf(
			"Expected scores to be %v but got %v instead",
			expectedScore,
			gameState.Scores[0],
		)
	}
	return nil
}
