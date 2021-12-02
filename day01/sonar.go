package day01

import (
	"aoc-2021/util"
)

func countIncreases(input []string) (count int) {
	nums := util.ConvertStrSliceToIntSlice(input)

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
	nums := util.ConvertStrSliceToIntSlice(input)

	lastWindowSum := nums[0] + nums[1] + nums[2]

	for i := range nums {
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

func windowSum(nums []int, i int) int {
	return nums[i-1] + nums[i] + nums[i+1]
}
