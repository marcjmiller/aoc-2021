package day01

import (
	"aoc-2021/util"
	"strconv"
)

func countIncreases(input []string) (count int) {
	nums, e := parseInputToInts(input)
	util.Check(e)

	lastNum := nums[0]

	for _, line := range nums {
		if lastNum < line {
			count++
		}
		lastNum = line
	}

	return
}

func countIncreasesPartTwo(input []string) (count int) {
	nums, e := parseInputToInts(input)
	util.Check(e)

	lastWindowSum := nums[0] + nums[1] + nums[2]

	for i, _ := range nums {
		if i == 0 || i == len(nums)-1 {
			continue
		}

		thisWindowSum := windowSum(nums, i)
		if thisWindowSum > lastWindowSum {
			count++
		}
		lastWindowSum = thisWindowSum
	}

	return
}

func parseInputToInts(input []string) (nums []int, e error) {
	nums = make([]int, 0, len(input))

	for _, l := range input {
		if len(l) == 0 {
			continue
		}

		n, e := strconv.Atoi(l)
		util.Check(e)

		nums = append(nums, n)
	}
	return
}

func windowSum(nums []int, i int) int {
	return nums[i-1] + nums[i] + nums[i+1]
}
