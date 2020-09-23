package e2e

import (
	"testing"

	"github.com/ThomasFerro/armadora/infra/api"
	"github.com/ThomasFerro/armadora/e2e/helpers"
)

func TestEndToEnd(t *testing.T) {
	go api.StartApi()

	partyId, err := helpers.CreateParty()
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.AddPlayers(partyId)
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.StartTheGame(partyId)
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.PlaySomeTurns(partyId) 
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.PassTurns(partyId) 
	if err != nil {
		t.Fatal(err)
	}

	helpers.CheckFinishedGame(partyId)
	if err != nil {
		t.Fatal(err)
	}
}
