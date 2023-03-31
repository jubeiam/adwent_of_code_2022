package main

import (
	"advent_of_code/pkg/day1"
	"advent_of_code/pkg/day2"
	"log"
)

func main() {
	log.Println(day1.FindElfWithCalories("/pkg/day1/input.txt"))
	log.Println(day1.Find3ElfsWithCalories("/pkg/day1/input.txt"))

	log.Println(day2.GetScore("/pkg/day2/input.txt"))
	log.Println(day2.GetPredictedScore("/pkg/day2/input.txt"))
}
