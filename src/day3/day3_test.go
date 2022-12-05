package day3

import (
	"adventofcode/src/util"
	"fmt"
	"testing"
)

func TestChallenge3(t *testing.T) {
	datasets, _ := util.ParseFile("./dataset.txt")
	result := challenge3(datasets)

	fmt.Printf("\nSomme des items : %v", result.Puzzle1)
	fmt.Printf("\nSomme des items : %v\n", result.Puzzle2)
}
