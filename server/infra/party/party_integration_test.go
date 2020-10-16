package party_test

import (
	"context"
	"testing"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/party"
	"github.com/ThomasFerro/armadora/infra/storage"
)

const PARTIES_INTEGRATION_TEST_COLLECTION = "PARTIES_INTEGRATION_TEST_COLLECTION"

func TestCreateAPublicParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := "My new party"
	partyIsPublic := true
	newPartyID, _ := createTestParty(t, partiesManager, partyName, partyIsPublic)

	if newPartyID == "" {
		t.Fatalf("The returned created party has no ID.")
	}
}

func TestCreateAPrivateParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := "My new party"
	partyIsPublic := false
	newPartyID, _ := createTestParty(t, partiesManager, partyName, partyIsPublic)

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

func TestGetVisibleParties(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	firstVisiblePartyName := "first visible party"
	secondVisiblePartyName := "second visible party"

	createTestParty(t, partiesManager, firstVisiblePartyName, true)
	createTestParty(t, partiesManager, "first private party", false)
	createTestParty(t, partiesManager, secondVisiblePartyName, true)

	parties, err := partiesManager.GetVisibleParties()

	if err != nil {
		t.Fatalf("An error has occurred while getting visible parties: %v", err)
	}

	if len(parties) != 2 {
		t.Fatalf("Only the two visible parties should have been returned. Got this instead: %v", parties)
	}

	expectedFirstParty := party.Party{
		Name:        firstVisiblePartyName,
		Restriction: party.Public,
	}

	if !partiesAreEqual(parties[0], expectedFirstParty) {
		t.Fatalf("Invalid first visible party. Expected %v but got this instead: %v", expectedFirstParty, parties[0])
	}

	expectedSecondParty := party.Party{
		Name:        secondVisiblePartyName,
		Restriction: party.Public,
	}

	if !partiesAreEqual(parties[1], expectedSecondParty) {
		t.Fatalf("Invalid second visible party. Expected %v but got this instead: %v", expectedSecondParty, parties[1])
	}
}

func createTestParty(t *testing.T, partyManager party.PartiesManager, partyName string, partyIsPublic bool) (party.PartyID, error) {
	newPartyID, err := partyManager.CreateParty(partyName, partyIsPublic)

	if err != nil {
		t.Fatalf("An error has occurred while creating the party: %v", err)
	}
	return newPartyID, nil
}

func partiesAreEqual(firstParty, secondParty party.Party) bool {
	return firstParty.Name == secondParty.Name &&
		firstParty.Restriction == secondParty.Restriction
}

func getIntegrationTestsPartiesManager() party.PartiesManager {
	partiesRepository := party.NewPartiesMongoRepository(PARTIES_INTEGRATION_TEST_COLLECTION)
	return party.NewPartiesManager(partiesRepository)
}

func dropIntegrationTestsPartiesDatabase() {
	mongoClient := storage.MongoClient{
		Uri:        config.GetConfiguration("MONGO_URI"),
		Database:   config.GetConfiguration("MONGO_DATABASE_NAME"),
		Collection: PARTIES_INTEGRATION_TEST_COLLECTION,
	}
	collectionToClose, err := mongoClient.GetCollection()
	if err != nil {
		return
	}
	defer collectionToClose.Close()
	collectionToClose.Collection.Drop(context.TODO())
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
  - Add a "status" to the party, open by default
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
