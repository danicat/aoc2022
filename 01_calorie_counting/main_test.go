package main

import (
	"strings"
	"testing"
)

func TestCalorieCount(t *testing.T) {
	input := `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	expected := 24000
	result := MaxCalorieCount(strings.NewReader(input))

	if expected != result {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
