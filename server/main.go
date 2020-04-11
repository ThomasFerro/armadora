package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ThomasFerro/armadora/infra"
)

var allowedOrigin string

func main() {
	http.HandleFunc("/games", handleGameCreation)

	http.HandleFunc("/parties", handleGetParties)

	// TODO: Manage get on specific party
	http.HandleFunc("/parties/", handlePartyRequest)

	allowedOrigin = os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:8081"
		log.Printf("No allowed origin provided in ALLOWED_ORIGIN, falling back to %v", allowedOrigin)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("No port provided in PORT, falling back to %v", port)
	}
	log.Printf("Serving Armadora on port: %v\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Cannot start the server: %v\n", err)
	}
}

func handleGameCreation(w http.ResponseWriter, r *http.Request) {
	manageCors(&w)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	newParty, err := infra.CreateParty()

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
	partyId := infra.PartyId(urlParts[len(urlParts)-1])

	switch r.Method {
	case "GET":
		handleGetPartyState(partyId, w, r)
	case "POST":
		handlePostPartyCommand(partyId, w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
	}
}

func handleGetPartyState(partyId infra.PartyId, w http.ResponseWriter, r *http.Request) {
	party, err := infra.GetParty(partyId)
	if err != nil {
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

func handlePostPartyCommand(partyId infra.PartyId, w http.ResponseWriter, r *http.Request) {
	log.Printf("Command received for party %v: %v\n", partyId, r.Body)
	decoder := json.NewDecoder(r.Body)
	var command infra.Command
	err := decoder.Decode(&command)
	if err != nil {
		manageError(&w, err)
	}

	err = infra.ReceiveCommand(partyId, command)

	if err != nil {
		manageError(&w, err)
		return
	}
	handleGetPartyState(partyId, w, r)
}

func handleGetParties(w http.ResponseWriter, r *http.Request) {
	manageCors(&w)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	// TODO: Don't check infra.Parties but call a method that return every parties (stateless)

	log.Printf("Returning the %v parties\n", len(infra.Parties))
	partiesIdJson, err := json.Marshal(infra.Parties)
	if err != nil {
		manageError(&w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(
		partiesIdJson,
	)
}

func manageError(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write([]byte(err.Error()))
}

func manageCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", allowedOrigin)
}
