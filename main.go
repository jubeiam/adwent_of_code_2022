package main

import (
	"advent_of_code/pkg/day1"
	"advent_of_code/pkg/day2"
	"advent_of_code/pkg/day3"
	"advent_of_code/pkg/day4"
	"log"
)

func main() {
	log.Println(day1.FindElfWithCalories("/pkg/day1/input.txt"))
	log.Println(day1.Find3ElfsWithCalories("/pkg/day1/input.txt"))

	log.Println(day2.GetScore("/pkg/day2/input.txt"))
	log.Println(day2.GetPredictedScore("/pkg/day2/input.txt"))

	log.Println(day3.SumOfThePrioritiesOfBothCompartmentsOfEachRucksack("/pkg/day3/input.txt"))
	log.Println(day3.SumOfTheBadgesOfThreeElf("/pkg/day3/input.txt"))

	log.Println(day4.FullyContains("/pkg/day4/input.txt"))
}
