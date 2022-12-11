package day8

import (
	"adventofcode/src/util"
	"github.com/samber/lo"
)

func ChallengeOfTheDay(dataset [][]int) util.ChallengeResult {
	var edgeTree int
	var visibleTree int
	var scenicScore []int
	for i := 0; i < len(dataset); i++ {
		for j := 0; j < len(dataset); j++ {
			if i == 0 || i == len(dataset)-1 || j == 0 || j == len(dataset[i])-1 {
				edgeTree++
			} else {
				tree := dataset[i][j]
				datasetCopy := make([][]int, len(dataset))
				for index := range dataset {
					datasetCopy[index] = make([]int, len(dataset[index]))
					copy(datasetCopy[index], dataset[index])
				}
				challenge1(datasetCopy, i, j, tree, &visibleTree)
				challenge2(datasetCopy, i, j, tree, &scenicScore)
			}
		}
	}
	return util.ChallengeResult{Puzzle1: edgeTree + visibleTree, Puzzle2: lo.Max(scenicScore)}
}

func challenge1(dataset [][]int, i int, j int, tree int, visibleTree *int) {
	left := isVisible(dataset[i][:j], tree)
	right := isVisible(dataset[i][j+1:], tree)
	up := isVisible(getColumn(dataset, j)[:i], tree)
	down := isVisible(getColumn(dataset, j)[i+1:], tree)
	if left || right || up || down {
		*visibleTree++
	}
}

func isVisible(dataset []int, value int) bool {
	return lo.EveryBy(dataset, func(item int) bool {
		return value > item
	})
}

func challenge2(dataset [][]int, i int, j int, tree int, scenicScore *[]int) {
	left := treeScore(lo.Reverse(dataset[i][:j]), tree)
	right := treeScore(dataset[i][j+1:], tree)
	up := treeScore(lo.Reverse(getColumn(dataset, j)[:i]), tree)
	down := treeScore(getColumn(dataset, j)[i+1:], tree)
	*scenicScore = append(*scenicScore, left*right*up*down)
}

func treeScore(treeRow []int, tree int) int {
	cpt := 0
	if len(treeRow) == 1 && treeRow[0] == 0 {
		return 0
	}
	for _, nextTree := range treeRow {
		visible := tree > nextTree
		blocked := tree <= nextTree
		if blocked {
			cpt++
		}
		if visible {
			cpt++
		} else {
			break
		}
	}
	return cpt
}

func getColumn(dataset [][]int, index int) []int {
	column := make([]int, len(dataset))
	for i, row := range dataset {
		column[i] = row[index]
	}
	return column
}
