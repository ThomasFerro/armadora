package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ThomasFerro/armadora/infra"
	"github.com/gorilla/websocket"
)

var parties = map[infra.PartyId]infra.Party{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/games", handleGameCreation)

	http.HandleFunc("/parties", handleGetParties)

	http.HandleFunc("/parties/", handleConnectionsToPartyWs)

	port := ":80"
	log.Printf("Serving Armadora on port: %v\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Cannot start the server: %v\n", err)
	}
}

func handleGameCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	newParty := infra.CreateParty()
	log.Printf("Creating a new party: %v\n", newParty)
	w.Write(
		[]byte(
			fmt.Sprintf("{\"id\": \"%v\"}", newParty),
		),
	)
}

func handleConnectionsToPartyWs(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.String(), "/")
	partyId := infra.PartyId(urlParts[len(urlParts)-1])
	log.Printf("Connection to party %v\n", partyId)

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error while receiving initial WS request: %v\n", err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	infra.AddClientToParty(partyId, ws)

	for {
		var msg infra.Command
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error while reading party message: %v", err)
			infra.RemoveClientFromParty(partyId, ws)
			break
		}
		infra.ReceiveCommand(partyId, msg)
	}
}

func handleGetParties(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	log.Printf("Returning the %v parties\n", len(infra.Parties))
	partiesId := make([]infra.PartyId, 0, len(infra.Parties))
	for partyId := range infra.Parties {
		partiesId = append(partiesId, partyId)
	}
	partiesIdJson, err := json.Marshal(partiesId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.Write(
		partiesIdJson,
	)
}
