package infra

import (
	"fmt"
	"log"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/infra/dto"
	"github.com/ThomasFerro/armadora/infra/storage"
	"github.com/google/uuid"
)

var eventStore = storage.NewEventStore()

// PartyID A party identifier
type PartyID string

// ReceiveCommand Manage a received command
func ReceiveCommand(partyID PartyID, command Command) error {
	log.Printf("Receiving the following command for party %v: %v\n", partyID, command)

	history, err := eventStore.GetHistory(string(partyID))

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

	eventStore.AppendToHistory(string(partyID), history.SequenceNumber, dto.ToEventsDto(newEvents))

	return nil
}

// CreateParty Create a new party
func CreateParty() (PartyID, error) {
	partyID := PartyID(uuid.New().String())
	history, err := ManageCommand([]event.Event{}, Command{
		CommandType: "CreateGame",
	})
	if err != nil {
		return "", err
	}
	err = eventStore.AppendToHistory(
		string(partyID),
		"",
		dto.ToEventsDto(history),
	)
	if err != nil {
		return "", err
	}
	return partyID, nil
}

// GetParty Get the current state of a party
func GetParty(partyID PartyID) (dto.GameDto, error) {
	history, err := eventStore.GetHistory(string(partyID))
	if err != nil {
		return dto.GameDto{}, err
	}
	return dto.ToGameDto(
		game.ReplayHistory(
			dto.FromEventsDto(history.Events),
		),
	), nil
}

// GetParties Retrieve every available parties
func GetParties() ([]PartyID, error) {
	parties, err := eventStore.GetParties()
	if err != nil {
		return nil, err
	}
	returnedParties := []PartyID{}
	for _, party := range parties {
		returnedParties = append(returnedParties, PartyID(party))
	}
	return returnedParties, nil
}
