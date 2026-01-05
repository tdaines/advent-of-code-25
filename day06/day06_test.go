package day06_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day06"
)

func TestParseWorksheet(t *testing.T) {
	var worksheet = []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}

	var numbers, operations = day06.ParseWorksheet(worksheet)

	assert.Equal(t, 3, len(numbers))
	assert.Equal(t, []int{123, 328, 51, 64}, numbers[0])
	assert.Equal(t, []int{45, 64, 387, 23}, numbers[1])
	assert.Equal(t, []int{6, 98, 215, 314}, numbers[2])

	assert.Equal(t, []byte{'*', '+', '*', '+'}, operations)
}

func TestSolverWorksheet(t *testing.T) {
	var worksheet = []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}

	var numbers, operations = day06.ParseWorksheet(worksheet)
	var total = day06.SolveWorksheet(numbers, operations)

	assert.Equal(t, 4277556, total)
}

func TestParseWorksheetRtL(t *testing.T) {
	var worksheet = []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}

	var numbers, operations = day06.ParseWorksheetRtL(worksheet)

	assert.Equal(t, 4, len(numbers))
	assert.Equal(t, []int{1, 24, 356}, numbers[0])
	assert.Equal(t, []int{369, 248, 8}, numbers[1])
	assert.Equal(t, []int{32, 581, 175}, numbers[2])
	assert.Equal(t, []int{623, 431, 4}, numbers[3])

	assert.Equal(t, []byte{'*', '+', '*', '+'}, operations)
}

func TestSolveWorksheetRtL(t *testing.T) {
	var worksheet = []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}

	var numbers, operations = day06.ParseWorksheetRtL(worksheet)
	var total = day06.SolveWorksheetRtL(numbers, operations)

	assert.Equal(t, 3263827, total)
}
