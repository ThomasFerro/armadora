package party

import (
	"fmt"
)

// PartiesManager Exposes every possible actions on parties
type PartiesManager struct {
	repository PartiesRepository
}

// CreateParty Manage the creation of a a new party
func (partiesManager PartiesManager) CreateParty(partyName PartyName, partyIsPublic bool) (PartyName, error) {
	if partyName == "" {
		return "", CannotCreateAPartyWithoutName{}
	}
	_, err := partiesManager.repository.GetParty(partyName)
	if err == nil {
		return "", CannotCreateAPartyWithAnAlreadyTakenName{
			name: partyName,
		}
	}
	if _, partyNotFound := err.(NotFound); !partyNotFound {
		return "", fmt.Errorf("An error has occurred while checking if the party already exists: %w", err)
	}

	restriction := Private
	if partyIsPublic {
		restriction = Public
	}
	return partiesManager.repository.CreateParty(partyName, restriction, Open)
}

// GetVisibleParties Get all visible parties
func (partiesManager PartiesManager) GetVisibleParties() ([]Party, error) {
	return partiesManager.repository.GetParties(Public, Open)
}

// GetParty Get a specific parties
func (partiesManager PartiesManager) GetParty(partyName PartyName) (Party, error) {
	if partyName == PartyName("") {
		return Party{}, NoPartyNameProvided{}
	}

	return partiesManager.repository.GetParty(partyName)
}

// CloseAParty Close a party
func (partiesManager PartiesManager) CloseAParty(partyName PartyName) error {
	if partyName == PartyName("") {
		return NoPartyNameProvided{}
	}

	partyToClose, err := partiesManager.repository.GetParty(partyName)

	if err != nil {
		return NotFound{
			partyName,
		}
	}
	closedParty := partyToClose.Close()

	return partiesManager.repository.UpdateParty(closedParty)
}

// NewPartiesManager Create a new PartiesManager
func NewPartiesManager(repository PartiesRepository) PartiesManager {
	return PartiesManager{
		repository,
	}
}
