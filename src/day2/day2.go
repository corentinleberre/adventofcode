package day2

import (
	"adventofcode/src/util"
)

func challenge2(gameCombinations []string) util.ChallengeResult {
	part1Total := 0
	part2Total := 0

	part1combinations := part1Combinations()
	part2combinations := part2Combinations()

	for _, dataset := range gameCombinations {
		part1Total += part1combinations[dataset]
		part2Total += part2combinations[dataset]
	}

	return util.ChallengeResult{
		Puzzle1: part1Total,
		Puzzle2: part2Total,
	}
}

func part1Combinations() map[string]int {
	// A | X Rock = 1
	// B | Y Paper = 2
	// C | Z Scissor = 3

	return map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6,
	}
}

func part2Combinations() map[string]int {
	// New rules

	// X Lose
	// Y Draw
	// Z Win

	// A Rock = 1
	// B Paper = 2
	// C Scissor = 3
	return map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7,
	}
}
