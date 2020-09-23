package helpers

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/ThomasFerro/armadora/infra/dto"
)

func CheckFinishedGame(partyId string) error {
	_, err := getFinishedGame(partyId)
	if err != nil {
		return err
	}

	// TODO: Check finished game attributes
	return nil
}

func getFinishedGame(partyId string) (*dto.GameDto, error) {
	url := fmt.Sprintf("http://localhost/parties/%v", partyId)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Check finished game - Wrong response code: %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var finishedGame dto.GameDto
	err = json.Unmarshal(body, &finishedGame)
	if err != nil {
		return nil, err
	}

	return &finishedGame, nil
}