package gold

import (
	"testing"
)

type StackToFind struct {
	expected map[int]int
	found    map[int]int
}

func TestDistributeTheGoldStacks(t *testing.T) {
	goldToDistribute := GoldToDistribute()

	stackToFind := StackToFind{
		expected: map[int]int{
			3: 1,
			4: 2,
			5: 2,
			6: 2,
			7: 1,
		},
		found: map[int]int{
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
		},
	}

	for _, stack := range goldToDistribute {
		stackToFind.found[stack] = stackToFind.found[stack] + 1
	}

	for stackValue, expectedStackCount := range stackToFind.expected {
		if stackToFind.found[stackValue] != expectedStackCount {
			t.Errorf("There should be %v stacks of %v gold, but %v was found", expectedStackCount, stackValue, stackToFind.found[stackValue])
			return
		}
	}
}
