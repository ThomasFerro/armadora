package warrior_test

import (
	"testing"

	"github.com/ThomasFerro/armadora/game/warrior"
)

func TestDistributeTheWarriorsForTwoPlayers(t *testing.T) {
	warriorsToDistribute := warrior.WarriorsToDistribute(2)

	if warriorsToDistribute.OnePoint() != 11 {
		t.Error("A player does not have the requested number of 1 point warriors")
		return
	}
	if warriorsToDistribute.TwoPoints() != 2 {
		t.Error("A player does not have the requested number of 2 points warriors")
		return
	}
	if warriorsToDistribute.ThreePoints() != 3 {
		t.Error("A player does not have the requested number of 3 points warriors")
		return
	}
	if warriorsToDistribute.FourPoints() != 1 {
		t.Error("A player does not have the requested number of 4 points warriors")
		return
	}
	if warriorsToDistribute.FivePoints() != 1 {
		t.Error("A player does not have the requested number of 5 points warriors")
	}
}

func TestDistributeTheWarriorsForThreePlayers(t *testing.T) {
	warriorsToDistribute := warrior.WarriorsToDistribute(3)

	if warriorsToDistribute.OnePoint() != 7 {
		t.Error("A player does not have the requested number of 1 point warriors")
		return
	}
	if warriorsToDistribute.TwoPoints() != 2 {
		t.Error("A player does not have the requested number of 2 points warriors")
		return
	}
	if warriorsToDistribute.ThreePoints() != 1 {
		t.Error("A player does not have the requested number of 3 points warriors")
		return
	}
	if warriorsToDistribute.FourPoints() != 1 {
		t.Error("A player does not have the requested number of 4 points warriors")
		return
	}
	if warriorsToDistribute.FivePoints() != 0 {
		t.Error("A player does not have the requested number of 5 points warriors")
	}
}

func TestDistributeTheWarriorsForFourPlayers(t *testing.T) {
	warriorsToDistribute := warrior.WarriorsToDistribute(4)

	if warriorsToDistribute.OnePoint() != 5 {
		t.Error("A player does not have the requested number of 1 point warriors")
		return
	}
	if warriorsToDistribute.TwoPoints() != 1 {
		t.Error("A player does not have the requested number of 2 points warriors")
		return
	}
	if warriorsToDistribute.ThreePoints() != 1 {
		t.Error("A player does not have the requested number of 3 points warriors")
		return
	}
	if warriorsToDistribute.FourPoints() != 1 {
		t.Error("A player does not have the requested number of 4 points warriors")
		return
	}
	if warriorsToDistribute.FivePoints() != 0 {
		t.Error("A player does not have the requested number of 5 points warriors")
	}
}
