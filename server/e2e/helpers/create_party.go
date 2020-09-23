package helpers

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"
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
	return newParty.Id, nil
}