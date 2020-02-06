package infra

import (
	"log"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Parties = map[PartyId]Party{}

var eventStore = NewEventStore()

type PartyId string

type Party struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan PartyId
}

func ReceiveCommand(partyId PartyId, command Command) {
	log.Printf("Receiving the following command for party %v: %v\n", partyId, command)

	// TODO: Error management
	history, _ := eventStore.GetHistory(string(partyId))

	newEvents := ManageCommand(
		FromEventsDto(history),
		command,
	)

	eventStore.AppendToHistory(string(partyId), ToEventsDto(newEvents))

	Parties[partyId].Broadcast <- partyId
}

func CreateParty() PartyId {
	partyId := PartyId(uuid.New().String())
	history := ManageCommand([]event.Event{}, Command{
		CommandType: "CreateGame",
	})
	eventStore.AppendToHistory(string(partyId), ToEventsDto(history))
	Parties[partyId] = Party{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan PartyId),
	}
	go HandleParty(Parties[partyId])
	return partyId
}

func AddClientToParty(partyId PartyId, ws *websocket.Conn) {
	Parties[partyId].Clients[ws] = true
	Parties[partyId].Broadcast <- partyId
}

func RemoveClientFromParty(partyId PartyId, ws *websocket.Conn) {
	delete(Parties[partyId].Clients, ws)
}

func HandleParty(party Party) {
	for {
		// Grab the next message from the broadcast channel
		partyId := <-party.Broadcast
		// TODO: Error management
		history, _ := eventStore.GetHistory(string(partyId))
		updatedParty := ToGameDto(
			game.ReplayHistory(
				FromEventsDto(
					history,
				),
			),
		)
		log.Println("Receiving party message to broadcast")
		// Send it out to every client that is currently connected
		for client := range party.Clients {
			log.Println("Sending the party to client")
			err := client.WriteJSON(updatedParty)
			if err != nil {
				log.Printf("Error while sending party information: %v", err)
				client.Close()
				delete(party.Clients, client)
			}
		}
	}
}
