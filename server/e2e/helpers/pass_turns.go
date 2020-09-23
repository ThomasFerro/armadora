package helpers

import (
	"github.com/ThomasFerro/armadora/infra"
)

func PassTurns(partyId string) error {
	for i := 0; i < 4; i++ {
		err := passTurn(partyId)
		if err != nil {
			return err
		}
	}
	return nil
}

func passTurn(partyId string) error {
	passTurnCommand := infra.Command{
		CommandType: "PassTurn",
	}

	return PostACommand(partyId, passTurnCommand, "Pass turn")
}