package helpers

import (
	"github.com/ThomasFerro/armadora/infra"
)

type turnInformation struct {
	Action string
	Payload map[string]string
}

func PlaySomeTurns(partyId string) error {
	turns := []turnInformation{
		turnInformation{
			Action: "PutWarrior",
			Payload: map[string]string{
				"Warrior": "1",
				"X": "2",
				"Y": "0",
			},
		},
		turnInformation{
			Action: "PutWarrior",
			Payload: map[string]string{
				"Warrior": "1",
				"X": "2",
				"Y": "1",
			},
		},
		turnInformation{
			Action: "PutWarrior",
			Payload: map[string]string{
				"Warrior": "4",
				"X": "2",
				"Y": "2",
			},
		},
		turnInformation{
			Action: "PutWarrior",
			Payload: map[string]string{
				"Warrior": "3",
				"X": "0",
				"Y": "0",
			},
		},
		turnInformation{
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
		},
	}

	for _, turn := range turns {
		err := playTurn(partyId, turn)
		if err != nil {
			return err
		}
	}
	return nil
}

func playTurn(partyId string, turn turnInformation) error {
	playTurnCommand := infra.Command{
		CommandType: turn.Action,
		Payload: turn.Payload,
	}

	return PostACommand(partyId, playTurnCommand, "Play a turn")
}