package party_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/infra/party"
)

const PARTIES_INTEGRATION_TEST_COLLECTION = "PARTIES_INTEGRATION_TEST_COLLECTION"

func TestCreateAPublicParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := "My new party"
	partyIsPublic := true
	newPartyID, err := partiesManager.CreateParty(partyName, partyIsPublic)

	if err != nil {
		t.Fatalf("An error has occurred while creating the party: %v", err)
	}

	if newPartyID == "" {
		t.Fatalf("The returned created party has no ID.")
	}
}

func TestCreateAPrivateParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := "My new party"
	partyIsPublic := false
	newPartyID, err := partiesManager.CreateParty(partyName, partyIsPublic)

	if err != nil {
		t.Fatalf("An error has occurred while creating the party: %v", err)
	}

	if newPartyID == "" {
		t.Fatalf("The returned created party has no ID.")
	}
}

func TestCannotCreateAPartyWithoutName(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := ""
	partyIsPublic := false
	_, err := partiesManager.CreateParty(partyName, partyIsPublic)

	if err == nil {
		t.Fatalf("An error should have occurred while creating a party without name")
	}

	if _, expectedType := err.(party.CannotCreateAPartyWithoutName); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected CannotCreateAPartyWithoutName but got %v", err)
	}
}

// TODO: Remove ?
func partiesAreEqual(firstParty, secondParty party.Party) bool {
	return firstParty.Name == secondParty.Name &&
		firstParty.Restriction == secondParty.Restriction
}

func getIntegrationTestsPartiesManager() party.PartiesManager {
	partiesRepository := party.NewPartiesMongoRepository(PARTIES_INTEGRATION_TEST_COLLECTION)
	return party.NewPartiesManager(partiesRepository)
}

func dropIntegrationTestsPartiesDatabase() {
	// TODO ? drop PARTIES_INTEGRATION_TEST_COLLECTION
}

/*
TODO:
- Create a party:
  - Creation date / ID ?
  - Cannot create the party: third party (repository) error
- Get visible parties
  - Nominal case: find every public and opened parties
  - Cannot get the parties: third party (repository) error
- Get a party
  - Nominal case: the party exists
  - Cannot get the party: third party (repository) error
  - Cannot get the party: The party does no exists
- Close a party - Use case: Once the game is finished, the party is closed
  - Nominal case: the party is closed
  - Cannot close the party: third party (repository) error

Things to do elsewhere:
- Use this party manager
- The stream_id become the party id (MVP) ? Or a party can play many games (possibily later) ?
- Call the party closer when receiving a "GameFinished" event
- Modify the API:
  - POST /parties and not POST /games
  - Returns name + id for each parties
- Modify the e2e tests
- Check that the dockerfile still works
*/
