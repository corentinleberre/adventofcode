package day6

import "adventofcode/src/util"

func ChallengeOfTheDay(sequence string) util.ChallengeResult {
	puzzle1 := findFirstUniqueMarker(sequence, 4)
	puzzle2 := findFirstUniqueMarker(sequence, 14)
	return util.ChallengeResult{Puzzle1: puzzle1, Puzzle2: puzzle2}
}

func findFirstUniqueMarker(sequence string, nbChar int) int {
	for i := 0; i < len(sequence); i++ {
		if len(sequence)-i > nbChar-1 {
			if AreAllCharactersUnique(sequence[i : i+nbChar]) {
				return i + nbChar
			}
		}
	}
	return -1
}

func AreAllCharactersUnique(str string) bool {
	seen := make(map[rune]bool)

	for _, ch := range str {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}

	return true
}
