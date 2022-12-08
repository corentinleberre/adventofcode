package day7

import (
	"adventofcode/src/util"
	"fmt"
	"testing"
)

func TestChallengeOfTheDay(t *testing.T) {
	dataset, _ := util.ParseFile("./dataset.txt")
	result := ChallengeOfTheDay(dataset)

	fmt.Printf("\nDirectoryAtMost100k : %d", result.Puzzle1)
	fmt.Printf("\nDirectory to delete : %d\n", result.Puzzle2)
}
