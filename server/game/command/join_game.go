package command

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/character"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/exception"
)

// JoinGamePayload Containing every information about the player who wants to join
type JoinGamePayload struct {
	Nickname  string
	Character character.Character
}

// JoinGame Join a game
func JoinGame(history []event.Event, joinGamePayload JoinGamePayload) ([]event.Event, error) {
	gameToJoin := game.ReplayHistory(history)

	if gameToJoin.State() != game.WaitingForPlayers {
		return nil, exception.GameAlreadyStarted{}
	}

	if len(gameToJoin.Players()) == 4 {
		return nil, exception.GameAlreadyFull{}
	}

	for _, player := range gameToJoin.Players() {
		if player.Character() == joinGamePayload.Character {
			return nil, exception.CharacterAlreadyChosen{
				Character: joinGamePayload.Character,
			}
		}
	}

	return []event.Event{
		event.PlayerJoined{
			Nickname:  joinGamePayload.Nickname,
			Character: joinGamePayload.Character,
		},
	}, nil
}
