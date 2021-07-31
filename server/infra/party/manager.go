package party

import (
	"context"
	"fmt"
)

// PartiesManager Exposes every possible actions on parties
type PartiesManager struct {
	repository PartiesRepository
}

// CreateParty Manage the creation of a a new party
func (partiesManager PartiesManager) CreateParty(ctx context.Context, partyName PartyName, partyIsPublic bool) (PartyName, error) {
	if partyName == "" {
		return "", CannotCreateAPartyWithoutName{}
	}
	_, err := partiesManager.repository.GetParty(ctx, partyName)
	if err == nil {
		return "", CannotCreateAPartyWithAnAlreadyTakenName{
			name: partyName,
		}
	}
	if _, partyNotFound := err.(NotFound); !partyNotFound {
		return "", fmt.Errorf("an error has occurred while checking if the party already exists: %w", err)
	}

	restriction := Private
	if partyIsPublic {
		restriction = Public
	}
	return partiesManager.repository.CreateParty(ctx, partyName, restriction, Open)
}

// GetVisibleParties Get all visible parties
func (partiesManager PartiesManager) GetVisibleParties(ctx context.Context) ([]Party, error) {
	return partiesManager.repository.GetParties(ctx, Public, Open)
}

// GetParty Get a specific parties
func (partiesManager PartiesManager) GetParty(ctx context.Context, partyName PartyName) (Party, error) {
	if partyName == PartyName("") {
		return Party{}, NoPartyNameProvided{}
	}

	return partiesManager.repository.GetParty(ctx, partyName)
}

// CloseAParty Close a party
func (partiesManager PartiesManager) CloseAParty(ctx context.Context, partyName PartyName) error {
	if partyName == PartyName("") {
		return NoPartyNameProvided{}
	}

	partyToClose, err := partiesManager.repository.GetParty(ctx, partyName)

	if err != nil {
		return NotFound{
			partyName,
		}
	}
	closedParty := partyToClose.Close()

	return partiesManager.repository.UpdateParty(ctx, closedParty)
}

// NewPartiesManager Create a new PartiesManager
func NewPartiesManager(repository PartiesRepository) PartiesManager {
	return PartiesManager{
		repository,
	}
}
