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


func ConvertStrToInt(in string) (out int) {
	out, e := strconv.Atoi(in)
	Check(e)

	return
}