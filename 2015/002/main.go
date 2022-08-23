// solution for Advent of Code 2015, Day 2: I Was Told There Would Be No Math
// https://adventofcode.com/2015/day/2

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("2015/002/input.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	length := getWrappingPaperLength(string(file))
	fmt.Printf("Required wrapping paper: %d\n", length)
}

func getWrappingPaperLength(file string) int {

	lines := strings.Split(string(file), "\n")
	var area int = 0

	for _, line := range lines {
		num := strings.Split(line, "x")
		area += getArea(num)
	}
	return area
}

func getArea(strings []string) (area int) {
	var ints = make([]int, len(strings))
	var areas []int

	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}

	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			areas = append(areas, ints[i]*ints[j])
		}
	}

	var sum = 0
	smallest := areas[0]
	for i := range areas {
		if areas[i] < smallest {
			smallest = areas[i]
		}
		sum += areas[i]
	}

	area = 2*sum + smallest
	return
}
