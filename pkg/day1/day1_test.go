package day1

import (
	test "testing"
)

func TestFindElfWithCalories(test *test.T) {
	calories := FindElfWithCalories("/test.txt")

	if calories != 24000 {
		test.Errorf("Expected 24000, got %d", calories)
	}
}

func TestFind3ElfsWithCalories(test *test.T) {
	calories := Find3ElfsWithCalories("/test.txt")

	if calories != 45000 {
		test.Errorf("Expected 45000, got %d", calories)
	}
}
