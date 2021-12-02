package day02

import (
	"aoc-2021/util"
	"fmt"
	"testing"
)

var testInput2 = []string{"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func Test__Day_Two_Part_One(t *testing.T) {
	expected := 150
	actual := calculatePosition(testInput2)

	if expected != actual {
		t.Errorf("Expected %d does not match actual, %d", expected, actual)
	} else {
		fmt.Printf("\nDay 2 Part 1 result: %d \n\n", calculatePosition(util.ReadInput(2)))
	}
}

func Test__Day_Two_Part_Two(t *testing.T) {
	expected := 900
	actual := calculatePositionPartTwo(testInput2)

	if expected != actual {
		t.Errorf("Expected %d does not match actual, %d", expected, actual)
	} else {
		fmt.Printf("\nDay 2 Part 2 result: %d \n\n", calculatePositionPartTwo(util.ReadInput(2)))
	}
}
