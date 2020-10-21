package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ThomasFerro/armadora/infra"
	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/party"
)

var allowedOrigin string
var armadoraService infra.ArmadoraService

func StartApi() {
	armadoraService = infra.NewArmadoraService()

	http.HandleFunc("/parties", handlePartiesRequest)

	http.HandleFunc("/parties/", handlePartyRequest)

	allowedOrigin = config.GetConfiguration("ALLOWED_ORIGIN")

	port := config.GetConfiguration("PORT")
	log.Printf("Serving Armadora on port: %v\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Cannot start the server: %v\n", err)
	}
}

func handlePartiesRequest(w http.ResponseWriter, r *http.Request) {
	manageCors(&w)

	switch r.Method {
	case "GET":
		handleGetParties(w, r)
	case "POST":
		handlePartyCreation(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
	}
}

func handleGetParties(w http.ResponseWriter, r *http.Request) {
	manageCors(&w)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	parties, err := armadoraService.GetVisibleParties()
	if err != nil {
		manageError(&w, err)
		return
	}

	log.Printf("Returning the %v parties\n", len(parties))
	partiesIdJson, err := json.Marshal(parties)
	if err != nil {
		manageError(&w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(
		partiesIdJson,
	)
}

func handlePartyCreation(w http.ResponseWriter, r *http.Request) {
	manageCors(&w)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	newParty, err := armadoraService.CreateParty()

	if err != nil {
		log.Printf("Cannot create a new party: %v\n", err)
		manageError(&w, err)
		return
	}

	log.Printf("Creating a new party: %v\n", newParty)
	w.Header().Add("Content-Type", "application/json")
	w.Write(
		[]byte(
			fmt.Sprintf("{\"id\": \"%v\"}", newParty),
		),
	)
}

func handlePartyRequest(w http.ResponseWriter, r *http.Request) {
	manageCors(&w)
	urlParts := strings.Split(r.URL.String(), "/")
	partyName := party.PartyName(urlParts[len(urlParts)-1])

	switch r.Method {
	case "GET":
		handleGetPartyState(partyName, w, r)
	case "POST":
		handlePostPartyCommand(partyName, w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
	}
}

func handleGetPartyState(partyName party.PartyName, w http.ResponseWriter, r *http.Request) {
	party, err := armadoraService.GetPartyGameState(partyName)
	if err != nil {
		log.Printf("Cannot get the party %v: %v\n", partyName, err)
		manageError(&w, err)
		return
	}
	partyJson, err := json.Marshal(
		party,
	)
	if err != nil {
		manageError(&w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(partyJson)
}

func handlePostPartyCommand(partyName party.PartyName, w http.ResponseWriter, r *http.Request) {
	log.Printf("Command received for party %v: %v\n", partyName, r.Body)
	decoder := json.NewDecoder(r.Body)
	// TODO: Pay the tech debt when managing authent
	// Do not trust the user with the provided player_id, but retrieve it based on the authentication token
	var command infra.Command
	err := decoder.Decode(&command)
	if err != nil {
		manageError(&w, err)
	}

	err = armadoraService.ReceiveCommand(partyName, command)

	if err != nil {
		manageError(&w, err)
		return
	}
	handleGetPartyState(partyName, w, r)
}

func manageError(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write([]byte(err.Error()))
}

func manageCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", allowedOrigin)
}
