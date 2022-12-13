package main

import (
	"strings"
	"testing"
)

func TestRockPaperScissors(t *testing.T) {
	input := `
A Y
B X
C Z
`

	expected := 15
	result := RockPaperScissors(strings.NewReader(input), fixed)

	if expected != result {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestCleverRockPaperScissors(t *testing.T) {
	input := `
A Y
B X
C Z
`

	expected := 12
	result := RockPaperScissors(strings.NewReader(input), clever)

	if expected != result {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
