package day04

import (
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var input []byte

func Part1() (answer int, elapsed time.Duration) {
	var start = time.Now()

	inputCopy := make([]byte, len(input))
	_ = copy(inputCopy, input)

	inputCopy = bytes.TrimRight(inputCopy, "\n")
	var grid = bytes.Split(inputCopy, []byte{'\n'})
	var ans = CountAccessiblePaper(grid)

	return ans, time.Since(start)
}

func Part2() (answer int, elapsed time.Duration) {
	var start = time.Now()

	input = bytes.TrimRight(input, "\n")
	var grid = bytes.Split(input, []byte{'\n'})
	answer = CountAndRemoveAccessiblePaper(grid)

	return answer, time.Since(start)
}

type Grid [][]byte

const (
	PAPER   byte = '@'
	EMPTY   byte = '.'
	REMOVED byte = 'x'
)

func CountAccessiblePaper(grid Grid) int {
	var sum = 0
	for r := range len(grid) {
		for c := range len(grid[r]) {
			if grid[r][c] == PAPER {
				if CountPaperNeighbors(grid, r, c) < 4 {
					grid[r][c] = REMOVED
					sum += 1
				}
			}
		}
	}

	return sum
}

func RemoveAccessiblePaper(grid Grid) {
	for r := range len(grid) {
		for c := range len(grid[r]) {
			if grid[r][c] == REMOVED {
				grid[r][c] = EMPTY
			}
		}
	}
}

func CountAndRemoveAccessiblePaper(grid Grid) int {
	var total = 0

	for {
		var sum = CountAccessiblePaper(grid)
		if sum == 0 {
			break
		}

		total += sum
		RemoveAccessiblePaper(grid)
	}

	return total
}

func CountPaperNeighbors(grid Grid, row, col int) int {
	// Clockwise from straight up
	return IsPaper(grid, row-1, col) +
		IsPaper(grid, row-1, col+1) +
		IsPaper(grid, row, col+1) +
		IsPaper(grid, row+1, col+1) +
		IsPaper(grid, row+1, col) +
		IsPaper(grid, row+1, col-1) +
		IsPaper(grid, row, col-1) +
		IsPaper(grid, row-1, col-1)
}

func IsPaper(grid Grid, row, col int) int {
	if row < 0 || col < 0 {
		return 0
	}

	if row >= len(grid) {
		return 0
	}

	if col >= len(grid[row]) {
		return 0
	}

	if grid[row][col] == PAPER || grid[row][col] == REMOVED {
		return 1
	}

	return 0
}
