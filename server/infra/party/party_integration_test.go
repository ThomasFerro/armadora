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

	partyName := party.PartyName("My new party")
	partyIsPublic := true
	newPartyName, _ := createTestParty(t, partiesManager, partyName, partyIsPublic)

	if newPartyName == "" {
		t.Fatalf("The returned created party has no name.")
	}

	if newPartyName != partyName {
		t.Fatalf("Created party name (%v) does not match with the provided one (%v).", newPartyName, partyName)
	}
}

func TestCreateAPrivateParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("My new party")
	partyIsPublic := false
	newPartyName, _ := createTestParty(t, partiesManager, partyName, partyIsPublic)

	if newPartyName == "" {
		t.Fatalf("The returned created party has no name.")
	}

	if newPartyName != partyName {
		t.Fatalf("Created party name (%v) does not match with the provided one (%v).", newPartyName, partyName)
	}
}

func TestCannotCreateAPartyWithoutName(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("")
	partyIsPublic := false
	_, err := partiesManager.CreateParty(partyName, partyIsPublic)

	if err == nil {
		t.Fatalf("An error should have occurred while creating a party without name")
	}

	if _, expectedType := err.(party.CannotCreateAPartyWithoutName); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected CannotCreateAPartyWithoutName but got %v", err)
	}
}

func TestCannotCreateAPartyWithAnAlreadyExistingName(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	alreadyExistingPartyName := party.PartyName("awesome party")
	createTestParty(t, partiesManager, alreadyExistingPartyName, false)

	newPartyName := party.PartyName("awesome party")
	newPartyIsPublic := false
	_, err := partiesManager.CreateParty(newPartyName, newPartyIsPublic)

	if err == nil {
		t.Fatalf("An error should have occurred while creating a party with an already taken name")
	}

	if _, expectedType := err.(party.CannotCreateAPartyWithAnAlreadyTakenName); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected CannotCreateAPartyWithAnAlreadyTakenName but got %v", err)
	}
}

func TestGetVisibleParties(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	firstVisiblePartyName := party.PartyName("first visible party")
	secondVisiblePartyName := party.PartyName("second visible party")

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

func TestGetASpecificPublicParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("my awesome party")
	partyIsPublic := true
	createTestParty(t, partiesManager, partyName, partyIsPublic)

	newlyCreatedParty, err := partiesManager.GetParty(partyName)

	if err != nil {
		t.Fatalf("Unable to get the newly created party: %v", err)
	}

	expectedParty := party.Party{
		Name:        partyName,
		Restriction: party.Public,
	}

	if !partiesAreEqual(newlyCreatedParty, expectedParty) {
		t.Fatalf("Invalid public party. Expected %v but got this instead: %v", expectedParty, newlyCreatedParty)
	}
}

func TestGetASpecificPrivateParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("my awesome party")
	partyIsPublic := false
	createTestParty(t, partiesManager, partyName, partyIsPublic)

	newlyCreatedParty, err := partiesManager.GetParty(partyName)

	if err != nil {
		t.Fatalf("Unable to get the newly created party: %v", err)
	}

	expectedParty := party.Party{
		Name:        partyName,
		Restriction: party.Private,
	}

	if !partiesAreEqual(newlyCreatedParty, expectedParty) {
		t.Fatalf("Invalid private party. Expected %v but got this instead: %v", expectedParty, newlyCreatedParty)
	}
}

func TestCannotGetAPartyThatDoesNotExist(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("my awesome party")
	_, err := partiesManager.GetParty(partyName)

	if err == nil {
		t.Fatalf("Should not be able to get a party that does not exists")
	}

	if _, expectedType := err.(party.NotFound); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NotFound but got %v", err)
	}
}

func createTestParty(t *testing.T, partyManager party.PartiesManager, partyName party.PartyName, partyIsPublic bool) (party.PartyName, error) {
	newPartyName, err := partyManager.CreateParty(partyName, partyIsPublic)

	if err != nil {
		t.Fatalf("An error has occurred while creating the party: %v", err)
	}
	return newPartyName, nil
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
