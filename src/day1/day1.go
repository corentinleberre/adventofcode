package day1

import (
	"adventofcode/src/util"
	"github.com/samber/lo"
	"sort"
	"strconv"
)

func challenge1(fruitsPickedByElphs []string) util.ChallengeResult {
	amountOfCaloriesByElph := []int{0}
	elphIndex := 0

	for _, fruitsPickedByElph := range fruitsPickedByElphs {
		if len(fruitsPickedByElph) > 0 {
			var fruit, _ = strconv.Atoi(fruitsPickedByElph)
			amountOfCaloriesByElph[elphIndex] += fruit
		} else {
			amountOfCaloriesByElph = append(amountOfCaloriesByElph, 0)
			elphIndex++
		}
	}

	sortedSlice := append([]int{}, amountOfCaloriesByElph...)
	sort.Ints(sortedSlice)

	maxValues := lo.Subset(sortedSlice, -3, 3)

	return util.ChallengeResult{
		Puzzle1: maxValues[len(maxValues)-1],
		Puzzle2: lo.Sum(maxValues),
	}
}
