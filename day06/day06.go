package day06

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

	var numbers, operations = ParseWorksheet(lines)
	answer = SolveWorksheet(numbers, operations)

	return answer, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	return answer, time.Since(start)
}

func ParseWorksheet(lines []string) (numbers [][]int, operations []byte) {
	for _, line := range lines {
		if len(line) > 0 {
			if line[0] == '*' || line[0] == '+' {
				operations = parseOperations(line)
			} else {
				numbers = append(numbers, parseNumbers(line))
			}
		}
	}

	return
}

func parseNumbers(line string) (numbers []int) {
	var parts = strings.Fields(line)
	for _, val := range parts {
		var num, _ = strconv.Atoi(val)
		numbers = append(numbers, num)
	}

	return
}

func parseOperations(line string) (operations []byte) {
	line = strings.ReplaceAll(line, " ", "")
	operations = []byte(line)

	return
}

func SolveWorksheet(numbers [][]int, operations []byte) (total int) {
	for i, op := range operations {
		if op == '+' {
			total += sum(numbers, i)
		} else {
			total += mult(numbers, i)
		}
	}

	return
}

func sum(numbers [][]int, col int) (sum int) {
	for row := 0; row < len(numbers); row++ {
		sum += numbers[row][col]
	}

	return
}

func mult(numbers [][]int, col int) (product int) {
	product = 1
	for row := 0; row < len(numbers); row++ {
		product *= numbers[row][col]
	}

	return
}
