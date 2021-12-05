package day03

import (
	"aoc-2021/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func runDiagnosticsPartOne(input []string) int {
	gamma := make([]string, len(input[0]))
	epsilon := make([]string, len(input[0]))

	for i := range strings.Split(input[0], "") {
		gamma = append(gamma, fmt.Sprint(getMostCommonBitAtPosition(input, i)))
		epsilon = append(epsilon, fmt.Sprint(getLeastCommonBitAtPosition(input, i)))
	}

	strEpsilon := strings.Join(epsilon, "")
	strGamma := strings.Join(gamma, "")
	intEpsilon, e := strconv.ParseInt(strEpsilon, 2, 64)
	util.Check(e)
	intGamma, e := strconv.ParseInt(strGamma, 2, 64)
	util.Check(e)

	return int(intEpsilon * intGamma)
}

func runDiagnosticsPartTwo(input []string) int {
	oRating := findORating(input)
	co2Rating := findCO2Rating(input)

	return int(oRating * co2Rating)
}

func getLeastCommonBitAtPosition(input []string, position ...int) (lsb int) {
	return int(math.Abs(float64(getMostCommonBitAtPosition(input, position[0]) - 1)))
}

func filterList(input []string, search int, position ...int) (output []string) {
	if len(position) == 0 {
		position = append(position, 0)
	}

	for _, v := range input {
		if strings.Split(v, "")[position[0]] == fmt.Sprint(search) {
			output = append(output, v)
		}
	}
	return
}

func getMostCommonBitAtPosition(input []string, position ...int) (mcb int) {
	var ones int
	var zeroes int

	for _, v := range input {
		ones += util.ConvertStrToInt(strings.Split(v, "")[position[0]])
	}

	zeroes = len(input) - ones
	if ones >= zeroes {
		mcb = 1
	} else {
		mcb = 0
	}

	return
}

func findORating(input []string, pos ...int) int {
	if len(pos) == 0 {
		pos = append(pos, 0)
	}

	if len(input) <= 1 {
		output, e := strconv.ParseInt(input[0], 2, 64)
		util.Check(e)
		return int(output)
	}

	nextInput := filterList(input, getMostCommonBitAtPosition(input, pos[0]), pos[0])

	return findORating(nextInput, pos[0]+1)
}

func findCO2Rating(input []string, pos ...int) int {
	if len(pos) == 0 {
		pos = append(pos, 0)
	}

	if len(input) <= 1 {
		output, e := strconv.ParseInt(input[0], 2, 64)
		util.Check(e)
		return int(output)
	}

	nextInput := filterList(input, getLeastCommonBitAtPosition(input, pos[0]), pos[0])

	return findCO2Rating(nextInput, pos[0]+1)
}
