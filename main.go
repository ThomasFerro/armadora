package main

import (
	"log"
	"net/http"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/infra"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan infra.GameDto)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// TODO: Manage multiple games
// TODO: Manage reconnection
var gameHistory = []event.Event{
	command.CreateGame(),
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Cannot start the server: %v\n", err)
	}
}

func currentGameState() infra.GameDto {
	return infra.ToGameDto(
		game.ReplayHistory(gameHistory),
	)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error while receiving initial WS request: %v\n", err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	// Send the current state of the game
	sendGameToClient(ws, currentGameState())

	for {
		var msg infra.Command
		// Read in a new message as JSON and map it to a Command object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error while reading the command: %v\n", err)
			delete(clients, ws)
			break
		}

		gameHistory = append(
			gameHistory,
			infra.ManageCommand(gameHistory, msg)...,
		)
		// Send the new game state to the broadcast channel
		gameDto := infra.ToGameDto(
			game.ReplayHistory(gameHistory),
		)
		log.Printf("Sending new game state: %v\n", gameDto)
		broadcast <- gameDto
	}
}

func sendGameToClient(client *websocket.Conn, gameDto infra.GameDto) {
	err := client.WriteJSON(gameDto)
	if err != nil {
		log.Printf("Error while writing to the WS: %v\n", err)
		client.Close()
		delete(clients, client)
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		gameDto := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			sendGameToClient(client, gameDto)
		}
	}
}
