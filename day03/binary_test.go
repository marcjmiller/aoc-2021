package day03

import (
	"aoc-2021/util"
	"fmt"
	"testing"
)

var testInput3 = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func Test__RunDiagnosticsPartOne(t *testing.T) {
	expected := 198
	actual := runDiagnosticsPartOne(testInput3)

	if expected != actual {
		t.Errorf("Expected %d does not match actual, %d", expected, actual)
	} else {
		fmt.Printf("\nDay 3 Part 1 result: %d \n\n", runDiagnosticsPartOne(util.ReadInput(3)))
	}
}

func Test__RunDiagnosticsPartTwo(t *testing.T) {
	expected := 230
	actual := runDiagnosticsPartTwo(testInput3)

	if expected != actual {
		t.Errorf("Expected %d does not match actual, %d", expected, actual)
	} else {
		fmt.Printf("\nDay 3 Part 2 result: %d \n\n", runDiagnosticsPartTwo(util.ReadInput(3)))
	}
}
