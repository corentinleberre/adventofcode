package day8

import (
	"adventofcode/src/util"
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"testing"
)

func TestChallengeOfTheDay(t *testing.T) {
	dataset, _ := util.ParseFile("./dataset.txt")

	data := lo.Map(dataset, func(item string, _ int) []int {
		var values []int
		for _, char := range item {
			value, _ := strconv.Atoi(string(char))
			values = append(values, value)
		}
		return values
	})

	result := ChallengeOfTheDay(data)
	fmt.Printf("\nTree visible from outside of the grid : %d", result.Puzzle1)
	fmt.Printf("\nMax scenic score on the grid : %d\n", result.Puzzle2)
}
