package infra

import (
	"log"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
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

	newEvents := ManageCommand(
		FromEventsDto(history),
		command,
	)

	eventStore.AppendToHistory(string(partyId), ToEventsDto(newEvents))

	return nil
}

func CreateParty() PartyId {
	partyId := PartyId(uuid.New().String())
	history := ManageCommand([]event.Event{}, Command{
		CommandType: "CreateGame",
	})
	Parties = append(Parties, partyId)
	eventStore.AppendToHistory(string(partyId), ToEventsDto(history))
	return partyId
}

func GetParty(partyId PartyId) (GameDto, error) {
	history, err := eventStore.GetHistory(string(partyId))
	if err != nil {
		return GameDto{}, err
	}
	return ToGameDto(
		game.ReplayHistory(
			FromEventsDto(history),
		),
	), nil
}
