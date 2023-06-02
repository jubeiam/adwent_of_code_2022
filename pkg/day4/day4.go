package day4

import (
	"advent_of_code/pkg/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Assignment struct {
	Original string
	Min      int
	Max      int
}

type Pair struct {
	Uno Assignment
	Due Assignment
}

func (a Assignment) Overlaps(b Assignment) bool {
	return a.Min <= b.Min && b.Max <= a.Max
}

func (a Assignment) Contains(b Assignment) bool {
	return b.Min <= a.Max && b.Max >= a.Min
}

func normalizeParirs(file string) []Pair {
	listOfPairs := []Pair{}

	rows := helpers.SplitByNewline(helpers.LoadFile(file))

	for row := range rows {
		pairs := strings.Split(rows[row], ",")

		ass1 := strings.Split(pairs[0], "-")
		ass2 := strings.Split(pairs[1], "-")

		max1, _ := strconv.Atoi(ass1[1])
		max2, _ := strconv.Atoi(ass2[1])

		min1, _ := strconv.Atoi(ass1[0])
		min2, _ := strconv.Atoi(ass2[0])

		a1 := Assignment{Original: pairs[0], Max: max1, Min: min1}
		a2 := Assignment{Original: pairs[1], Max: max2, Min: min2}

		p := Pair{a1, a2}

		listOfPairs = append(listOfPairs, p)
	}

	return listOfPairs
}

func FullyContains(file string) int {
	listOfPairs := normalizeParirs(file)
	found := 0
	for _, pair := range listOfPairs {

		if pair.Uno.Overlaps(pair.Due) {
			found++
			fmt.Println(pair.Uno.Original, "overlaps", pair.Due.Original)
			continue
		}

		if pair.Due.Overlaps(pair.Uno) {
			found++
			fmt.Println(pair.Due.Original, "overlaps", pair.Uno.Original)
		}
	}

	return found
}

func AnyContains(file string) int {
	listOfPairs := normalizeParirs(file)

	found := 0
	for _, pair := range listOfPairs {

		if pair.Uno.Contains(pair.Due) {
			found++
			fmt.Println(pair.Uno.Original, "contains", pair.Due.Original)
			continue
		}

		if pair.Due.Contains(pair.Uno) {
			found++
			fmt.Println(pair.Due.Original, "contains", pair.Uno.Original)
		}
	}

	return found
}
