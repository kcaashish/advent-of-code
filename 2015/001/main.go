// solution for Advent of Code 2015, Day 1: Not Quite Lisp
// https://adventofcode.com/2015/day/1

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("2015/001/input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	floor := getCurrentFloor(0, string(input))
	basePos := getFloorEntry(0, string(input), -1)
	fmt.Printf("Current floor: %v\n", floor)
	fmt.Printf("First basement entry: %v\n", basePos)
}

func getCurrentFloor(floor int, moves string) int {
	floor = floor + strings.Count(moves, "(") - strings.Count(moves, ")")
	return floor
}

func getFloorEntry(floor int, moves string, target int) int {
	for pos, value := range string(moves) {
		if string(value) == "(" {
			floor += 1
		} else {
			floor -= 1
		}

		if floor == target {
			return pos + 1
		}
	}
	return 0
}
