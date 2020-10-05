package score_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/score"
)

func TestComputeScoresForOnePlayer(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(12, []int{0}),
	}

	scores := score.ComputeScores(territories)

	if len(scores) != 1 {
		t.Errorf("Expected to get one score, got this instead: %v", scores)
		return
	}

	winnerScore, hasWinner := scores[1]

	if !hasWinner {
		t.Errorf("Expected to get a winning player, got this instead: %v", scores)
		return
	}

	if winnerScore.Player() != 0 || winnerScore.TotalGold() != 12 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
	}
}

func TestComputeScoresForMultiplePlayers(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(2, []int{2}),
		board.NewTerritory(4, []int{1}),
		board.NewTerritory(7, []int{3}),
		board.NewTerritory(12, []int{0}),
		board.NewTerritory(1, []int{0}),
	}

	scores := score.ComputeScores(territories)

	if len(scores) != 4 {
		t.Errorf("Expected to get four scores, got this instead: %v", scores)
		return
	}

	if winnerScore, hasWinner := scores[1]; !hasWinner || winnerScore.Player() != 0 || winnerScore.TotalGold() != 13 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
		return
	}

	if secondPlayer, hasSecondPlayer := scores[2]; !hasSecondPlayer || secondPlayer.Player() != 3 || secondPlayer.TotalGold() != 7 {
		t.Errorf("Invalid second player, got this instead: %v", secondPlayer)
		return
	}

	if thirdPlayer, hasThirdPlayer := scores[3]; !hasThirdPlayer || thirdPlayer.Player() != 1 || thirdPlayer.TotalGold() != 4 {
		t.Errorf("Invalid third player, got this instead: %v", thirdPlayer)
		return
	}

	if fourthPlayer, hasFourthPlayer := scores[4]; !hasFourthPlayer || fourthPlayer.Player() != 2 || fourthPlayer.TotalGold() != 2 {
		t.Errorf("Invalid fourth player, got this instead: %v", fourthPlayer)
	}
}

func TestSplitTheGold(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(2, []int{1}),
		board.NewTerritory(12, []int{0, 1}),
	}

	scores := score.ComputeScores(territories)

	if winnerScore, hasWinner := scores[1]; !hasWinner || winnerScore.Player() != 1 || winnerScore.TotalGold() != 8 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
		return
	}

	if secondPlayer, hasSecondPlayer := scores[2]; !hasSecondPlayer || secondPlayer.Player() != 0 || secondPlayer.TotalGold() != 6 {
		t.Errorf("Invalid second player, got this instead: %v", secondPlayer)
	}
}

func TestSplitTheGoldIndivisible(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(2, []int{0}),
		board.NewTerritory(7, []int{0, 1}),
	}

	scores := score.ComputeScores(territories)

	if winnerScore, hasWinner := scores[1]; !hasWinner || winnerScore.Player() != 0 || winnerScore.TotalGold() != 5 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
		return
	}

	if secondPlayer, hasSecondPlayer := scores[2]; !hasSecondPlayer || secondPlayer.Player() != 1 || secondPlayer.TotalGold() != 3 {
		t.Errorf("Invalid second player, got this instead: %v", secondPlayer)
	}
}

func TestComputeScoresWithTies(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(5, []int{0}),
		board.NewTerritory(7, []int{0}),
		board.NewTerritory(5, []int{1}),
		board.NewTerritory(5, []int{1}),
		board.NewTerritory(2, []int{1}),
	}

	scores := score.ComputeScores(territories)

	if winnerScore, hasWinner := scores[1]; !hasWinner || winnerScore.Player() != 0 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
		return
	}

	if secondPlayer, hasSecondPlayer := scores[2]; !hasSecondPlayer || secondPlayer.Player() != 1 {
		t.Errorf("Invalid second player, got this instead: %v", secondPlayer)
	}
}

func TestComputeScoresWithTiesOnASingleTerritory(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(7, []int{0, 1}),
	}

	scores := score.ComputeScores(territories)

	if winnerScore, hasWinner := scores[1]; !hasWinner || winnerScore.Player() != 0 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
		return
	}

	if secondPlayer, hasSecondPlayer := scores[2]; !hasSecondPlayer || secondPlayer.Player() != 1 {
		t.Errorf("Invalid second player, got this instead: %v", secondPlayer)
	}
}

func TestComputeScoresWithComplexTies(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(7, []int{0}),
		board.NewTerritory(7, []int{1}),
		board.NewTerritory(6, []int{0}),
		board.NewTerritory(6, []int{1}),
		board.NewTerritory(5, []int{0}),
		board.NewTerritory(5, []int{1}),
		board.NewTerritory(5, []int{1}),
		board.NewTerritory(3, []int{0}),
		board.NewTerritory(2, []int{0}),
	}

	scores := score.ComputeScores(territories)

	if winnerScore, hasWinner := scores[1]; !hasWinner || winnerScore.Player() != 1 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
		return
	}

	if secondPlayer, hasSecondPlayer := scores[2]; !hasSecondPlayer || secondPlayer.Player() != 0 {
		t.Errorf("Invalid second player, got this instead: %v", secondPlayer)
	}
}
func TestComputeScoresWithComplexTiesWithThreePlayers(t *testing.T) {
	territories := []board.Territory{
		board.NewTerritory(7, []int{0}),
		board.NewTerritory(7, []int{1}),
		board.NewTerritory(7, []int{2}),
		board.NewTerritory(6, []int{0}),
		board.NewTerritory(6, []int{1}),
		board.NewTerritory(6, []int{2}),
		board.NewTerritory(5, []int{0}),
		board.NewTerritory(5, []int{1}),
		board.NewTerritory(5, []int{2}),
		board.NewTerritory(5, []int{0}),
		board.NewTerritory(3, []int{2}),
		board.NewTerritory(2, []int{2}),
		board.NewTerritory(2, []int{1}),
		board.NewTerritory(2, []int{1}),
		board.NewTerritory(1, []int{1}),
	}

	scores := score.ComputeScores(territories)

	if winnerScore, hasWinner := scores[1]; !hasWinner || winnerScore.Player() != 0 {
		t.Errorf("Invalid winning player, got this instead: %v", winnerScore)
		return
	}

	if secondPlayer, hasSecondPlayer := scores[2]; !hasSecondPlayer || secondPlayer.Player() != 2 {
		t.Errorf("Invalid second player, got this instead: %v", secondPlayer)
		return
	}

	if thirdPlayer, hasThridPlayer := scores[3]; !hasThridPlayer || thirdPlayer.Player() != 1 {
		t.Errorf("Invalid third player, got this instead: %v", thirdPlayer)
	}
}
