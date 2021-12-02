package day01

import (
	"fmt"
	"testing"

	util "aoc2021/util"
)

var testInput = []string{"199",
	"200",
	"208",
	"210",
	"200",
	"207",
	"240",
	"269",
	"260",
	"263"}

func Test__Day_One_Part_One(t *testing.T) {
	expected := 7
	actual := countIncreases(testInput)

	if expected != actual {
		t.Fatalf("Expected %d does not match actual, %d", expected, actual)
	}

	fmt.Printf("Day 1 Part 1 result: %d \n\n", countIncreases(util.ReadInput(1)))
}

func Test__Day_One_Part_TWo(t *testing.T) {
	expected := 5
	actual := countIncreasesPartTwo(testInput)
	if expected != actual {
		t.Fatalf("Expected %d does not match actual, %d", expected, actual)
	}

	fmt.Printf("Day 1 Part 2 result: %d \n\n", countIncreasesPartTwo(util.ReadInput(1)))
}
