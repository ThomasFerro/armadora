package party_test

import (
	"context"
	"os"
	"testing"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/party"
	"github.com/ThomasFerro/armadora/infra/storage"
)

func ignoreIntegrationTests(t *testing.T) {
	if os.Getenv("IGNORE_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping integration tests")
	}
}

const PARTIES_INTEGRATION_TEST_COLLECTION = "PARTIES_INTEGRATION_TEST_COLLECTION"

func getIntegrationTestsPartiesManager() (error, party.PartiesManager, *storage.ConnectionToClose) {
	mongoClient := storage.MongoClient{
		Uri:      config.GetConfiguration("MONGO_URI"),
		Database: config.GetConfiguration("MONGO_DATABASE_NAME"),
	}
	connectionToClose, err := mongoClient.GetConnection()
	if err != nil {
		return err, party.PartiesManager{}, nil
	}
	partiesRepository := party.NewPartiesMongoRepository(connectionToClose, PARTIES_INTEGRATION_TEST_COLLECTION)
	return nil, party.NewPartiesManager(partiesRepository), connectionToClose
}

func dropIntegrationTestsPartiesDatabase(connectionToClose *storage.ConnectionToClose) {
	defer connectionToClose.Close()
	connectionToClose.Database.Collection(PARTIES_INTEGRATION_TEST_COLLECTION).Drop(context.TODO())
}

func TestCreateAPublicParty(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

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
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

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
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	partyName := party.PartyName("")
	partyIsPublic := false
	createPartyContext := context.Background()
	_, err = partiesManager.CreateParty(createPartyContext, partyName, partyIsPublic)

	if err == nil {
		t.Fatalf("An error should have occurred while creating a party without name")
	}

	if _, expectedType := err.(party.CannotCreateAPartyWithoutName); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected CannotCreateAPartyWithoutName but got %v", err)
	}
}

func TestCannotCreateAPartyWithAnAlreadyExistingName(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	alreadyExistingPartyName := party.PartyName("awesome party")
	createTestParty(t, partiesManager, alreadyExistingPartyName, false)

	newPartyName := party.PartyName("awesome party")
	newPartyIsPublic := false
	createPartyContext := context.Background()
	_, err = partiesManager.CreateParty(createPartyContext, newPartyName, newPartyIsPublic)

	if err == nil {
		t.Fatalf("An error should have occurred while creating a party with an already taken name")
	}

	if _, expectedType := err.(party.CannotCreateAPartyWithAnAlreadyTakenName); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected CannotCreateAPartyWithAnAlreadyTakenName but got %v", err)
	}
}

func TestGetVisibleParties(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	firstVisiblePartyName := party.PartyName("first visible party")
	secondVisiblePartyName := party.PartyName("second visible party")

	createTestParty(t, partiesManager, firstVisiblePartyName, true)
	createTestParty(t, partiesManager, "first private party", false)
	createTestParty(t, partiesManager, secondVisiblePartyName, true)

	getVisiblePartiesContext := context.Background()
	parties, err := partiesManager.GetVisibleParties(getVisiblePartiesContext)

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
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	visiblePartyName := party.PartyName("first visible party")
	closedPartyName := party.PartyName("closed party")
	createTestParty(t, partiesManager, visiblePartyName, true)
	createTestParty(t, partiesManager, closedPartyName, true)

	closePartyContext := context.Background()
	err = partiesManager.CloseAParty(closePartyContext, closedPartyName)

	if err != nil {
		t.Fatalf("An error has occurred while closing the party: %v", err)
	}

	getVisiblePartiesContext := context.Background()
	parties, err := partiesManager.GetVisibleParties(getVisiblePartiesContext)

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
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	partyName := party.PartyName("my awesome party")
	partyIsPublic := true
	createTestParty(t, partiesManager, partyName, partyIsPublic)

	getPartyContext := context.Background()
	newlyCreatedParty, err := partiesManager.GetParty(getPartyContext, partyName)

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
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	partyName := party.PartyName("my awesome party")
	partyIsPublic := false
	createTestParty(t, partiesManager, partyName, partyIsPublic)

	getPartyContext := context.Background()
	newlyCreatedParty, err := partiesManager.GetParty(getPartyContext, partyName)

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
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	getPartyContext := context.Background()
	partyName := party.PartyName("my awesome party")
	_, err = partiesManager.GetParty(getPartyContext, partyName)

	if err == nil {
		t.Fatalf("Should not be able to get a party that does not exists")
	}

	if _, expectedType := err.(party.NotFound); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NotFound but got %v", err)
	}
}

func TestNoPartyNameProvidedToGetTheParty(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	getPartyContext := context.Background()
	partyName := party.PartyName("")
	_, err = partiesManager.GetParty(getPartyContext, partyName)

	if err == nil {
		t.Fatalf("Should not be able to get a party that does not exists")
	}

	if _, expectedType := err.(party.NoPartyNameProvided); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NoPartyNameProvided but got %v", err)
	}
}

func TestAPartyIsOpenByDefault(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	partyName := party.PartyName("my awesome party")
	createTestParty(t, partiesManager, partyName, true)

	getPartyContext := context.Background()
	newParty, err := partiesManager.GetParty(getPartyContext, partyName)
	if err != nil {
		t.Fatalf("An error has occurred while getting the newly created party: %v", err)
	}

	if newParty.Status != party.Open {
		t.Fatalf("The created party should be opened. Got this instead: %v", newParty)
	}
}

func TestCloseAParty(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	partyName := party.PartyName("my awesome party")
	createTestParty(t, partiesManager, partyName, true)

	closePartyContext := context.Background()
	err = partiesManager.CloseAParty(closePartyContext, partyName)

	if err != nil {
		t.Fatalf("An error has occurred while closing the party: %v", err)
	}

	getPartyContext := context.Background()
	newleClosedParty, err := partiesManager.GetParty(getPartyContext, partyName)

	if err != nil {
		t.Fatalf("An error has occurred while getting the newly closed party: %v", err)
	}

	if newleClosedParty.Status != party.Close {
		t.Fatalf("The newly closed party's status is invalid. Expected to be closed but got this instead: %v", newleClosedParty)
	}
}

func TestCannotFindThePartyToClose(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	closePartyContext := context.Background()
	partyName := party.PartyName("my awesome party")
	err = partiesManager.CloseAParty(closePartyContext, partyName)

	if err == nil {
		t.Fatalf("Expected not to be able to close a party that does no exists")
	}

	if _, expectedType := err.(party.NotFound); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NotFound but got %v", err)
	}
}

func TestNoPartyNameProvidedForThePartyToClose(t *testing.T) {
	ignoreIntegrationTests(t)
	err, partiesManager, connectionToClose := getIntegrationTestsPartiesManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsPartiesDatabase(connectionToClose)

	closePartyContext := context.Background()
	partyName := party.PartyName("")
	err = partiesManager.CloseAParty(closePartyContext, partyName)

	if err == nil {
		t.Fatalf("Expected not to be able to close a party that does no exists")
	}

	if _, expectedType := err.(party.NoPartyNameProvided); !expectedType {
		t.Fatalf("The returned error is of the wrong type. Expected NoPartyNameProvided but got %v", err)
	}
}

func createTestParty(t *testing.T, partiesManager party.PartiesManager, partyName party.PartyName, partyIsPublic bool) (party.PartyName, error) {
	createPartyContext := context.Background()
	newPartyName, err := partiesManager.CreateParty(createPartyContext, partyName, partyIsPublic)

	if err != nil {
		t.Fatalf("An error has occurred while creating the party: %v", err)
	}
	return newPartyName, nil
}

func partiesAreEqual(firstParty, secondParty party.Party) bool {
	return firstParty.Name == secondParty.Name &&
		firstParty.Restriction == secondParty.Restriction
}
