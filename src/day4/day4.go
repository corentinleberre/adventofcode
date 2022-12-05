package day4

import (
	"adventofcode/src/util"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func challenge4(pairs []string) util.ChallengeResult {
	cleanedPairs := lo.Map(pairs, func(pair string, _ int) [][]int {
		values := lo.Map(strings.Split(pair, ","), func(item string, _ int) []int {
			return lo.Map(strings.Split(item, "-"), func(item string, _ int) int {
				value, _ := strconv.Atoi(item)
				return value
			})
		})
		return [][]int{util.MakeRange(values[0][0], values[0][1]), util.MakeRange(values[1][0], values[1][1])}
	})

	totalFullOverlapping := 0
	totalPartialOverlapping := 0

	for _, pair := range cleanedPairs {
		var smallerGroup []int
		var biggerGroup []int

		if len(pair[0]) > len(pair[1]) {
			smallerGroup = append(smallerGroup, pair[1]...)
			biggerGroup = append(biggerGroup, pair[0]...)
		} else {
			smallerGroup = append(smallerGroup, pair[0]...)
			biggerGroup = append(biggerGroup, pair[1]...)
		}

		count := 0

		for i := 0; i < len(smallerGroup); i++ {
			for j := 0; j < len(biggerGroup); j++ {
				if smallerGroup[i] == biggerGroup[j] {
					count++
				}
			}
		}

		if count == len(smallerGroup) {
			totalFullOverlapping++
		}

		if count > 0 {
			totalPartialOverlapping++
		}

	}

	return util.ChallengeResult{
		Puzzle1: totalFullOverlapping,
		Puzzle2: totalPartialOverlapping,
	}
}
