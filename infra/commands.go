package infra

import (
	"fmt"
	"log"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/google/uuid"
)

type Command struct {
	Id          string            `json:"id"`
	CommandType string            `json:"command_type"`
	Payload     map[string]string `json:"payload"`
}

func ManageCommand(msg Command) GameDto {
	fmt.Printf("Receiving command with id %v\n", msg.Id)
	gameId := msg.Id
	history := getHistory(gameId)
	newEvents := []event.Event{}
	switch msg.CommandType {
	case "CreateGame":
		gameId = generateNewId()
		newEvents = createGame(history, msg)
	case "JoinGame":
		newEvents = joinGame(history, msg)
	case "StartTheGame":
		newEvents = startTheGame(history, msg)
	}
	saveNewEvents(gameId, newEvents)
	return ToGameDto(
		game.ReplayHistory(
			append(
				history,
				newEvents...,
			),
		),
	)
}

func generateNewId() string {
	return uuid.New().String()
}

var gameHistory = map[string][]event.Event{}

func saveNewEvents(id string, newEvents []event.Event) {
	gameHistory[id] = append(gameHistory[id], newEvents...)
}

func getHistory(id string) []event.Event {
	return gameHistory[id]
}

func createGame(gameHistory []event.Event, msg Command) []event.Event {
	return []event.Event{
		command.CreateGame(),
	}
}

func joinGame(history []event.Event, msg Command) []event.Event {
	characterValue := getCharacter(msg.Payload["Character"])
	return command.JoinGame(history, command.JoinGamePayload{
		Nickname:  msg.Payload["Nickname"],
		Character: characterValue,
	})
}

func startTheGame(history []event.Event, msg Command) []event.Event {
	return command.StartTheGame(history)
}

func getCharacter(characterName string) character.Character {
	switch characterName {
	case "Orc":
		return character.Orc
	case "Goblin":
		return character.Goblin
	case "Elf":
		return character.Elf
	case "Mage":
		return character.Mage
	}
	log.Printf("Invalid character %v, falling back to the Orc\n", characterName)
	return character.Orc
}
