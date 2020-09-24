package helpers

import (
	"fmt"

	"github.com/ThomasFerro/armadora/infra"
)

type playerInformation struct {
	Nickname  string `json:"nickname"`
	Character string `json:"character"`
}

func AddPlayers(partyId string) error {
	players := []playerInformation{
		playerInformation{
			Nickname:  "Kileek",
			Character: "Goblin",
		},
		playerInformation{
			Nickname:  "Qonor",
			Character: "Elf",
		},
		playerInformation{
			Nickname:  "HackID",
			Character: "Orc",
		},
		playerInformation{
			Nickname:  "LaNinjaBanané",
			Character: "Mage",
		},
	}

	for _, player := range players {
		err := addPlayer(partyId, player)
		if err != nil {
			return err
		}
	}

	return checkGameState(partyId, players)
}

func addPlayer(partyId string, player playerInformation) error {
	addPlayerCommand := infra.Command{
		CommandType: "JoinGame",
		Payload: map[string]string{
			"Nickname":  player.Nickname,
			"Character": player.Character,
		},
	}

	return PostACommand(partyId, addPlayerCommand, "Add a player")
}

func checkGameState(partyId string, expectedPlayers []playerInformation) error {
	gameState, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	for index, actualPlayer := range gameState.Players {
		expectedPlayer := expectedPlayers[index]
		if expectedPlayer.Character != string(actualPlayer.Character) {
			return fmt.Errorf(
				"Expected player %v's character to be %v, got %v instead",
				index,
				expectedPlayer.Character,
				actualPlayer.Character,
			)
		}
		if expectedPlayer.Nickname != actualPlayer.Nickname {
			return fmt.Errorf(
				"Expected player %v's nickname to be %v, got %v instead",
				index,
				expectedPlayer.Nickname,
				actualPlayer.Nickname,
			)
		}
	}
	return nil
}
