package gold

import (
	"math/rand"
	"time"
)

var GoldStacks = []int{
	3,
	4,
	4,
	5,
	5,
	6,
	6,
	7,
}

// GoldToDistribute Return the gold to be distributed
func GoldToDistribute() []int {
	// Shuffle method based on this example: https://yourbasic.org/golang/shuffle-slice-array/
	gold := GoldStacks
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(gold), func(i, j int) { gold[i], gold[j] = gold[j], gold[i] })
	return gold
}
