package main

import (
	"fmt"
	"github.com/samber/lo"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// challenge1()
	// challenge2()
	// challenge3()
	// challenge3Part2()
	// challenge4Part1()
	challenge4Part2()
}

func challenge4Part2() {
	pairs, _ := ParseFile("./challenge4.txt")

	cleanedPairs := lo.Map(pairs, func(pair string, _ int) [][]int {
		values := lo.Map(strings.Split(pair, ","), func(item string, _ int) []int {
			return lo.Map(strings.Split(item, "-"), func(item string, _ int) int {
				value, _ := strconv.Atoi(item)
				return value
			})
		})
		return [][]int{makeRange(values[0][0], values[0][1]), makeRange(values[1][0], values[1][1])}
	})

	total := 0

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

		if count > 0 {
			total++
		}

	}

	fmt.Printf("\nOverlap : %d", total)
}

func challenge4Part1() {
	pairs, _ := ParseFile("./challenge4.txt")

	cleanedPairs := lo.Map(pairs, func(pair string, _ int) [][]int {
		values := lo.Map(strings.Split(pair, ","), func(item string, _ int) []int {
			return lo.Map(strings.Split(item, "-"), func(item string, _ int) int {
				value, _ := strconv.Atoi(item)
				return value
			})
		})
		return [][]int{makeRange(values[0][0], values[0][1]), makeRange(values[1][0], values[1][1])}
	})

	total := 0

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
			total++
		}

	}

	print(total)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func challenge3Part2() {
	rucksacks, _ := ParseFile("./challenge3.txt")

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

	r := []rune(finalScore)
	r = lo.Map(r, func(item rune, index int) rune {
		if item <= 90 {
			return item - 38
		}
		return item - 96
	})

	sum := lo.Sum(r)

	fmt.Printf("\nSomme des items : %v", sum)
}

func challenge3() {
	rucksacks, _ := ParseFile("./challenge3.txt")

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

	r := []rune(finalScore)
	r = lo.Map(r, func(item rune, index int) rune {
		if item <= 90 {
			return item - 38
		}
		return item - 96
	})

	sum := lo.Sum(r)

	fmt.Printf("\nSomme des items : %v", sum)
}

func challenge1() {
	fruitsPickedByElphs, _ := ParseFile("./challenge1.txt")

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

	print("\nPuzzle 1 :\nTotal calories picked by the best Elph : ", maxValues[len(maxValues)-1])
	print("\nPuzzle 2 :\nTotal calories picked by the 3 bests Elphs : ", lo.Sum(maxValues))
}

func challenge2() {
	datasets, _ := ParseFile("./challenge2.txt")

	finalScore := 0
	// A | X Rock = 1
	// B | Y Paper = 2
	// C | Z Scissor = 3

	combinations := part2Combinations()

	for _, dataset := range datasets {
		finalScore += combinations[dataset]
	}

	print("\n Score final : ", finalScore)
}

func part1Combinations() map[string]int {
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
