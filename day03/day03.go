package day03

import (
	_ "embed"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Part1() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")
	var sum = 0
	for _, line := range lines {
		if len(line) > 0 {
			sum += FindLargestNDigitJoltage(line, 2)
		}
	}

	return sum, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	var lines = strings.Split(input, "\n")
	var sum = 0
	for _, line := range lines {
		if len(line) > 0 {
			sum += FindLargestNDigitJoltage(line, 12)
		}
	}

	return sum, time.Since(start)
}

func findLargestBattery(batteries string) (maxValue, index int) {
	for i := 0; i < len(batteries); i++ {
		var val = int(batteries[i] - '0')

		if val > maxValue {
			maxValue = val
			index = i
		}
	}

	return
}

func FindLargestNDigitJoltage(batteryBank string, numDigits int) int {

	var leftIndex = 0
	var rightIndex = len(batteryBank) - (numDigits - 1)
	var joltage = ""

	for numDigits > 0 {
		var searchSpace = batteryBank[leftIndex:rightIndex]
		var battery, index = findLargestBattery(searchSpace)
		joltage += strconv.Itoa(battery)
		numDigits -= 1
		leftIndex += index + 1
		rightIndex += 1
	}

	var value, _ = strconv.Atoi(joltage)
	return value
}
