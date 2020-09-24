package helpers

import (
	"github.com/ThomasFerro/armadora/infra"
)

func StartTheGame(partyId string) error {
	startTheGameCommand := infra.Command{
		CommandType: "StartTheGame",
	}

	return PostACommand(partyId, startTheGameCommand, "Start the game")
}