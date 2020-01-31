package board

import "fmt"

import "github.com/ThomasFerro/armadora/game/palisade"
import "github.com/ThomasFerro/armadora/game/board/cell"

// BoardInvalid The board is invalid
type BoardInvalid struct {
	reason string
}

func (b BoardInvalid) Error() string {
	return fmt.Sprintf("The board is invalid for the following reason: %v", b.reason)
}

type territoryId int

// Territory A territory on the board
type Territory interface {
	Gold() int
	WinningPlayers() []int
}

type territory struct {
	gold           int
	winningPlayers []int
}

func (t territory) Gold() int {
	return t.gold
}

func (t territory) WinningPlayers() []int {
	return t.winningPlayers
}

func (t territory) String() string {
	return fmt.Sprintf("Territory containing %v gold with winning players %v", t.Gold(), t.WinningPlayers())
}

func findTerritoryWinnerAndGolds(cells []cell.Cell) Territory {
	playersStrength := map[int]int{}
	greatestStrength := 0
	totalGold := 0
	for _, nextCell := range cells {
		switch typedCell := nextCell.(type) {
		case cell.Gold:
			totalGold += typedCell.Stack()
		case cell.Warrior:
			playerStrength, isPresent := playersStrength[typedCell.Player()]
			if !isPresent {
				playersStrength[typedCell.Player()] = 0
			}
			playerStrengthInTerritory := playerStrength + typedCell.Strength()
			playersStrength[typedCell.Player()] = playerStrength + typedCell.Strength()
			if playerStrengthInTerritory > greatestStrength {
				greatestStrength = playerStrengthInTerritory
			}
		}
	}

	winningPlayers := []int{}
	for player, strength := range playersStrength {
		if strength == greatestStrength {
			winningPlayers = append(winningPlayers, player)
		}
	}

	return NewTerritory(totalGold, winningPlayers)
}

type cellWithTerritoryId struct {
	territoryId territoryId
	cell        cell.Cell
	palisades   []palisade.Palisade
}

func (c cellWithTerritoryId) String() string {
	return fmt.Sprint(c.territoryId)
}

func manageGridCell(grid [][]cellWithTerritoryId, x, y int, nextTerritoryId int) {
	if grid[y][x].territoryId != 0 {
		return
	}
	grid[y][x].territoryId = territoryId(nextTerritoryId)
	// Left neighbor with no palisade
	if x > 0 {
		leftNeighbor := grid[y][x-1]
		palisadeFound := false
		for _, nextPalisade := range leftNeighbor.palisades {
			if (nextPalisade.X1 == x && nextPalisade.Y1 == y && nextPalisade.X2 == x-1 && nextPalisade.Y2 == y) ||
				(nextPalisade.X1 == x-1 && nextPalisade.Y1 == y && nextPalisade.X2 == x && nextPalisade.Y2 == y) {
				palisadeFound = true
				break
			}
		}
		if !palisadeFound {
			manageGridCell(grid, x-1, y, nextTerritoryId)
		}
	}
	// Right neighbor with no palisade
	if x < len(grid[y])-1 {
		rightNeighbor := grid[y][x+1]
		palisadeFound := false
		for _, nextPalisade := range rightNeighbor.palisades {
			if (nextPalisade.X1 == x && nextPalisade.Y1 == y && nextPalisade.X2 == x+1 && nextPalisade.Y2 == y) ||
				(nextPalisade.X1 == x+1 && nextPalisade.Y1 == y && nextPalisade.X2 == x && nextPalisade.Y2 == y) {
				palisadeFound = true
				break
			}
		}
		if !palisadeFound {
			manageGridCell(grid, x+1, y, nextTerritoryId)
		}
	}
	// Top neighbor with no palisade
	if y > 0 {
		topNeighbor := grid[y-1][x]
		palisadeFound := false
		for _, nextPalisade := range topNeighbor.palisades {
			if (nextPalisade.X1 == x && nextPalisade.Y1 == y && nextPalisade.X2 == x && nextPalisade.Y2 == y-1) ||
				(nextPalisade.X1 == x && nextPalisade.Y1 == y-1 && nextPalisade.X2 == x && nextPalisade.Y2 == y) {
				palisadeFound = true
				break
			}
		}
		if !palisadeFound {
			manageGridCell(grid, x, y-1, nextTerritoryId)
		}
	}
	// Bottom neighbor with no palisade
	if y < len(grid)-1 {
		bottomNeighbor := grid[y+1][x]
		palisadeFound := false
		for _, nextPalisade := range bottomNeighbor.palisades {
			if (nextPalisade.X1 == x && nextPalisade.Y1 == y && nextPalisade.X2 == x && nextPalisade.Y2 == y+1) ||
				(nextPalisade.X1 == x && nextPalisade.Y1 == y+1 && nextPalisade.X2 == x && nextPalisade.Y2 == y) {
				palisadeFound = true
				break
			}
		}
		if !palisadeFound {
			manageGridCell(grid, x, y+1, nextTerritoryId)
		}
	}
}

func getCellsWithTerritoryId(grid [][]cellWithTerritoryId) []cellWithTerritoryId {
	nextTerritoryId := 1
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			manageGridCell(grid, x, y, nextTerritoryId)
			nextTerritoryId++
		}
	}

	cellsWithTerritoryId := []cellWithTerritoryId{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			cellsWithTerritoryId = append(cellsWithTerritoryId, grid[y][x])
		}
	}
	return cellsWithTerritoryId
}

func findInterestingPalisades(boardPalisades []palisade.Palisade, x, y int) []palisade.Palisade {
	interestingPalisades := []palisade.Palisade{}

	for _, boardPalisade := range boardPalisades {
		if (boardPalisade.X1 == x && boardPalisade.Y1 == y) || (boardPalisade.X2 == x && boardPalisade.Y2 == y) {
			interestingPalisades = append(interestingPalisades, boardPalisade)
		}
	}

	return interestingPalisades
}

func initCellsWithTerritoryId(board Board) [][]cellWithTerritoryId {
	cellsWithTerritoryId := [][]cellWithTerritoryId{}
	for y := 0; y < board.Height(); y++ {
		cellsWithTerritoryId = append(cellsWithTerritoryId, []cellWithTerritoryId{})
		for x := 0; x < board.Width(); x++ {
			cellsWithTerritoryId[y] = append(
				cellsWithTerritoryId[y],
				cellWithTerritoryId{
					territoryId: 0,
					cell: board.Cell(Position{
						X: x,
						Y: y,
					}),
					palisades: findInterestingPalisades(board.Palisades(), x, y),
				},
			)
		}
	}
	return cellsWithTerritoryId
}

func extractTerritories(boardToCompute Board) (map[territoryId][]cell.Cell, error) {
	cells := getCellsWithTerritoryId(
		initCellsWithTerritoryId(boardToCompute),
	)

	extractedTerritories := map[territoryId][]cell.Cell{}

	for _, nextCell := range cells {
		if _, alreadyInTheMap := extractedTerritories[nextCell.territoryId]; !alreadyInTheMap {
			extractedTerritories[nextCell.territoryId] = []cell.Cell{}
		}
		extractedTerritories[nextCell.territoryId] = append(extractedTerritories[nextCell.territoryId], nextCell.cell)
	}

	for _, territory := range extractedTerritories {
		if len(territory) < 4 {
			return nil, BoardInvalid{
				reason: "At least one territory is smaller than four cells",
			}
		}
	}

	return extractedTerritories, nil
}

func NewTerritory(gold int, winningPlayers []int) Territory {
	return territory{
		gold:           gold,
		winningPlayers: winningPlayers,
	}
}

// FindTerritories Find the territories in the provided board
func FindTerritories(boardToCompute Board) ([]Territory, error) {
	extractedTerritories, err := extractTerritories(boardToCompute)
	if err != nil {
		return nil, err
	}
	territories := []Territory{}
	for _, territoryToCompute := range extractedTerritories {
		territories = append(territories, findTerritoryWinnerAndGolds(territoryToCompute))
	}
	return territories, nil
}
