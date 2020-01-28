package infra

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/palisade"
)

type Command struct {
	CommandType string            `json:"command_type"`
	Payload     map[string]string `json:"payload"`
}

// TODO: Error management

func ManageCommand(history []event.Event, msg Command) []event.Event {
	switch msg.CommandType {
	case "CreateGame":
		return createGame(history, msg)
	case "JoinGame":
		return joinGame(history, msg)
	case "StartTheGame":
		return startTheGame(history, msg)
	case "PutWarrior":
		return putWarrior(history, msg)
	case "PutPalisades":
		return putPalisades(history, msg)
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

func putWarrior(history []event.Event, msg Command) []event.Event {
	// TODO: Pay the tech debt when managing authent
	currentGame := game.ReplayHistory(history)
	warrior, _ := strconv.Atoi(msg.Payload["Warrior"])
	x, _ := strconv.Atoi(msg.Payload["X"])
	y, _ := strconv.Atoi(msg.Payload["Y"])
	return command.PutWarrior(history, command.PutWarriorPayload{
		Player:  currentGame.CurrentPlayer(),
		Warrior: warrior,
		Position: board.Position{
			X: x,
			Y: y,
		},
	})
}

func putPalisades(history []event.Event, msg Command) []event.Event {
	// TODO: Pay the tech debt when managing authent
	currentGame := game.ReplayHistory(history)
	// TODO: Error management
	var palisades []palisade.Palisade
	json.Unmarshal([]byte(msg.Payload["Palisades"]), &palisades)
	return command.PutPalisades(history, command.PutPalisadesPayload{
		Player:    currentGame.CurrentPlayer(),
		Palisades: palisades,
	})
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
