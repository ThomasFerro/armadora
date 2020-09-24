package helpers

import (
	"fmt"

	"github.com/ThomasFerro/armadora/infra/dto"

	"github.com/ThomasFerro/armadora/infra"
)

type turnInformation struct {
	Action  string
	Payload map[string]string
}

type warriorInformation struct {
	Player   int
	Strength int
	X        int
	Y        int
}

var warriorsPlacedByPlayers = []warriorInformation{
	warriorInformation{
		Player:   0,
		Strength: 1,
		X:        2,
		Y:        0,
	},
	warriorInformation{
		Player:   1,
		Strength: 1,
		X:        2,
		Y:        1,
	},
	warriorInformation{
		Player:   2,
		Strength: 4,
		X:        2,
		Y:        2,
	},
	warriorInformation{
		Player:   3,
		Strength: 3,
		X:        0,
		Y:        0,
	},
}

type palisadeInformation struct {
	X1, Y1, X2, Y2 int
}

func (p palisadeInformation) String() string {
	return fmt.Sprintf("{%v, %v}, {%v, %v}", p.X1, p.Y1, p.X2, p.Y2)
}

var palisadesPlacedByPlayers = []palisadeInformation{
	palisadeInformation{
		X1: 2,
		Y1: 0,
		X2: 3,
		Y2: 0,
	},
	palisadeInformation{
		X1: 2,
		Y1: 1,
		X2: 3,
		Y2: 1,
	},
	palisadeInformation{
		X1: 1,
		Y1: 1,
		X2: 2,
		Y2: 1,
	},
}

func PlaySomeTurns(partyId string) error {
	turns := []turnInformation{}

	for _, warriorToPut := range warriorsPlacedByPlayers {
		turns = append(turns, turnInformation{
			Action: "PutWarrior",
			Payload: map[string]string{
				"Warrior": fmt.Sprint(warriorToPut.Strength),
				"X":       fmt.Sprint(warriorToPut.X),
				"Y":       fmt.Sprint(warriorToPut.Y),
			},
		})
	}

	turns = append(turns, turnInformation{
		Action: "PutPalisades",
		Payload: map[string]string{
			"Palisades": "[{\"x1\":2,\"y1\":0,\"x2\":3,\"y2\":0},{\"x1\":2,\"y1\":1,\"x2\":3,\"y2\":1}]",
		},
	},
		turnInformation{
			Action: "PutPalisades",
			Payload: map[string]string{
				"Palisades": "[{\"x1\":1,\"y1\":1,\"x2\":2,\"y2\":1}]",
			},
		})

	for turnIndex, turn := range turns {
		err := playTurn(partyId, turn)
		if err != nil {
			return err
		}
		err = checkTurn(partyId, turnIndex)
		if err != nil {
			return err
		}
	}

	gameState, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	err = checkPlayersState(gameState)
	if err != nil {
		return err
	}

	err = checkWarriorsInBoard(gameState)
	if err != nil {
		return err
	}

	return checkPalisadesInBoard(gameState)
}

func playTurn(partyId string, turn turnInformation) error {
	playTurnCommand := infra.Command{
		CommandType: turn.Action,
		Payload:     turn.Payload,
	}

	return PostACommand(partyId, playTurnCommand, "Play a turn")
}

func checkTurn(partyId string, previousTurnIndex int) error {
	game, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	expectedCurrentPlayer := (previousTurnIndex + 1) % 4
	if game.CurrentPlayer != expectedCurrentPlayer {
		return fmt.Errorf(
			"Expected current player to be %v but got %v instead",
			expectedCurrentPlayer,
			game.CurrentPlayer,
		)
	}
	return nil
}

func checkPlayersState(gameState *dto.GameDto) error {
	expectedWarriorsPerPlayer := []dto.WarriorsDto{
		dto.WarriorsDto{
			OnePoint:    4,
			TwoPoints:   1,
			ThreePoints: 1,
			FourPoints:  1,
			FivePoints:  0,
		},
		dto.WarriorsDto{
			OnePoint:    4,
			TwoPoints:   1,
			ThreePoints: 1,
			FourPoints:  1,
			FivePoints:  0,
		},
		dto.WarriorsDto{
			OnePoint:    5,
			TwoPoints:   1,
			ThreePoints: 1,
			FourPoints:  0,
			FivePoints:  0,
		},
		dto.WarriorsDto{
			OnePoint:    5,
			TwoPoints:   1,
			ThreePoints: 0,
			FourPoints:  1,
			FivePoints:  0,
		},
	}

	for playerIndex, expectedWarriors := range expectedWarriorsPerPlayer {
		playerWarriors := gameState.Players[playerIndex].Warriors
		if playerWarriors != expectedWarriors {
			return fmt.Errorf(
				"Exepcted player %v's warriors to be %v, got %v instead",
				playerIndex,
				expectedWarriors,
				playerWarriors,
			)
		}
	}

	return nil
}

func checkWarriorsInBoard(gameState *dto.GameDto) error {
	for _, expectedWarrior := range warriorsPlacedByPlayers {
		cell := gameState.Board.Cells[expectedWarrior.Y][expectedWarrior.X]
		if cell.Type != "warrior" {
			return fmt.Errorf("Expected cell to contain a warrior, got %v instead", cell)
		}

		expectedPlayer := gameState.Players[expectedWarrior.Player]
		if cell.Character != string(expectedPlayer.Character) {
			return fmt.Errorf(
				"Expected cell to contain a warrior of the player %v, got %v instead",
				expectedPlayer.Character,
				cell.Character,
			)
		}
	}

	return nil
}

func checkPalisadesInBoard(gameState *dto.GameDto) error {
	for _, expectedPalisade := range palisadesPlacedByPlayers {
		palisadeFound := false
		actualPalisades := gameState.Board.Palisades
		for _, actualPalisade := range actualPalisades {
			if actualPalisade.X1 == expectedPalisade.X1 &&
				actualPalisade.Y1 == expectedPalisade.Y1 &&
				actualPalisade.X2 == expectedPalisade.X2 &&
				actualPalisade.Y2 == expectedPalisade.Y2 {
				palisadeFound = true
				break
			}
		}

		if !palisadeFound {
			return fmt.Errorf(
				"Expected the board to contain the palisade %v but only contains %v",
				expectedPalisade,
				gameState.Board.Palisades,
			)
		}
	}

	return nil
}
