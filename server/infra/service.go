package infra

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ThomasFerro/armadora/infra/party"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/infra/dto"
	"github.com/ThomasFerro/armadora/infra/storage"
)

// ArmadoraService Service managing Armadora games
type ArmadoraService struct {
	eventStore         storage.EventStoreWithProjection
	partiesManager     party.PartiesManager
	transactionManager storage.TransactionManager
}

// GetVisibleParties Retrieve every available parties
func (armadoraService ArmadoraService) GetVisibleParties(getVisiblePartiesContext context.Context) ([]party.PartyName, error) {
	parties, err := armadoraService.partiesManager.GetVisibleParties(getVisiblePartiesContext)
	if err != nil {
		return nil, err
	}
	returnedParties := []party.PartyName{}
	for _, partyFromEventStore := range parties {
		returnedParties = append(returnedParties, partyFromEventStore.Name)
	}
	return returnedParties, nil
}

// CreateParty Create a new party
func (armadoraService ArmadoraService) CreateParty(createPartyContext context.Context) (party.PartyName, error) {
	// TODO: Private parties
	createPartyWorkflow := func(ctx context.Context) (interface{}, error) {
		nameOfThePartyToCreate, err := generateNewPartyName(ctx, armadoraService.partiesManager)

		if err != nil {
			return nil, fmt.Errorf("an error has occurred while getting a new party name: %w", err)
		}

		newPartyName, err := armadoraService.partiesManager.CreateParty(ctx, nameOfThePartyToCreate, true)
		if err != nil {
			return nil, fmt.Errorf("an error has occurred while creating a new party: %w", err)
		}

		_, err = manageCommand(ctx, armadoraService, newPartyName, []event.Event{}, "", Command{
			CommandType: "CreateGame",
		})
		return newPartyName, err
	}

	returnedNewPartyName, err := armadoraService.transactionManager.RunTransation(createPartyContext, createPartyWorkflow)

	if err != nil {
		return "", fmt.Errorf("an error has occurred in the party creation transaction: %w", err)
	}

	newPartyName, returnedNewPartyNameOfTheRightType := returnedNewPartyName.(party.PartyName)
	if !returnedNewPartyNameOfTheRightType {
		return "", errors.New("created party name type mismatch")
	}

	return newPartyName, nil
}

// GetPartyGameState Get the current state of a party's game
func (armadoraService ArmadoraService) GetPartyGameState(getPartyGameStateContext context.Context, partyName party.PartyName) (dto.GameDto, error) {
	requestedPartyExists, err := partyExists(getPartyGameStateContext, armadoraService.partiesManager, partyName)

	if err != nil {
		return dto.GameDto{}, fmt.Errorf("an error has occurred while checking if the party %v exists: %w", partyName, err)
	}

	if !requestedPartyExists {
		return dto.GameDto{}, fmt.Errorf("the party %v does not exists", partyName)
	}

	gameDtoToCheck, err := armadoraService.eventStore.GetProjection(getPartyGameStateContext, string(partyName))
	if err != nil {
		return dto.GameDto{}, fmt.Errorf("an error has occurred while getting the party %v state: %w", partyName, err)
	}
	// TODO: Using an interface{} in the projection store make the cast not working :/
	// gameDto, castingFailed := gameDtoToCheck.(dto.GameDto)
	// if !castingFailed {
	// 	return dto.GameDto{}, fmt.Errorf("party %v state type mismatch %v %v", partyName, gameDto, gameDtoToCheck)
	// }
	return gameDtoToCheck, nil
}

// ReceiveCommand Manage a received command
func (armadoraService ArmadoraService) ReceiveCommand(receiveCommandContext context.Context, partyName party.PartyName, command Command) error {
	receiveCommandWorkflow := func(ctx context.Context) (interface{}, error) {
		requestedPartyExists, err := partyExists(ctx, armadoraService.partiesManager, partyName)

		if err != nil {
			return nil, fmt.Errorf("an error has occurred while checking if the party %v exists: %w", partyName, err)
		}

		if !requestedPartyExists {
			return nil, fmt.Errorf("the party %v does not exists", partyName)
		}

		log.Printf("Receiving the following command for party %v: %v\n", partyName, command)

		history, err := armadoraService.eventStore.GetHistory(ctx, string(partyName))

		if err != nil {
			return nil, fmt.Errorf("an error has occurred while retrieving the history before managing the command %v, %w", command, err)
		}

		historyEvents := dto.FromEventsDto(history.Events)
		newEvents, err := manageCommand(ctx, armadoraService, partyName, historyEvents, history.SequenceNumber, command)
		if err != nil {
			return nil, fmt.Errorf("an error has occurred while managing the command %v for the party %v, %w", command, partyName, err)
		}

		partyNeedsToBeClosed := false

		for _, newEvent := range newEvents {
			if _, isOfRightEventType := newEvent.(event.GameFinished); isOfRightEventType {
				partyNeedsToBeClosed = true
			}
		}

		if partyNeedsToBeClosed {
			return nil, armadoraService.partiesManager.CloseAParty(ctx, partyName)
		}

		return nil, nil
	}

	_, err := armadoraService.transactionManager.RunTransation(receiveCommandContext, receiveCommandWorkflow)

	if err != nil {
		return fmt.Errorf("an error has occurred in the command management transaction for the party %v: %w", partyName, err)
	}
	return nil
}

func partyExists(ctx context.Context, partiesManager party.PartiesManager, partyName party.PartyName) (bool, error) {
	_, err := partiesManager.GetParty(ctx, partyName)
	if _, partyNotFound := err.(party.NotFound); partyNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}

func generateNewPartyName(ctx context.Context, partiesManager party.PartiesManager) (party.PartyName, error) {
	tries := 1
	maxTries := 10

	for ; tries < maxTries; tries++ {
		nextPartyNameToTry := party.PartyName(
			generateNewName(),
		)
		_, err := partiesManager.GetParty(ctx, nextPartyNameToTry)
		if err == nil {
			continue
		}
		if _, partyNotFound := err.(party.NotFound); partyNotFound {
			return nextPartyNameToTry, nil
		}
		return "", err
	}

	return "", fmt.Errorf("unable to find a new party name after %v tries", tries)
}

func manageCommand(ctx context.Context, armadoraService ArmadoraService, partyName party.PartyName, eventsHistory []event.Event, sequenceNumber storage.SequenceNumber, command Command) ([]event.Event, error) {
	newEvents, err := ManageCommand(eventsHistory, command)
	if err != nil {
		return nil, fmt.Errorf("an error has occurred while managing the command for the party %v: %w", partyName, err)
	}

	newHistory := append(eventsHistory, newEvents...)
	err = armadoraService.eventStore.AppendToHistory(
		ctx,
		string(partyName),
		sequenceNumber,
		dto.ToEventsDto(newEvents),
	)
	if err != nil {
		return nil, fmt.Errorf("an error has occurred while storing the new events for the party %v: %w", partyName, err)
	}

	gameState := dto.ToGameDto(
		game.ReplayHistory(
			newHistory,
		),
	)
	err = armadoraService.eventStore.PersistProjection(
		ctx,
		string(partyName),
		gameState,
	)
	if err != nil {
		return nil, fmt.Errorf("an error has occurred while storing the party's %v projection: %w", partyName, err)
	}

	return newEvents, nil
}

// NewArmadoraService Create a new Armadora service
func NewArmadoraService(eventStore storage.EventStoreWithProjection, partiesRepository party.PartiesRepository, transactionManager storage.TransactionManager) ArmadoraService {
	return ArmadoraService{
		eventStore:         eventStore,
		partiesManager:     party.NewPartiesManager(partiesRepository),
		transactionManager: transactionManager,
	}
}
