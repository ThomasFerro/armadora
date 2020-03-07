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

// TODO: Error management: do not create error event anymore but return those errors

func ManageCommand(history []event.Event, msg Command) ([]event.Event, error) {
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
	case "PassTurn":
		return passTurn(history, msg)
	}
	return []event.Event{}, nil
}

func createGame(gameHistory []event.Event, msg Command) ([]event.Event, error) {
	return command.CreateGame()
}

func joinGame(history []event.Event, msg Command) ([]event.Event, error) {
	characterValue := getCharacter(msg.Payload["Character"])
	return command.JoinGame(history, command.JoinGamePayload{
		Nickname:  msg.Payload["Nickname"],
		Character: characterValue,
	})
}

func startTheGame(history []event.Event, msg Command) ([]event.Event, error) {
	return command.StartTheGame(history)
}

func putWarrior(history []event.Event, msg Command) ([]event.Event, error) {
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
	}), nil
}

func putPalisades(history []event.Event, msg Command) ([]event.Event, error) {
	// TODO: Pay the tech debt when managing authent
	currentGame := game.ReplayHistory(history)
	var palisades []palisade.Palisade
	err := json.Unmarshal([]byte(msg.Payload["Palisades"]), &palisades)
	if err != nil {
		return nil, err
	}
	return command.PutPalisades(history, command.PutPalisadesPayload{
		Player:    currentGame.CurrentPlayer(),
		Palisades: palisades,
	}), nil
}

func passTurn(history []event.Event, msg Command) ([]event.Event, error) {
	// TODO: Pay the tech debt when managing authent
	currentGame := game.ReplayHistory(history)
	return command.PassTurn(history, command.PassTurnPayload{
		Player: currentGame.CurrentPlayer(),
	}), nil
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
