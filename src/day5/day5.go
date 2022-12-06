package day5

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

func parseArrayOfContainers(lines []string) {
	result := [][]string{
		{"F", "D", "B", "Z", "T", "J", "R", "N"},
		{"R", "S", "N", "J", "H"},
		{"C", "R", "N", "J", "G", "Z", "F", "Q"},
		{"F", "V", "N", "G", "R", "T", "Q"},
		{"L", "T", "Q", "F"},
		{"Q", "C", "W", "Z", "B", "R", "G", "N"},
		{"F", "C", "L", "S", "N", "H", "M"},
		{"D", "N", "Q", "M", "T", "J"},
		{"P", "G", "S"},
	}

	/*result := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}*/

	print(result)

	instructions := lo.Map(lines, func(line string, _ int) []int {
		return extractNumbers(line)
	})

	for _, instruction := range instructions {
		quantity := instruction[0]
		from := instruction[1] - 1
		to := instruction[2] - 1

		tempArray := []string{}

		for i := quantity - 1; i >= 0; i-- {
			tempArray = append(tempArray, result[from][len(result[from])-1])
			result[from] = result[from][:len(result[from])-1]
		}

		result[to] = append(result[to], tempArray...)

		print(tempArray, to)
	}

	finalState := []string{}

	for _, col := range result {
		finalState = append(finalState, col[len(col)-1])
	}

	print(finalState)
}

func extractNumbers(str string) []int {
	words := strings.Fields(str)
	re := regexp.MustCompile("[0-9]+")

	numbers := make([]int, 0)

	for _, word := range words {
		numStr := re.FindString(word)
		if numStr != "" {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
	}

	return numbers
}
