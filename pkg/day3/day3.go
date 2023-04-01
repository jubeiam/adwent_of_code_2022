package day3

import (
	"advent_of_code/pkg/helpers"
	"unicode"
)

func splitRucksackToCompartments(rucksack string) (string, string) {
	length := len(rucksack)

	return rucksack[:length/2], rucksack[length/2:]
}

func findCommonChar(one string, two string) string {
	var commonChar string

	for i := 0; i < len(one); i++ {
		for j := 0; j < len(two); j++ {
			if one[i] == two[j] {
				return string(one[i])
			}
		}
	}

	return commonChar
}

func findCommonChars(one string, two string) string {
	commonChar := ""

	for i := 0; i < len(one); i++ {
		for j := 0; j < len(two); j++ {
			if one[i] == two[j] {
				commonChar += string(one[i])
			}
		}
	}

	return commonChar
}

func char2priority(char string) int {
	r := []rune(char)
	ascii := int(r[0])
	delta := 0

	if unicode.IsLower(r[0]) {
		delta = 1 - 97
	}

	if unicode.IsUpper(r[0]) {
		delta = 27 - 65
	}

	return ascii + delta
}

func SumOfThePrioritiesOfBothCompartmentsOfEachRucksack(file string) int {
	rows := helpers.SplitByNewline(helpers.LoadFile(file))
	sum := 0
	for _, row := range rows {
		one, two := splitRucksackToCompartments(row)
		commonChar := findCommonChar(one, two)
		priority := char2priority(commonChar)
		sum += priority
	}

	return sum
}

func SumOfTheBadgesOfThreeElf(file string) int {
	rows := helpers.SplitByNewline(helpers.LoadFile(file))
	sum := 0
	commonChar := ""

	one := ""
	two := ""

	for _, row := range rows {
		if one == "" {
			one = row
			// log.Println("ONE", one)
			continue
		}

		if two == "" {
			two = findCommonChars(one, row)
			// log.Println("TWO", two, row)
			continue
		}

		commonChar = findCommonChar(two, row)
		// log.Println("THREE", commonChar, row)

		priority := char2priority(commonChar)
		sum += priority
		one = ""
		two = ""
	}

	return sum
}
