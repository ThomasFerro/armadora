package infra

import (
	"log"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/infra/dto"
	"github.com/google/uuid"
)

var Parties = []PartyId{}

var eventStore = NewEventStore()

type PartyId string

func ReceiveCommand(partyId PartyId, command Command) error {
	log.Printf("Receiving the following command for party %v: %v\n", partyId, command)

	history, err := eventStore.GetHistory(string(partyId))

	if err != nil {
		return err
	}

	newEvents, err := ManageCommand(
		dto.FromEventsDto(history),
		command,
	)

	if err != nil {
		return err
	}

	eventStore.AppendToHistory(string(partyId), dto.ToEventsDto(newEvents))

	return nil
}

func CreateParty() (PartyId, error) {
	partyId := PartyId(uuid.New().String())
	history, err := ManageCommand([]event.Event{}, Command{
		CommandType: "CreateGame",
	})
	if err != nil {
		return "", err
	}
	Parties = append(Parties, partyId)
	err = eventStore.AppendToHistory(string(partyId), dto.ToEventsDto(history))
	if err != nil {
		return "", err
	}
	return partyId, nil
}

func GetParty(partyId PartyId) (dto.GameDto, error) {
	history, err := eventStore.GetHistory(string(partyId))
	if err != nil {
		return dto.GameDto{}, err
	}
	return dto.ToGameDto(
		game.ReplayHistory(
			dto.FromEventsDto(history),
		),
	), nil
}
