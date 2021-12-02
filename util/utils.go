package util

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(day int) (output []string) {
	input, e := os.ReadFile(fmt.Sprintf("../input/day%02d.txt", day))
	check(e)

	for _, line := range strings.Split(string(input), "\n") {
		output = append(output, line)
	}
	
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}