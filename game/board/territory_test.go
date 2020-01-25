package board_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/gold"
	"github.com/ThomasFerro/armadora/game/palisade"
)

/*
TODO:
- Invalid territories
*/

func TestValidGridWithOneTerritory(t *testing.T) {
	testedBoard := board.NewBoard(gold.GoldStacks)
	testedBoard = testedBoard.PutWarriorInCell(
		board.Position{
			X: 0,
			Y: 0,
		},
		0,
		5,
	)
	testedBoard = testedBoard.PutWarriorInCell(
		board.Position{
			X: 1,
			Y: 0,
		},
		1,
		2,
	)
	totalGold := 0
	for _, stack := range gold.GoldStacks {
		totalGold += stack
	}

	territories, err := board.FindTerritories(testedBoard)
	if err != nil {
		t.Errorf("The board is invalid: %v", err)
		return
	}

	if len(territories) != 1 {
		t.Errorf("Expecting to get a single territory, got this instead: %v", territories)
		return
	}

	if territories[0].Gold() != totalGold {
		t.Errorf("Expecting the single territory to have all of the board's gold (%v), got this instead: %v", totalGold, territories[0].Gold())
		return
	}

	if territories[0].WinningPlayers()[0] != 0 {
		t.Errorf("Expecting the single territory to have player 0 winning, got this instead: %v", territories[0].WinningPlayers())
	}
}

func TestMultipleWinnersWithinOneTerritory(t *testing.T) {
	testedBoard := board.NewBoard(gold.GoldStacks)
	testedBoard = testedBoard.PutWarriorInCell(
		board.Position{
			X: 0,
			Y: 0,
		},
		0,
		5,
	)
	testedBoard = testedBoard.PutWarriorInCell(
		board.Position{
			X: 1,
			Y: 0,
		},
		1,
		5,
	)
	totalGold := 0
	for _, stack := range gold.GoldStacks {
		totalGold += stack
	}

	territories, _ := board.FindTerritories(testedBoard)

	if len(territories[0].WinningPlayers()) != 2 {
		t.Errorf("Expecting the territory to have the two players winning, got this instead: %v", territories[0].WinningPlayers())
	}
}

func TestValidFourCellsTerritory(t *testing.T) {
	testedBoard := board.NewBoard(gold.GoldStacks)
	testedBoard = testedBoard.PutWarriorInCell(
		board.Position{
			X: 0,
			Y: 0,
		},
		0,
		5,
	)
	testedBoard = testedBoard.PutWarriorInCell(
		board.Position{
			X: 5,
			Y: 0,
		},
		1,
		5,
	)
	palisadesToPut := []palisade.Palisade{
		palisade.Palisade{
			X1: 3,
			Y1: 0,
			X2: 3,
			Y2: 1,
		},
		palisade.Palisade{
			X1: 4,
			Y1: 0,
			X2: 4,
			Y2: 1,
		},
		palisade.Palisade{
			X1: 2,
			Y1: 1,
			X2: 3,
			Y2: 1,
		},
		palisade.Palisade{
			X1: 2,
			Y1: 2,
			X2: 3,
			Y2: 2,
		},
		palisade.Palisade{
			X1: 3,
			Y1: 2,
			X2: 3,
			Y2: 3,
		},
		palisade.Palisade{
			X1: 4,
			Y1: 2,
			X2: 4,
			Y2: 3,
		},
		palisade.Palisade{
			X1: 4,
			Y1: 1,
			X2: 5,
			Y2: 1,
		},
		palisade.Palisade{
			X1: 4,
			Y1: 2,
			X2: 5,
			Y2: 2,
		},
	}

	for _, palpalisadeToPut := range palisadesToPut {
		testedBoard = testedBoard.PutPalisade(palpalisadeToPut)
	}

	territories, _ := board.FindTerritories(testedBoard)

	if len(territories) != 2 {
		t.Errorf("Expecting to get two territories, got this instead: %v", territories)
		return
	}
	// TODO: Vérifier que le premier territoire est gagné par le joueur 1 et que le second est gagné par le joueur 2 ?
}
