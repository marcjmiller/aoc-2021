package day04

import (
	"fmt"
	"strings"
)

func bingoPartOne(input []string) (result int) {
	ins, boards := parseInput(input)
	
	fmt.Println(ins)

	for i, board := range boards {
		fmt.Println(i, board)
	}
	return 4511
}

func bingoPartTwo(input []string) (result int) {

	return 4512
}

func parseInput(input []string) (ins []string, boards [][]string) {
	ins = strings.Split(input[0], ",")
	boards = append(boards, []string{})
	var board int
	for _, v := range input[2:] {
		boards[board] = append(boards[board], strings.Fields(v)... )
		if v == "" {
			board++
			boards = append(boards, []string{})
			continue
		}
	}

	return
}