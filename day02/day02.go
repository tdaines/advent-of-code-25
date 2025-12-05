package day02

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
	var sum = SumInvalidProductIDs(lines[0])

	return sum, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	return 0, time.Since(start)
}

func ParseProductIDRange(idRange string) (start int, end int) {
	// "11-22"
	var parts = strings.Split(idRange, "-")

	start, _ = strconv.Atoi(parts[0])
	end, _ = strconv.Atoi(parts[1])

	return
}

func FindInvalidProductIDs(rangeStart int, rangeEnd int) []int {
	var invalid = make([]int, 0, 10)

	for i := rangeStart; i <= rangeEnd; i++ {
		var num = strconv.Itoa(i)
		if len(num)&1 == 0 {
			// Split the number into 2 halves
			// "11" -> "1" and "1"
			// "123456" -> "123" and "456"
			var left = num[:len(num)/2]
			var right = num[len(num)/2:]
			if left == right {
				invalid = append(invalid, i)
			}
		}
	}

	return invalid
}

func SumInvalidProductIDs(line string) (sum int) {
	var ranges = strings.Split(line, ",")

	for _, idRange := range ranges {
		start, end := ParseProductIDRange(idRange)
		invalid := FindInvalidProductIDs(start, end)
		sum += sumSlice(invalid)
	}

	return
}

func sumSlice(input []int) (sum int) {
	for _, v := range input {
		sum += v
	}

	return
}
