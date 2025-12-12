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

	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var numbers, operations = ParseWorksheetRtL(lines)
	answer = SolveWorksheetRtL(numbers, operations)

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

func ParseWorksheetRtL(lines []string) (numbers [][]int, operations []byte) {
	var indices = getColumnIndices(lines[len(lines)-1])

	for i := 0; i < len(indices); i++ {
		var colStart = indices[i]
		var colEnd = len(lines[0])
		if i < len(indices)-1 {
			colEnd = indices[i+1] - 2
		}

		numbers = append(numbers, parseNumbersRtL(lines, colStart, colEnd))
	}

	operations = parseOperations(lines[len(lines)-1])

	return
}

func parseNumbersRtL(lines []string, colStart, colEnd int) (numbers []int) {
	//
	// 123 328  51 64
	//  45 64  387 23
	//   6 98  215 314
	// *   +   *   +

	//
	// 724 2384 57 32 142 25  89  2 45  221 382 9887   24 64 723 66   31
	// 257 7211 49 81 994 872 64  7 841 943 348 8834 1829 29 283 6   143
	// 268 158  96 92 335 931 86 29 957 346   4  614 1872 51 576 1   876
	// 685 51   8  53 2   383 7  55 451 939   1   79 3192 21 654 3  3822
	// *   +    *  +  *   +   *  *  *   *   +   +    +    *  *   *  +
	//

	for i := colStart; i <= colEnd; i++ {
		numbers = append(numbers, parseNumberRtL(lines, i))
	}

	return
}

func parseNumberRtL(lines []string, col int) (num int) {
	for row := 0; row < len(lines)-1; row++ {
		if col < len(lines[row]) {
			var digit = lines[row][col]
			if '0' <= digit && digit <= '9' {
				num = (num * 10) + int(digit-'0')
			}
		}
	}

	return
}

func getColumnIndices(operations string) (indices []int) {
	for i := 0; i < len(operations); i++ {
		if operations[i] == '*' || operations[i] == '+' {
			indices = append(indices, i)
		}
	}

	return
}

func SolveWorksheetRtL(numbers [][]int, operations []byte) (total int) {
	for i, op := range operations {
		if op == '+' {
			total += sumRtL(numbers[i])
		} else {
			total += multRtL(numbers[i])
		}
	}

	return
}

func sumRtL(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}

	return
}

func multRtL(numbers []int) (product int) {
	product = 1
	for _, val := range numbers {
		product *= val
	}

	return
}
