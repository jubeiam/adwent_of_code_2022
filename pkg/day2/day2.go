package day2

import (
	"advent_of_code/pkg/helpers"
)

var (
	winingHand = map[string]string{"A": "Y", "B": "Z", "C": "X"}
	losingHand = map[string]string{"A": "Z", "B": "X", "C": "Y"}
	tieHand    = map[string]string{"A": "X", "B": "Y", "C": "Z"}
)

func getHandScore(hand string) int {
	switch hand {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		panic("Invalid hand")
	}
}

func getHandName(hand string) string {
	switch hand {
	case "A", "X":
		return "rock"
	case "B", "Y":
		return "paper"
	case "C", "Z":
		return "scissors"
	default:
		panic("Invalid hand")
	}
}

func compareHands(hand1 string, hand2 string) int {
	if winingHand[hand1] == hand2 {
		return 6
	}

	if tieHand[hand1] == hand2 {
		return 3
	}

	return 0
}

func predictHand(hand1 string, predict string) string {
	if predict == "Z" {
		return winingHand[hand1]
	}

	if predict == "Y" {
		return tieHand[hand1]
	}

	return losingHand[hand1]
}

func GetScore(file string) int {
	rows := helpers.SplitByNewline(helpers.LoadFile(file))
	sum := 0

	for _, row := range rows {
		oponent := row[0:1]
		me := row[2:3]

		round := compareHands(oponent, me)
		myHand := getHandScore(me)

		// log.Println(i, getHandName(oponent), getHandName(me), round, myHand)

		sum += round
		sum += myHand
	}

	return sum
}

func GetPredictedScore(file string) int {
	rows := helpers.SplitByNewline(helpers.LoadFile(file))
	sum := 0

	for _, row := range rows {
		oponent := row[0:1]
		predict := row[2:3]
		me := predictHand(oponent, predict)

		round := compareHands(oponent, me)
		myHand := getHandScore(me)

		// log.Println(i, getHandName(oponent), getHandName(me), round, myHand)

		sum += round
		sum += myHand
	}

	return sum
}
