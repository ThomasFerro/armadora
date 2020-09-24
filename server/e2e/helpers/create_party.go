package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type createPartyResponse struct {
	Id string `json:"id"`
}

func CreateParty() (string, error) {
	resp, err := http.Post("http://localhost/games", "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Party creation - Wrong response code: %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var newParty createPartyResponse
	err = json.Unmarshal(body, &newParty)
	if err != nil {
		return "", err
	}
	return newParty.Id, checkGameStateAfterCreatingTheParty(newParty.Id)
}

func checkGameStateAfterCreatingTheParty(partyId string) error {
	game, err := GetGameState(partyId)
	if err != nil {
		return err
	}

	if string(game.State) != "WaitingForPlayers" {
		return fmt.Errorf("Party creating - Wrong game state, expected WaitingForPlayers but got %v instead", game.State)
	}

	return nil
}
