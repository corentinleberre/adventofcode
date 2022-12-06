package day6

import (
	"adventofcode/src/util"
	"fmt"
	"testing"
)

func TestChallengeOfTheDay(t *testing.T) {
	dataset, _ := util.ParseFile("./dataset.txt")
	result := ChallengeOfTheDay(dataset[0])

	fmt.Printf("\nStart-of-packet marker detected at : %d", result.Puzzle1)
	fmt.Printf("\nStart-of-message marker detected at : %d\n", result.Puzzle2)
}
