package day2

import (
	"adventofcode/src/util"
	"fmt"
	"testing"
)

func TestRockPaperScissor(t *testing.T) {
	datasets, _ := util.ParseFile("./dataset.txt")
	result := challenge2(datasets)

	fmt.Printf("\nPartie 1 score final : %v", result.Puzzle1)
	fmt.Printf("\nPartie 2 score final : %v\n", result.Puzzle2)
}
