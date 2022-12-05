package day1

import (
	"adventofcode/src/util"
	"fmt"
	"testing"
)

func TestChallenge1(t *testing.T) {
	datasets, _ := util.ParseFile("./dataset.txt")
	result := challenge1(datasets)

	fmt.Printf("\nPuzzle 1 :\nTotal calories picked by the best Elph : %v", result.Puzzle1)
	fmt.Printf("\nPuzzle 2 :\nTotal calories picked by the 3 bests Elphs : %v\n", result.Puzzle2)
}
