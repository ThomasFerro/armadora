package party

// PartiesManager Exposes every possible actions on parties
type PartiesManager struct {
	repository PartiesRepository
}

// CreateParty Manage the creation of a a new party
func (partiesManager PartiesManager) CreateParty(partyName string, partyIsPublic bool) (PartyID, error) {
	if partyName == "" {
		return "", CannotCreateAPartyWithoutName{}
	}
	restriction := Private
	if partyIsPublic {
		restriction = Public
	}
	return partiesManager.repository.CreateParty(partyName, restriction)
}

// NewPartiesManager Create a new PartiesManager
func NewPartiesManager(repository PartiesRepository) PartiesManager {
	return PartiesManager{
		repository,
	}
}

// CannotCreateAPartyWithoutName Error thrown when attempting to create a party without name
type CannotCreateAPartyWithoutName struct{}

func (err CannotCreateAPartyWithoutName) Error() string {
	return "Unable to create a party without name"
}
