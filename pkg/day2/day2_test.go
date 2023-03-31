package day2

import "testing"

func TestGetScore(t *testing.T) {
	score := GetScore("/test.txt")

	if score != 15 {
		t.Errorf("Expected 15, got %d", score)
	}
}

func TestGetHandScore(t *testing.T) {
	score := getHandScore("C")

	if score != 3 {
		t.Errorf("Expected 3, got %d", score)
	}

	score = getHandScore("Z")

	if score != 3 {
		t.Errorf("Expected 3, got %d", score)
	}

	score = getHandScore("Y")

	if score != 2 {
		t.Errorf("Expected 2, got %d", score)
	}

	score = getHandScore("B")

	if score != 2 {
		t.Errorf("Expected 2, got %d", score)
	}

	score = getHandScore("X")

	if score != 1 {
		t.Errorf("Expected 1, got %d", score)
	}

	score = getHandScore("A")

	if score != 1 {
		t.Errorf("Expected 1, got %d", score)
	}
}

func TestCompareHands(t *testing.T) {
	score := compareHands("A", "X")

	if score != 3 {
		t.Errorf("Expected 3, got %d", score)
	}

	score = compareHands("B", "Y")

	if score != 3 {
		t.Errorf("Expected 3, got %d", score)
	}

	score = compareHands("C", "Z")

	if score != 3 {
		t.Errorf("Expected 3, got %d", score)
	}

	score = compareHands("C", "X")

	if score != 6 {
		t.Errorf("Expected 6, got %d", score)
	}

	score = compareHands("A", "Y")

	if score != 6 {
		t.Errorf("Expected 6, got %d", score)
	}

	score = compareHands("B", "Z")

	if score != 6 {
		t.Errorf("Expected 6, got %d", score)
	}
}

func TestGetPredictedScore(t *testing.T) {
	score := GetPredictedScore("/test.txt")

	if score != 12 {
		t.Errorf("Expected 12, got %d", score)
	}
}
