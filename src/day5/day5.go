package day5

import (
	"github.com/samber/lo"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func supplyStacks(cratesState [][]string, lines []string) string {
	rearrangementProcedure := lo.Map(lines, func(line string, _ int) []int {
		return convertCraneInstructions(line)
	})

	for _, instruction := range rearrangementProcedure {
		quantity := instruction[0]
		from := instruction[1] - 1
		to := instruction[2] - 1

		movedCrateArray, leftCrateArray := crateMover9001Logic(cratesState, from, quantity)

		cratesState[to] = append(cratesState[to], movedCrateArray...)
		cratesState[from] = leftCrateArray
	}

	finalState := lo.Map(cratesState, func(crate []string, index int) string {
		return crate[len(crate)-1]
	})

	concatResult := lo.Reduce(finalState, func(agg string, item string, index int) string {
		return agg + item
	}, "")

	return concatResult
}

func crateMover9001Logic(cratesState [][]string, from int, quantity int) ([]string, []string) {
	tempArray := lo.Subset(cratesState[from], -quantity, math.MaxUint)
	restArray := lo.Subset(cratesState[from], 0, uint(len(cratesState[from])-quantity))
	return tempArray, restArray
}

func convertCraneInstructions(str string) []int {
	words := strings.Fields(str)
	re := regexp.MustCompile("[0-9]+")

	instructions := make([]int, 0)

	for _, word := range words {
		numStr := re.FindString(word)
		if numStr != "" {
			num, _ := strconv.Atoi(numStr)
			instructions = append(instructions, num)
		}
	}

	return instructions
}
