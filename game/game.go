package game

import (
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/palisade"
	"github.com/ThomasFerro/armadora/game/warrior"
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
	ApplyPalisadesDistributed(event event.PalisadesDistributed) Game
	ApplyNextPlayer(event event.NextPlayer) Game
	ApplyWarriorPut(event event.WarriorPut) Game
	ApplyPalisadePut(event event.PalisadePut) Game
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

func (g game) ApplyPalisadesDistributed(event event.PalisadesDistributed) Game {
	g.board = g.Board().SetPalisadesLeft(event.Count)
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

func (g game) ApplyWarriorPut(event event.WarriorPut) Game {
	g.board = g.board.PutWarriorInCell(event.Position, event.Player, event.Strength)
	currentPlayer := g.players[g.currentPlayer]
	g.players[g.currentPlayer] = currentPlayer.SetWarriors(warrior.RemoveUsedWarrior(g.players[g.currentPlayer].Warriors(), event.Strength))
	return g
}

func (g game) ApplyPalisadePut(event event.PalisadePut) Game {
	g.board = g.Board().PutPalisade(palisade.Palisade{
		X1: event.X1,
		Y1: event.Y1,
		X2: event.X2,
		Y2: event.Y2,
	})
	return g
}

// ReplayHistory Replay the provided history to retrieve the game state
func ReplayHistory(history []event.Event) Game {
	var returnedGame Game
	returnedGame = game{}
	for _, nextEvent := range history {
		switch typedEvent := nextEvent.(type) {
		case event.GameCreated:
			returnedGame = returnedGame.ApplyGameCreated(typedEvent)
		case event.PlayerJoined:
			returnedGame = returnedGame.ApplyPlayerJoined(typedEvent)
		case event.WarriorsDistributed:
			returnedGame = returnedGame.ApplyWarriorsDistributed(typedEvent)
		case event.GameStarted:
			returnedGame = returnedGame.ApplyGameStarted(typedEvent)
		case event.GoldStacksDistributed:
			returnedGame = returnedGame.ApplyGoldStacksDistributed(typedEvent)
		case event.PalisadesDistributed:
			returnedGame = returnedGame.ApplyPalisadesDistributed(typedEvent)
		case event.NextPlayer:
			returnedGame = returnedGame.ApplyNextPlayer(typedEvent)
		case event.WarriorPut:
			returnedGame = returnedGame.ApplyWarriorPut(typedEvent)
		case event.PalisadePut:
			returnedGame = returnedGame.ApplyPalisadePut(typedEvent)
		}
	}
	return returnedGame
}
