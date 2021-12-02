package util

import (
	"fmt"
	"testing"
)

func Test__ReadInputGetsAllLines(t *testing.T) {
	input := ReadInput(0)
	expected := 9
	actual := len(input)

	if expected != actual {
		t.Fatalf("Expected %v does not match actual, %v", expected, actual)
	} else {
		fmt.Printf("Expected: %v matches actual: %v \n", expected, actual)
	}
}

func Test__ReadInputGetsCorrectLines(t *testing.T) {
	input := ReadInput(0)
	expected := []string{
		"this",
		"is",
		"a",
		"test",
		"of",
		"the",
		"emergency",
		"broadcast",
		"system",
	}

	for i, actual := range input {
		if expected[i] != actual {
			t.Fatalf("Expected %v does not match actual, %v, at position %v", expected, actual, i)
		} else {
			fmt.Printf("Expected: %v - Actual: %v -- Match \n", expected[i], actual)
		}
	}
}
