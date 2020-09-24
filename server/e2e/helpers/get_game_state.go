package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ThomasFerro/armadora/infra/dto"
)

func GetGameState(partyId string) (*dto.GameDto, error) {
	url := fmt.Sprintf("http://localhost/parties/%v", partyId)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Get game state - Wrong response code: %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gameState dto.GameDto
	err = json.Unmarshal(body, &gameState)
	if err != nil {
		return nil, err
	}

	return &gameState, nil
}
