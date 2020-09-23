package helpers

import (
	"github.com/ThomasFerro/armadora/infra"
)

type playerInformation struct {
	Nickname  string `json:"nickname"`
	Character string    `json:"character"`
}

func AddPlayers(partyId string) error {
	players := []playerInformation{
		playerInformation{
			Nickname: "Kileek",
			Character: "Goblin",
		},
		playerInformation{
			Nickname: "Qonor",
			Character: "Elf",
		},
		playerInformation{
			Nickname: "HackID",
			Character: "Orc",
		},
		playerInformation{
			Nickname: "LaNinjaBanané",
			Character: "Mage",
		},
	}

	for _, player := range players {
		err := addPlayer(partyId, player)
		if err != nil {
			return err
		}
	}
	return nil
}

func addPlayer(partyId string, player playerInformation) error {
	addPlayerCommand := infra.Command{
		CommandType: "JoinGame",
		Payload: map[string]string{
			"Nickname": player.Nickname,
			"Character": player.Character,
		},
	}

	return PostACommand(partyId, addPlayerCommand, "Add a player")
}