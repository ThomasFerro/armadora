package infra

import (
	"fmt"
	"log"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/party"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/infra/dto"
	"github.com/ThomasFerro/armadora/infra/storage"
	"github.com/google/uuid"
)

// ArmadoraService Service managing Armadora games
type ArmadoraService struct {
	eventStore     storage.EventStore
	partiesManager party.PartiesManager
}

// NewArmadoraService Create a new Armadora service
func NewArmadoraService() ArmadoraService {
	partiesRepository := party.NewPartiesMongoRepository(
		config.GetConfiguration("MONGO_PARTY_COLLECTION_NAME"),
	)

	return ArmadoraService{
		eventStore:     storage.NewEventStore(),
		partiesManager: party.NewPartiesManager(partiesRepository),
	}
}

// ReceiveCommand Manage a received command
func (armadoraService ArmadoraService) ReceiveCommand(partyName party.PartyName, command Command) error {
	log.Printf("Receiving the following command for party %v: %v\n", partyName, command)

	history, err := armadoraService.eventStore.GetHistory(string(partyName))

	if err != nil {
		return fmt.Errorf("An error has occurred while retrieving the history before managing the command %v, %w", command, err)
	}

	newEvents, err := ManageCommand(
		dto.FromEventsDto(history.Events),
		command,
	)

	if err != nil {
		return fmt.Errorf("An error has occurred while managing the command %v, %w", command, err)
	}

	err = armadoraService.eventStore.AppendToHistory(string(partyName), history.SequenceNumber, dto.ToEventsDto(newEvents))

	if err != nil {
		return fmt.Errorf("An error has occurred while appending the events to the party's %v history, %w", partyName, err)
	}

	partyNeedsToBeClosed := false

	for _, newEvent := range newEvents {
		if _, isOfRightEventType := newEvent.(event.GameFinished); isOfRightEventType {
			partyNeedsToBeClosed = true
		}
	}

	if partyNeedsToBeClosed {
		return armadoraService.partiesManager.CloseAParty(partyName)
	}

	return nil
}

// CreateParty Create a new party
func (armadoraService ArmadoraService) CreateParty() (party.PartyName, error) {
	// TODO: Shorter name
	// TODO: Private parties
	partyName := party.PartyName(uuid.New().String())
	newPartyName, err := armadoraService.partiesManager.CreateParty(partyName, true)
	if err != nil {
		return "", err
	}
	history, err := ManageCommand([]event.Event{}, Command{
		CommandType: "CreateGame",
	})
	if err != nil {
		return "", err
	}
	err = armadoraService.eventStore.AppendToHistory(
		string(newPartyName),
		"",
		dto.ToEventsDto(history),
	)
	if err != nil {
		return "", err
	}
	return newPartyName, nil
}

// GetPartyGameState Get the current state of a party's game
func (armadoraService ArmadoraService) GetPartyGameState(partyName party.PartyName) (dto.GameDto, error) {
	history, err := armadoraService.eventStore.GetHistory(string(partyName))
	if err != nil {
		return dto.GameDto{}, err
	}
	return dto.ToGameDto(
		game.ReplayHistory(
			dto.FromEventsDto(history.Events),
		),
	), nil
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
