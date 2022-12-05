package day3

import (
	"adventofcode/src/util"
	"github.com/samber/lo"
	"sort"
	"strings"
)

func challenge3(rucksacks []string) util.ChallengeResult {
	return util.ChallengeResult{
		Puzzle1: puzzle1(rucksacks),
		Puzzle2: puzzle2(rucksacks),
	}
}

func puzzle2(rucksacks []string) int {
	finalScore := ""

	for i := 0; i < len(rucksacks); i += 3 {
		group := rucksacks[i : i+3]
		sort.Strings(group)
		for i := 0; i < len(group[2]); i++ {
			if strings.Contains(group[0], string(group[2][i])) && strings.Contains(group[1], string(group[2][i])) {
				finalScore += string(group[2][i])
				break
			}
		}
	}

	sum := mapStringToNumberAndGiveFinalScore(finalScore)
	return int(sum)
}

func puzzle1(rucksacks []string) int {
	rucksacksCompartments := lo.Map(rucksacks, func(item string, index int) []string {
		return []string{item[0 : len(item)/2], item[len(item)/2:]}
	})

	finalScore := ""

	for _, rucksack := range rucksacksCompartments {
		temp := ""
		for i := 0; i < len(rucksack[0]); i++ {
			if strings.Contains(rucksack[1], string(rucksack[0][i])) {
				if !strings.Contains(temp, string(rucksack[0][i])) {
					temp += string(rucksack[0][i])
				}
			}
		}
		finalScore += temp
	}

	sum := mapStringToNumberAndGiveFinalScore(finalScore)
	return int(sum)
}

func mapStringToNumberAndGiveFinalScore(finalScore string) rune {
	r := []rune(finalScore)
	r = lo.Map(r, func(item rune, index int) rune {
		if item <= 90 {
			return item - 38
		}
		return item - 96
	})
	return lo.Sum(r)
}
