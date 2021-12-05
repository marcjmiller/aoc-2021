package day05

import (
	"aoc-2021/util"
	"fmt"
	"testing"
)

var testInput5 = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}

func Test__CalcHydrothermals_PartOne(t *testing.T) {
	expected := 4512
	actual := calcHydrothermalsPartOne(testInput5)

	if expected != actual {
		t.Errorf("Expected %d does not match actual, %d", expected, actual)
	} else {
		fmt.Printf("\nDay 3 Part 1 result: %d \n\n", calcHydrothermalsPartOne(util.ReadInput(5)))
	}
}

func Test__CalcHydrothermals_PartTwo(t *testing.T) {
	expected := 4512
	actual := calcHydrothermalsPartTwo(testInput5)

	if expected != actual {
		t.Errorf("Expected %d does not match actual, %d", expected, actual)
	} else {
		fmt.Printf("\nDay 3 Part 1 result: %d \n\n", calcHydrothermalsPartTwo(util.ReadInput(5)))
	}
}
