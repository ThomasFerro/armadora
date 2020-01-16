package game_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/command"
	"github.com/ThomasFerro/armadora/game/event"
)

/*
TODO:
- To determine: who can start the game ?
- Distribute the gold
*/

func TestCreateAGame(t *testing.T) {
	gameCreatedEvent := command.CreateGame()

	if len(gameCreatedEvent.EventMessage()) == 0 {
		t.Error("The game has not been created")
	}
}

func TestGameCreateInWaitingForPlayersState(t *testing.T) {
	gameCreatedEvent := command.CreateGame()

	newGame := game.ReplayHistory([]event.Event{
		gameCreatedEvent,
	})

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

	history = append(history, command.JoinGame(history, joinGameCommandPayload)...)

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
		history = append(history, command.JoinGame(history, nextCommand)...)
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
		history = append(history, command.JoinGame(history, nextCommand)...)
	}

	unauthorizedPlayerTryingToJoin := command.JoinGamePayload{
		Nickname:  "TheFifthPlayer",
		Character: character.Mage,
	}

	history = append(history, command.JoinGame(history, unauthorizedPlayerTryingToJoin)...)
	lastEvent := history[len(history)-1]

	newGame := game.ReplayHistory(history)

	if _, isOfRightEventType := lastEvent.(event.GameAlreadyFull); !isOfRightEventType {
		t.Error("The error event was not sent")
		return
	}

	if len(newGame.Players()) != 4 {
		t.Error("The fifth player was added to the game")
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

	history = append(history, command.JoinGame(history, joinGameCommandPayload)...)
	lastEvent := history[len(history)-1]
	newGame := game.ReplayHistory(history)

	if _, isOfRightEventType := lastEvent.(event.CharacterAlreadyChosen); !isOfRightEventType {
		t.Error("The error event was not sent")
		return
	}

	if len(newGame.Players()) != 2 {
		t.Error("The player with already chosen character was added to the game")
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

	history = append(history, command.StartTheGame(history)...)

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

	history = append(history, command.StartTheGame(history)...)
	lastEvent := history[len(history)-1]

	newGame := game.ReplayHistory(history)

	if _, isOfRightEventType := lastEvent.(event.NotEnoughPlayers); !isOfRightEventType {
		t.Error("The error event was not sent")
		return
	}

	if newGame.State() != game.WaitingForPlayers {
		t.Error("The game's state has change, it is not waiting for players anymore")
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

	history = append(history, command.StartTheGame(history)...)
	lastEvent := history[len(history)-1]

	newGame := game.ReplayHistory(history)

	if _, isOfRightEventType := lastEvent.(event.NotEnoughPlayers); !isOfRightEventType {
		t.Error("The error event was not sent")
		return
	}

	if newGame.State() != game.WaitingForPlayers {
		t.Error("The game's state has change, it is not waiting for players anymore")
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

	history = append(history, command.JoinGame(history, joinGameCommandPayload)...)
	lastEvent := history[len(history)-1]

	newGame := game.ReplayHistory(history)

	if _, isOfRightEventType := lastEvent.(event.GameAlreadyStarted); !isOfRightEventType {
		t.Error("The error event was not sent")
		return
	}

	if len(newGame.Players()) != 2 {
		t.Error("The player was added after the game started")
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

	history = append(history, command.StartTheGame(history)...)
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
