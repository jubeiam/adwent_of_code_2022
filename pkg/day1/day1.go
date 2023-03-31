package day1

import (
	"advent_of_code/pkg/helpers"
	"sort"
	"strconv"
)

func FindElfWithCalories(file string) (calories int) {
	rows := helpers.SplitByNewline(helpers.LoadFile(file))
	caloriesList := []int{}

	sum := 0

	for _, row := range rows {
		if row == "" {
			caloriesList = append(caloriesList, sum)
			sum = 0
		} else {
			i, err := strconv.Atoi(row)
			helpers.Check(err)
			sum += i
		}
	}

	sort.Ints(caloriesList)

	return caloriesList[len(caloriesList)-1]
}

func Find3ElfsWithCalories(file string) (sumOfCalories int) {
	rows := helpers.SplitByNewline(helpers.LoadFile(file))
	caloriesList := []int{}

	sum := 0

	for _, row := range rows {
		if row == "" {
			if sum > 0 {
				caloriesList = append(caloriesList, sum)
			}
			sum = 0
		} else {
			i, err := strconv.Atoi(row)
			helpers.Check(err)
			sum += i
		}
	}
	// Add last sum
	caloriesList = append(caloriesList, sum)

	sort.Ints(caloriesList)

	return caloriesList[len(caloriesList)-1] + caloriesList[len(caloriesList)-2] + caloriesList[len(caloriesList)-3]
}
