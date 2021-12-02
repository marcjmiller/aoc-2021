package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(day int) (output []string) {
	input, e := os.ReadFile(fmt.Sprintf("../input/day%02d.txt", day))
	Check(e)

	for _, line := range strings.Split(string(input), "\n") {
		output = append(output, line)
	}

	return
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConvertStrSliceToIntSlice(input []string) (nums []int) {
	nums = make([]int, 0, len(input))

	for _, l := range input {
		if len(l) == 0 {
			continue
		}

		n := ConvertStrToInt(l)

		nums = append(nums, n)
	}

	return
}

func ConvertStrToInt(in string) (out int) {
	out, e := strconv.Atoi(in)
	Check(e)

	return
}
