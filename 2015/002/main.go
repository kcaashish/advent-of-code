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

	ribLength := getRibbonLength(string(file))
	fmt.Printf("Required ribbon length: %d\n", ribLength)
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

func getArea(stringSlice []string) (area int) {
	ints := stringToInt(stringSlice)

	var areas []int
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

func getRibbonLength(file string) (length int) {
	lines := strings.Split(file, "\n")

	for _, line := range lines {
		numStr := strings.Split(line, "x")
		length += getLength(numStr)
	}
	return
}

func getLength(stringSlice []string) (length int) {
	ints := stringToInt(stringSlice)

	cpy := make([]int, len(ints))
	copy(cpy, ints)

	largest := cpy[0]
	indexlargest := 0

	for i := range ints {
		if cpy[i] > largest {
			indexlargest = i
			largest = cpy[i]
		}
	}

	// remove the element in indexlargest
	cpy = append(cpy[:indexlargest], cpy[indexlargest+1:]...)

	forPresent := 2 * (cpy[0] + cpy[1])
	forBow := ints[0] * ints[1] * ints[2]
	length = forPresent + forBow
	return
}

func stringToInt(stringSlice []string) []int {
	var intSlice = make([]int, len(stringSlice))
	for i, s := range stringSlice {
		intSlice[i], _ = strconv.Atoi(s)
	}
	return intSlice
}
