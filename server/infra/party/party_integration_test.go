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
func TestClosedPartiesNotConsideredAsVisible(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	visiblePartyName := party.PartyName("first visible party")
	closedPartyName := party.PartyName("closed party")
	createTestParty(t, partiesManager, visiblePartyName, true)
	createTestParty(t, partiesManager, closedPartyName, true)

	err := partiesManager.CloseAParty(closedPartyName)

	if err != nil {
		t.Fatalf("An error has occurred while closing the party: %v", err)
	}

	parties, err := partiesManager.GetVisibleParties()

	if err != nil {
		t.Fatalf("An error has occurred while getting visible parties: %v", err)
	}

	if len(parties) != 1 {
		t.Fatalf("Only the visible and open party should have been returned. Got this instead: %v", parties)
	}

	expectedVisibleParty := party.Party{
		Name:        visiblePartyName,
		Restriction: party.Public,
	}

	if !partiesAreEqual(parties[0], expectedVisibleParty) {
		t.Fatalf("Invalid only visible party. Expected %v but got this instead: %v", expectedVisibleParty, parties[0])
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

func TestNoPartyNameProvidedToGetTheParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("")
	_, err := partiesManager.GetParty(partyName)

	if err == nil {
		t.Fatalf("Should not be able to get a party that does not exists")
	}

	if _, expectedType := err.(party.NoPartyNameProvided); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NoPartyNameProvided but got %v", err)
	}
}

func TestAPartyIsOpenByDefault(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("my awesome party")
	createTestParty(t, partiesManager, partyName, true)

	newParty, err := partiesManager.GetParty(partyName)
	if err != nil {
		t.Fatalf("An error has occurred while getting the newly created party: %v", err)
	}

	if newParty.Status != party.Open {
		t.Fatalf("The created party should be opened. Got this instead: %v", newParty)
	}
}

func TestCloseAParty(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("my awesome party")
	createTestParty(t, partiesManager, partyName, true)

	err := partiesManager.CloseAParty(partyName)

	if err != nil {
		t.Fatalf("An error has occurred while closing the party: %v", err)
	}

	newleClosedParty, err := partiesManager.GetParty(partyName)

	if err != nil {
		t.Fatalf("An error has occurred while getting the newly closed party: %v", err)
	}

	if newleClosedParty.Status != party.Close {
		t.Fatalf("The newly closed party's status is invalid. Expected to be closed but got this instead: %v", newleClosedParty)
	}
}

func TestCannotFindThePartyToClose(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("my awesome party")
	err := partiesManager.CloseAParty(partyName)

	if err == nil {
		t.Fatalf("Expected not to be able to close a party that does no exists")
	}

	if _, expectedType := err.(party.NotFound); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NotFound but got %v", err)
	}
}

func TestNoPartyNameProvidedForThePartyToClose(t *testing.T) {
	partiesManager := getIntegrationTestsPartiesManager()
	defer dropIntegrationTestsPartiesDatabase()

	partyName := party.PartyName("")
	err := partiesManager.CloseAParty(partyName)

	if err == nil {
		t.Fatalf("Expected not to be able to close a party that does no exists")
	}

	if _, expectedType := err.(party.NoPartyNameProvided); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NoPartyNameProvided but got %v", err)
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
Things to do elsewhere:
- Modify the e2e tests
- Check that the dockerfile still works
*/
