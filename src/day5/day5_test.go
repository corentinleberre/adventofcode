package day5

import (
	"adventofcode/src/util"
	"fmt"
	"github.com/samber/lo"
	"testing"
)

func TestChallengeOfTheDay(t *testing.T) {
	datasets, _ := util.ParseFile("dataset.txt")

	cratesInitialState := [][]string{
		{"F", "D", "B", "Z", "T", "J", "R", "N"},
		{"R", "S", "N", "J", "H"},
		{"C", "R", "N", "J", "G", "Z", "F", "Q"},
		{"F", "V", "N", "G", "R", "T", "Q"},
		{"L", "T", "Q", "F"},
		{"Q", "C", "W", "Z", "B", "R", "G", "N"},
		{"F", "C", "L", "S", "N", "H", "M"},
		{"D", "N", "Q", "M", "T", "J"},
		{"P", "G", "S"},
	}

	containers := lo.Subset(datasets, 10, uint(len(datasets)))

	result := supplyStacks(cratesInitialState, containers)

	fmt.Printf("\n Result with the CrateMover 9001 : %s\n", result)
}
