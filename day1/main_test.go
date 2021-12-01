package day1

import (
	"fmt"
	"testing"
)

func Test__Day_One_Part_One(t *testing.T) {
	expected := 7
	actual := countIncreases(testInput)

	if expected != actual {
		t.Fatalf("Expected %d does not match actual, %d", expected, actual)
	}

	fmt.Printf("Day 1 Part 1 result: %d \n\n", countIncreases(day1input1))
}

func Test__Day_One_Part_TWo(t *testing.T) {
	expected := 5
	actual := countIncreasesPartTwo(testInput)
	if expected != actual {
		t.Fatalf("Expected %d does not match actual, %d", expected, actual)
	}

	fmt.Printf("Day 1 Part 2 result: %d \n\n", countIncreasesPartTwo(day1input1))
}
