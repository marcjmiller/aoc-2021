package day1

import (
	"strconv"
	"strings"
)

func countIncreases(input string) (count int) {
	nums, e := parseInputToInts(input)
	if e != nil {
		panic(e)
	}
	lastNum := nums[0]

	for _, line := range nums {
		if lastNum < line {
			count++
		}
		lastNum = line
	}

	return
}

func countIncreasesPartTwo(input string) (count int) {
	nums, e := parseInputToInts(input)
	if e != nil {
		panic(e)
	}

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

func parseInputToInts(input string) (nums []int, e error) {
	values := strings.Split(input, "\n")
	nums = make([]int, 0, len(values))

	for _, l := range values {
		if len(l) == 0 {
			continue
		}
		n, e := strconv.Atoi(l)
		if e != nil {
			return nil, e
		}
		nums = append(nums, n)
	}
	return
}

func windowSum(nums []int, i int) int {
	return nums[i-1] + nums[i] + nums[i+1]
}
