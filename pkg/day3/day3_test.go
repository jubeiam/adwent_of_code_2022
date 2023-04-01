package day3

import "testing"

func TestSumOfThePrioritiesOfBothCompartmentsOfEachRucksack(t *testing.T) {
	expected := 157
	actual := SumOfThePrioritiesOfBothCompartmentsOfEachRucksack("/test.txt")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSplitRucksackToCompartments(t *testing.T) {
	oneExpected := "ab"
	twoExpected := "cd"
	one, two := splitRucksackToCompartments("abcd")

	// Assert
	if one != oneExpected {
		t.Errorf("Expected %s but got %s", oneExpected, one)
	}

	// Assert
	if two != twoExpected {
		t.Errorf("Expected %s but got %s", twoExpected, two)
	}
}

func TestFindCommonChar(t *testing.T) {
	expected := "b"
	actual := findCommonChar("abx", "egb")

	// Assert
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestChar2Priority(t *testing.T) {
	expected := 1
	actual := char2priority("a")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	expected = 26
	actual = char2priority("z")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	expected = 27
	actual = char2priority("A")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	expected = 52
	actual = char2priority("Z")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSumOfTheBadgesOfThreeElf(t *testing.T) {
	expected := 70
	actual := SumOfTheBadgesOfThreeElf("/test.txt")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
