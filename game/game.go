package game

import (
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/event"
)

// Game Instance of an Aramdora game
type Game interface {
	State() State
	Players() []Player
	CurrentPlayer() int
	Board() board.Board
	ApplyGameCreated(event event.GameCreated) Game
	ApplyPlayerJoined(event event.PlayerJoined) Game
	ApplyWarriorsDistributed(event event.WarriorsDistributed) Game
	ApplyGameStarted(event event.GameStarted) Game
	ApplyGoldStacksDistributed(event event.GoldStacksDistributed) Game
	ApplyNextPlayer(event event.NextPlayer) Game
}

type game struct {
	state         State
	players       []Player
	currentPlayer int
	board         board.Board
}

// State The game's current state
func (g game) State() State {
	return g.state
}

// Player The game's Players
func (g game) Players() []Player {
	return g.players
}

// CurrentPlayer The game's current player
func (g game) CurrentPlayer() int {
	return g.currentPlayer
}

// Board The game's Board
func (g game) Board() board.Board {
	return g.board
}

func (g game) ApplyGameCreated(event event.GameCreated) Game {
	g.state = WaitingForPlayers
	return g
}

func (g game) ApplyPlayerJoined(event event.PlayerJoined) Game {
	g.players = append(g.players, NewPlayer(event.Nickname, event.Character))
	return g
}

func (g game) ApplyWarriorsDistributed(event event.WarriorsDistributed) Game {
	players := []Player{}
	for _, player := range g.Players() {
		players = append(players, player.SetWarriors(event.WarriorsDistributed))
	}
	g.players = players
	return g
}

func (g game) ApplyGameStarted(event event.GameStarted) Game {
	g.state = Started
	return g
}

func (g game) ApplyGoldStacksDistributed(event event.GoldStacksDistributed) Game {
	g.board = board.NewBoard(event.GoldStacks)
	return g
}

func (g game) ApplyNextPlayer(event event.NextPlayer) Game {
	if g.currentPlayer == len(g.players)-1 {
		g.currentPlayer = 0
	} else {
		g.currentPlayer++
	}
	return g
}

// ReplayHistory Replay the provided history to retrieve the game state
func ReplayHistory(history []event.Event) Game {
	var returnedGame Game
	returnedGame = game{}
	for _, nextEvent := range history {
		switch nextEvent.(type) {
		case event.GameCreated:
			gameCreatedEvent, _ := nextEvent.(event.GameCreated)
			returnedGame = returnedGame.ApplyGameCreated(gameCreatedEvent)
		case event.PlayerJoined:
			playerJoinedEvent, _ := nextEvent.(event.PlayerJoined)
			returnedGame = returnedGame.ApplyPlayerJoined(playerJoinedEvent)
		case event.WarriorsDistributed:
			warriorsDistributedEvent, _ := nextEvent.(event.WarriorsDistributed)
			returnedGame = returnedGame.ApplyWarriorsDistributed(warriorsDistributedEvent)
		case event.GameStarted:
			gameStartedEvent, _ := nextEvent.(event.GameStarted)
			returnedGame = returnedGame.ApplyGameStarted(gameStartedEvent)
		case event.GoldStacksDistributed:
			goldStacksDistributedEvent, _ := nextEvent.(event.GoldStacksDistributed)
			returnedGame = returnedGame.ApplyGoldStacksDistributed(goldStacksDistributedEvent)
		case event.NextPlayer:
			nextPlayerEvent, _ := nextEvent.(event.NextPlayer)
			returnedGame = returnedGame.ApplyNextPlayer(nextPlayerEvent)
		}
	}
	return returnedGame
}
