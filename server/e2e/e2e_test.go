package e2e

import (
	"os"
	"testing"
	"time"

	"github.com/ThomasFerro/armadora/e2e/helpers"
	"github.com/ThomasFerro/armadora/infra/api"
)

func ignoreIntegrationTests(t *testing.T) {
	if os.Getenv("IGNORE_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping integration tests")
	}
}

func TestEndToEnd(t *testing.T) {
	ignoreIntegrationTests(t)
	go api.StartApi()

	time.Sleep(5 * time.Second)

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

	game, err := helpers.GetGameState(partyId)
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.PassTurns(partyId, game.CurrentPlayer)
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.CheckFinishedGame(partyId)
	if err != nil {
		t.Fatal(err)
	}
}
