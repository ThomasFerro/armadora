package score

import (
	"fmt"
	"sort"

	"github.com/ThomasFerro/armadora/game/board"
)

// Scores The game's scores
type Scores map[int]Score

// Score One player's score
type Score interface {
	Player() int
	GoldStacks() []int
	TotalGold() int
}

type score struct {
	player     int
	goldStacks []int
}

func (s score) String() string {
	return fmt.Sprintf("Player %v with %v gold", s.Player(), s.TotalGold())
}

func (s score) Player() int {
	return s.player
}

func (s score) GoldStacks() []int {
	return s.goldStacks
}

func (s score) TotalGold() int {
	total := 0

	for _, nextStack := range s.GoldStacks() {
		total += nextStack
	}

	return total
}

// NewScore Create a new score
func NewScore(player int, goldStacks []int) Score {
	return score{
		player,
		goldStacks,
	}
}

func getScoreByPlayer(territories []board.Territory) map[int][]int {
	scoreByPlayer := map[int][]int{}

	for _, nextTerritory := range territories {
		for _, winningPlayer := range nextTerritory.WinningPlayers() {
			scoreByPlayer[winningPlayer] = append(scoreByPlayer[winningPlayer], nextTerritory.Gold()/len(nextTerritory.WinningPlayers()))
		}
	}

	return scoreByPlayer
}

type kv struct {
	Player     int
	GoldStacks []int
}

func sortTiedPlayers(firstPlayerStacks, secondPlayerStacks []int) bool {
	stackIndex := 0
	sortedFirstPlayerStacks := append([]int{}, firstPlayerStacks...)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedFirstPlayerStacks)))
	sortedSecondPlayerStacks := append([]int{}, secondPlayerStacks...)
	sort.Sort(sort.Reverse(sort.IntSlice(secondPlayerStacks)))

	for {
		if len(sortedFirstPlayerStacks) < stackIndex {
			return false
		}
		if len(sortedSecondPlayerStacks) < stackIndex {
			return true
		}
		if sortedFirstPlayerStacks[stackIndex] != sortedSecondPlayerStacks[stackIndex] {
			return sortedFirstPlayerStacks[stackIndex] > sortedSecondPlayerStacks[stackIndex]
		}
		stackIndex++
	}
}

func sortFunction(scoresToSort []kv) func(i, j int) bool {
	return func(firstScore, secondScore int) bool {
		totalFirstScore := 0
		for _, nextStack := range scoresToSort[firstScore].GoldStacks {
			totalFirstScore += nextStack
		}
		totalSecondScore := 0
		for _, nextStack := range scoresToSort[secondScore].GoldStacks {
			totalSecondScore += nextStack
		}
		if totalFirstScore == totalSecondScore {
			return sortTiedPlayers(scoresToSort[firstScore].GoldStacks, scoresToSort[secondScore].GoldStacks)
		}
		return totalFirstScore > totalSecondScore
	}
}

func sortScoresBasedOnGold(unsortedScores map[int][]int) []Score {
	scoresToSort := []kv{}

	for key, value := range unsortedScores {
		scoresToSort = append(scoresToSort, kv{
			Player:     key,
			GoldStacks: value,
		})
	}

	sort.Slice(
		scoresToSort,
		sortFunction(scoresToSort),
	)

	sortedScores := []Score{}

	for _, sortedScore := range scoresToSort {
		sortedScores = append(sortedScores, NewScore(sortedScore.Player, sortedScore.GoldStacks))
	}

	return sortedScores
}

// ComputeScores Compute the scores based on the territories
func ComputeScores(territories []board.Territory) Scores {
	scoreByPlayer := getScoreByPlayer(territories)
	sortedPlayers := sortScoresBasedOnGold(scoreByPlayer)

	returnedScored := Scores{}

	for _, sortedPlayer := range sortedPlayers {
		returnedScored[len(returnedScored)+1] = sortedPlayer
	}

	return returnedScored
}
