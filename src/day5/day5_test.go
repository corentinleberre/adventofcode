package day5

import (
	"adventofcode/src/util"
	"github.com/samber/lo"
	"testing"
)

func TestChallengeOfTheDay(t *testing.T) {
	datasets, _ := util.ParseFile("dataset.txt")

	containers := lo.Subset(datasets, 10, uint(len(datasets)))

	parseArrayOfContainers(containers)
}
