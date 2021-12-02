package day02

import (
	"aoc-2021/util"
	"strings"
)

type Step struct {
	index     int
	direction string
	amount    int
}

type Position struct {
	depth      int
	horizontal int
	aim        int
}

func calculatePosition(course []string) int {
	var position Position

	for i, line := range course {
		steps := strings.Split(line, " ")
		step := Step{i, steps[0], util.ConvertStrToInt(steps[1])}

		switch step.direction {
		case "down":
			position.depth += step.amount
		case "up":
			position.depth -= step.amount
		case "forward":
			position.horizontal += step.amount
		}
	}

	return position.depth * position.horizontal
}

func calculatePositionPartTwo(course []string) int {
	var position Position

	for i, line := range course {
		steps := strings.Split(line, " ")
		step := Step{i, steps[0], util.ConvertStrToInt(steps[1])}

		switch step.direction {
		case "down":
			position.aim += step.amount
		case "up":
			position.aim -= step.amount
		case "forward":
			position.horizontal += step.amount
			position.depth += position.aim * step.amount
		}
	}

	return position.depth * position.horizontal
}
