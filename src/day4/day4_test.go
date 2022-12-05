package day4

import (
	"adventofcode/src/util"
	"fmt"
	"testing"
)

func TestChallenge4(t *testing.T) {
	datasets, _ := util.ParseFile("./dataset.txt")
	result := challenge4(datasets)

	fmt.Printf("\nFull Overlap : %d", result.Puzzle1)
	fmt.Printf("\nPartial Overlap : %d\n", result.Puzzle2)
}
