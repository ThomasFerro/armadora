package infra

import (
	"log"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Parties = map[PartyId]Party{}

type PartyId string

type Party struct {
	History   []event.Event
	Clients   map[*websocket.Conn]bool
	Broadcast chan GameDto
}

func CreateParty() PartyId {
	partyId := PartyId(uuid.New().String())
	Parties[partyId] = Party{
		History:   []event.Event{},
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan GameDto),
	}
	go HandleParty(Parties[partyId])
	return partyId
}

func AddClientToParty(partyId PartyId, ws *websocket.Conn) {
	Parties[partyId].Clients[ws] = true
	Parties[partyId].Broadcast <- ToGameDto(
		game.ReplayHistory(Parties[partyId].History),
	)
}

func HandleParty(party Party) {
	for {
		// Grab the next message from the broadcast channel
		updatedParty := <-party.Broadcast
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
