package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/exception"
)

/*
TODO:
- To determine: who can start the game ?
*/

func TestCreateAGame(t *testing.T) {
	gameCreatedEvents, err := command.CreateGame()

	if err != nil {
		t.Errorf("An error has occurred while creating the game: %v", err)
		return
	}

	if len(gameCreatedEvents) != 1 || len(gameCreatedEvents[0].EventMessage()) == 0 {
		t.Error("The game has not been created")
	}
}

func TestGameCreateInWaitingForPlayersState(t *testing.T) {
	gameCreatedEvent, err := command.CreateGame()

	if err != nil {
		t.Errorf("An error has occurred while creating the game: %v", err)
		return
	}

	newGame := game.ReplayHistory(gameCreatedEvent)

	if newGame.State() != game.WaitingForPlayers {
		t.Error("The new game is not in 'WaitingForPlayers' state")
	}
}

func TestJoinAGame(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
	}

	joinGameCommandPayload := command.JoinGamePayload{
		Nickname:  "README.md",
		Character: character.Goblin,
	}

	joinGameEvents, err := command.JoinGame(history, joinGameCommandPayload)

	if err != nil {
		t.Errorf("The player could not join the game: %v", err)
		return
	}

	history = append(history, joinGameEvents...)

	newGame := game.ReplayHistory(history)

	if len(newGame.Players()) != 1 {
		t.Error("The player did not join the game")
		return
	}

	newPlayer := newGame.Players()[0]

	if newPlayer.Nickname() != "README.md" || newPlayer.Character() != character.Goblin {
		t.Error("The player's information are not set correctly")
	}
}

func TestFourPlayersJoinAGame(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
	}

	joinGameCommandPayloads := []command.JoinGamePayload{
		command.JoinGamePayload{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		command.JoinGamePayload{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
		command.JoinGamePayload{
			Nickname:  "Kileek",
			Character: character.Orc,
		},
		command.JoinGamePayload{
			Nickname:  "LaNinjaBabané",
			Character: character.Mage,
		},
	}

	for _, nextCommand := range joinGameCommandPayloads {
		joinGameEvents, err := command.JoinGame(history, nextCommand)
		if err != nil {
			t.Errorf("The player could not join the game: %v", err)
			return
		}
		history = append(history, joinGameEvents...)
	}

	newGame := game.ReplayHistory(history)

	if len(newGame.Players()) != 4 {
		t.Error("There should be four players")
		return
	}

	for i, nextCommand := range joinGameCommandPayloads {
		player := newGame.Players()[i]
		if player.Nickname() != nextCommand.Nickname || player.Character() != nextCommand.Character {
			t.Error("One player's information does not match what was expected")
			return
		}
	}
}

func TestAFifthPlayerCannotJoinAGame(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
	}

	joinGameCommandPayloads := []command.JoinGamePayload{
		command.JoinGamePayload{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		command.JoinGamePayload{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
		command.JoinGamePayload{
			Nickname:  "Kileek",
			Character: character.Orc,
		},
		command.JoinGamePayload{
			Nickname:  "LaNinjaBabané",
			Character: character.Mage,
		},
	}

	for _, nextCommand := range joinGameCommandPayloads {
		joinGameEvents, err := command.JoinGame(history, nextCommand)
		if err != nil {
			t.Errorf("The player could not join the game: %v", err)
			return
		}
		history = append(history, joinGameEvents...)
	}

	unauthorizedPlayerTryingToJoin := command.JoinGamePayload{
		Nickname:  "TheFifthPlayer",
		Character: character.Mage,
	}

	_, err := command.JoinGame(history, unauthorizedPlayerTryingToJoin)
	if _, isOfRightExceptionType := err.(exception.GameAlreadyFull); !isOfRightExceptionType {
		t.Errorf("The player should not be able to join an already full game %v", err)
		return
	}
}

func TestANewPlayerCannotSelectACharacterAlreadyChosen(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		event.PlayerJoined{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
	}

	joinGameCommandPayload := command.JoinGamePayload{
		Nickname:  "Kileek",
		Character: character.Goblin,
	}

	_, err := command.JoinGame(history, joinGameCommandPayload)

	if _, isOfRightExceptionType := err.(exception.CharacterAlreadyChosen); !isOfRightExceptionType {
		t.Errorf("The player should not be able to select an already chosen character: %v", err)
		return
	}
}

func TestStartTheGame(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		event.PlayerJoined{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
	}

	startTheGameEvents, err := command.StartTheGame(history)

	if err != nil {
		t.Errorf("Could not start the game: %v", err)
		return
	}

	history = append(history, startTheGameEvents...)

	newGame := game.ReplayHistory(history)

	if newGame.State() != game.Started {
		t.Error("The new game is not in 'Started' state")
		return
	}

	if newGame.CurrentPlayer() != 0 {
		t.Error("The current player is not the first one")
	}
}

func TestCannotStartAGameWhenThereIsNoPlayer(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
	}

	_, err := command.StartTheGame(history)

	if _, isOfRightExceptionType := err.(exception.NotEnoughPlayers); !isOfRightExceptionType {
		t.Errorf("The game cannot start without players: %v", err)
		return
	}
}

func TestOnePlayerCannotStartAGameAlone(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
	}

	_, err := command.StartTheGame(history)

	if _, isOfRightExceptionType := err.(exception.NotEnoughPlayers); !isOfRightExceptionType {
		t.Errorf("The game cannot start with one player: %v", err)
		return
	}
}

func TestPlayersCannotJoinAGameOnceItStarted(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		event.PlayerJoined{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
		event.GameStarted{},
	}

	joinGameCommandPayload := command.JoinGamePayload{
		Nickname:  "Kileek",
		Character: character.Mage,
	}

	_, err := command.JoinGame(history, joinGameCommandPayload)

	if _, isOfRightExceptionType := err.(exception.GameAlreadyStarted); !isOfRightExceptionType {
		t.Errorf("The player should not be able to join the game %v", err)
		return
	}
}

func TestWarriorsShouldBeDistributedOnGameStart(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		event.PlayerJoined{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
		event.GameStarted{},
	}

	startTheGameEvents, err := command.StartTheGame(history)

	if err != nil {
		t.Errorf("Could not start the game: %v", err)
		return
	}

	history = append(history, startTheGameEvents...)
	warriorsDistributedEventFound := false

	for _, nextEvent := range history {
		if _, isOfRightEventType := nextEvent.(event.WarriorsDistributed); isOfRightEventType {
			warriorsDistributedEventFound = true
		}
	}

	if !warriorsDistributedEventFound {
		t.Error("The warriors were not distributed")
	}

	newGame := game.ReplayHistory(history)
	for _, player := range newGame.Players() {
		if player.Warriors() == nil {
			t.Error("A player has no warrior")
			return
		}
	}
}

func TestGoldStacksShouldBeDistributedOnGameStart(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		event.PlayerJoined{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
		event.GameStarted{},
	}

	startTheGameEvents, err := command.StartTheGame(history)

	if err != nil {
		t.Errorf("Could not start the game: %v", err)
		return
	}

	history = append(history, startTheGameEvents...)
	var goldStacksDistributed []int

	for _, nextEvent := range history {
		if rightEvent, isOfRightEventType := nextEvent.(event.GoldStacksDistributed); isOfRightEventType {
			goldStacksDistributed = rightEvent.GoldStacks
		}
	}

	if goldStacksDistributed == nil {
		t.Error("The gold stacks were not distributed")
		return
	}

	newGame := game.ReplayHistory(history)
	boardGoldStacks := newGame.Board().GoldStacks()
	if len(boardGoldStacks) != len(goldStacksDistributed) {
		t.Error("The game's board did not receive all of the gold to distribute")
		return
	}
	for index, stack := range boardGoldStacks {
		if goldStacksDistributed[index] != stack {
			t.Error("The distributed gold stack does not match with the board")
			return
		}
	}
}

func TestPalisadesShouldBeDistributedOnGameStart(t *testing.T) {
	history := []event.Event{
		event.GameCreated{},
		event.PlayerJoined{
			Nickname:  "README.md",
			Character: character.Goblin,
		},
		event.PlayerJoined{
			Nickname:  "Javadoc",
			Character: character.Elf,
		},
		event.GameStarted{},
	}

	startTheGameEvents, err := command.StartTheGame(history)

	if err != nil {
		t.Errorf("Could not start the game: %v", err)
		return
	}

	history = append(history, startTheGameEvents...)

	var palisadesDistributedEvent event.PalisadesDistributed

	for _, nextEvent := range history {
		if typedEvent, isOfRightEventType := nextEvent.(event.PalisadesDistributed); isOfRightEventType {
			palisadesDistributedEvent = typedEvent
		}
	}

	if palisadesDistributedEvent.Count != 35 {
		t.Errorf("The palisades are not distributed, event: %v", palisadesDistributedEvent)
		return
	}

	newGame := game.ReplayHistory(history)

	if newGame.Board().PalisadesLeft() != 35 {
		t.Errorf("The distributed palisades does not match with the rules, should have 35 but got %v", newGame.Board().PalisadesLeft())
	}
}
