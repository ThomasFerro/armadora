package command

import (
	"math"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/palisade"
)

// PutPalisadesPayload Data about the palisades to put on the board
type PutPalisadesPayload struct {
	Player    int
	Palisades []palisade.Palisade
}

func validPalisade(palisadeToCheck palisade.Palisade) bool {
	diff := 0.

	diff += math.Abs(float64(palisadeToCheck.X1 - palisadeToCheck.X2))
	diff += math.Abs(float64(palisadeToCheck.Y1 - palisadeToCheck.Y2))

	return diff == 1.
}

func equals(firstPalisade, secondPalisade palisade.Palisade) bool {
	return firstPalisade.X1 == secondPalisade.X1 &&
		firstPalisade.Y1 == secondPalisade.Y1 &&
		firstPalisade.X2 == secondPalisade.X2 &&
		firstPalisade.Y2 == secondPalisade.Y2
}

func vacantBorder(currentGame game.Game, palisadeToCheck palisade.Palisade) bool {
	for _, gridPalisade := range currentGame.Board().Palisades() {
		if equals(palisadeToCheck, gridPalisade) {
			return false
		}
	}
	return true
}

func getDistinctPalisades(palisades []palisade.Palisade) []palisade.Palisade {
	distinct := []palisade.Palisade{}

	for _, nextPalisade := range palisades {
		alreadyInSlice := false

		for _, nextDistinctPalisade := range distinct {
			if equals(nextDistinctPalisade, nextPalisade) {
				alreadyInSlice = true
				break
			}
		}

		if !alreadyInSlice {
			distinct = append(distinct, nextPalisade)
		}
	}

	return distinct
}

func breaksGridValidity(currentGame game.Game, palisadeToCheck palisade.Palisade) bool {
	_, err := board.FindTerritories(currentGame.Board().PutPalisade(palisadeToCheck))
	return err != nil
}

// PutPalisades Put palisades on the board
func PutPalisades(history []event.Event, payload PutPalisadesPayload) []event.Event {
	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != payload.Player {
		return []event.Event{
			event.NotThePlayerTurn{
				PlayerWhoTriedToPlay: payload.Player,
			},
		}
	}

	events := []event.Event{}

	distinctPalisades := getDistinctPalisades(payload.Palisades)

	if len(distinctPalisades) > currentGame.Board().PalisadesLeft() {
		return []event.Event{
			event.NoMorePalisadeLeft{},
		}
	}

	for _, palisade := range distinctPalisades {
		if !validPalisade(palisade) {
			return []event.Event{
				event.InvalidPalisadePosition{
					Player: payload.Player,
					X1:     palisade.X1,
					Y1:     palisade.Y1,
					X2:     palisade.X2,
					Y2:     palisade.Y2,
				},
			}
		}

		if !vacantBorder(currentGame, palisade) {
			return []event.Event{
				event.BorderAlreadyTaken{
					Player: payload.Player,
					X1:     palisade.X1,
					Y1:     palisade.Y1,
					X2:     palisade.X2,
					Y2:     palisade.Y2,
				},
			}
		}

		if breaksGridValidity(currentGame, palisade) {
			return []event.Event{
				event.InvalidPalisadePosition{
					Player: payload.Player,
					X1:     palisade.X1,
					Y1:     palisade.Y1,
					X2:     palisade.X2,
					Y2:     palisade.Y2,
				},
			}
		}

		events = append(events, event.PalisadePut{
			Player: payload.Player,
			X1:     palisade.X1,
			Y1:     palisade.Y1,
			X2:     palisade.X2,
			Y2:     palisade.Y2,
		})
		currentGame = game.ReplayHistory(append(history, events...))
	}

	return append(events, event.NextPlayer{})
}
