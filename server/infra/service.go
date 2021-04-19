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
func (armadoraService ArmadoraService) GetVisibleParties() ([]party.PartyName, error) {
	parties, err := armadoraService.partiesManager.GetVisibleParties()
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
func (armadoraService ArmadoraService) CreateParty() (party.PartyName, error) {
	// TODO: Private parties
	createPartyContext := context.Background()

	createPartyWorkflow := func(ctx context.Context) (interface{}, error) {
		nameOfThePartyToCreate, err := generateNewPartyName(armadoraService.partiesManager)

		if err != nil {
			return nil, fmt.Errorf("An error has occurred while getting a new party name: %w", err)
		}

		newPartyName, err := armadoraService.partiesManager.CreateParty(ctx, nameOfThePartyToCreate, true)
		if err != nil {
			return nil, fmt.Errorf("An error has occurred while creating a new party: %w", err)
		}

		// TODO: extract ManageCommand + AppendToHistory + PersistProjection and reuse it in ReceiveCommand
		history, err := ManageCommand([]event.Event{}, Command{
			CommandType: "CreateGame",
		})
		if err != nil {
			return nil, fmt.Errorf("An error has occurred while creating the original event for the new party: %w", err)
		}
		err = armadoraService.eventStore.AppendToHistory(
			string(newPartyName),
			"",
			dto.ToEventsDto(history),
		)
		if err != nil {
			return nil, fmt.Errorf("An error has occurred while storing the new party: %w", err)
		}
		gameState := dto.ToGameDto(
			game.ReplayHistory(
				history,
			),
		)
		err = armadoraService.eventStore.PersistProjection(
			ctx,
			string(newPartyName),
			gameState,
		)
		if err != nil {
			return nil, fmt.Errorf("an error has occurred while storing the party's projection: %w", err)
		}
		return newPartyName, nil
	}

	returnedNewPartyName, err := armadoraService.transactionManager.RunTransation(createPartyContext, createPartyWorkflow)

	if err != nil {
		return "", fmt.Errorf("An error has occurred in the party creation transaction: %w", err)
	}

	newPartyName, returnedNewPartyNameOfTheRightType := returnedNewPartyName.(party.PartyName)
	if !returnedNewPartyNameOfTheRightType {
		return "", errors.New("Created party name type mismatch")
	}

	return newPartyName, nil
}

// GetPartyGameState Get the current state of a party's game
func (armadoraService ArmadoraService) GetPartyGameState(partyName party.PartyName) (dto.GameDto, error) {
	requestedPartyExists, err := partyExists(armadoraService.partiesManager, partyName)

	if err != nil {
		return dto.GameDto{}, fmt.Errorf("an error has occurred while checking if the party %v exists: %w", partyName, err)
	}

	if !requestedPartyExists {
		return dto.GameDto{}, fmt.Errorf("the party %v does not exists", partyName)
	}

	gameDtoToCheck, err := armadoraService.eventStore.GetProjection(context.Background(), string(partyName))
	if err != nil {
		return dto.GameDto{}, fmt.Errorf("an error has occurred while getting the party %v state: %w", partyName, err)
	}
	// gameDto, castingFailed := gameDtoToCheck.(dto.GameDto)
	// if !castingFailed {
	// 	return dto.GameDto{}, fmt.Errorf("party %v state type mismatch %v %v", partyName, gameDto, gameDtoToCheck)
	// }
	return gameDtoToCheck, nil
}

// ReceiveCommand Manage a received command
func (armadoraService ArmadoraService) ReceiveCommand(partyName party.PartyName, command Command) error {
	receiveCommandContext := context.Background()
	receiveCommandWorkflow := func(ctx context.Context) (interface{}, error) {
		requestedPartyExists, err := partyExists(armadoraService.partiesManager, partyName)

		if err != nil {
			return nil, fmt.Errorf("an error has occurred while checking if the party %v exists: %w", partyName, err)
		}

		if !requestedPartyExists {
			return nil, fmt.Errorf("the party %v does not exists", partyName)
		}

		log.Printf("Receiving the following command for party %v: %v\n", partyName, command)

		history, err := armadoraService.eventStore.GetHistory(string(partyName))

		if err != nil {
			return nil, fmt.Errorf("an error has occurred while retrieving the history before managing the command %v, %w", command, err)
		}

		historyEvents := dto.FromEventsDto(history.Events)
		newEvents, err := ManageCommand(
			historyEvents,
			command,
		)

		if err != nil {
			return nil, fmt.Errorf("an error has occurred while managing the command %v, %w", command, err)
		}

		// TODO: Put AppendToHistory + CloseAParty in a transaction
		err = armadoraService.eventStore.AppendToHistory(string(partyName), history.SequenceNumber, dto.ToEventsDto(newEvents))

		if err != nil {
			return nil, fmt.Errorf("an error has occurred while appending the events to the party's %v history, %w", partyName, err)
		}

		newGameHistory := append(historyEvents, newEvents...)
		gameState := dto.ToGameDto(game.ReplayHistory(newGameHistory))

		err = armadoraService.eventStore.PersistProjection(
			ctx,
			string(partyName),
			gameState,
		)
		if err != nil {
			return nil, fmt.Errorf("an error has occurred while storing the party's projection: %w", err)
		}

		partyNeedsToBeClosed := false

		for _, newEvent := range newEvents {
			if _, isOfRightEventType := newEvent.(event.GameFinished); isOfRightEventType {
				partyNeedsToBeClosed = true
			}
		}

		if partyNeedsToBeClosed {
			return nil, armadoraService.partiesManager.CloseAParty(partyName)
		}

		return nil, nil
	}

	_, err := armadoraService.transactionManager.RunTransation(receiveCommandContext, receiveCommandWorkflow)

	if err != nil {
		return fmt.Errorf("An error has occurred in the command management transaction: %w", err)
	}
	return nil
}

func partyExists(partiesManager party.PartiesManager, partyName party.PartyName) (bool, error) {
	_, err := partiesManager.GetParty(partyName)
	if _, partyNotFound := err.(party.NotFound); partyNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}

func generateNewPartyName(partiesManager party.PartiesManager) (party.PartyName, error) {
	tries := 1
	maxTries := 10

	for ; tries < maxTries; tries++ {
		nextPartyNameToTry := party.PartyName(
			generateNewName(),
		)
		_, err := partiesManager.GetParty(nextPartyNameToTry)
		if err == nil {
			continue
		}
		if _, partyNotFound := err.(party.NotFound); partyNotFound {
			return nextPartyNameToTry, nil
		}
		return "", err
	}

	return "", fmt.Errorf("Unable to find a new party name after %v tries", tries)
}

// NewArmadoraService Create a new Armadora service
func NewArmadoraService(eventStore storage.EventStoreWithProjection, partiesRepository party.PartiesRepository, transactionManager storage.TransactionManager) ArmadoraService {
	return ArmadoraService{
		eventStore:         eventStore,
		partiesManager:     party.NewPartiesManager(partiesRepository),
		transactionManager: transactionManager,
	}
}
