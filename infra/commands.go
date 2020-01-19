package infra

import (
	"log"

	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
)

type Command struct {
	CommandType string            `json:"command_type"`
	Payload     map[string]string `json:"payload"`
}

func ManageCommand(history []event.Event, msg Command) []event.Event {
	switch msg.CommandType {
	case "CreateGame":
		return createGame(history, msg)
	case "JoinGame":
		return joinGame(history, msg)
	case "StartTheGame":
		return startTheGame(history, msg)
	}
	return []event.Event{}
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
