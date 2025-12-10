package day04_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day04"
)

func TestCountPaperNeighbors(t *testing.T) {
	var grid = [][]byte{
		[]byte("..@@.@@@@."),
		[]byte("@@@.@.@.@@"),
		[]byte("@@@@@.@.@@"),
		[]byte("@.@@@@..@."),
		[]byte("@@.@@@@.@@"),
		[]byte(".@@@@@@@.@"),
		[]byte(".@.@.@.@@@"),
		[]byte("@.@@@.@@@@"),
		[]byte(".@@@@@@@@."),
		[]byte("@.@.@@@.@."),
	}

	assert.Equal(t, 2, day04.CountPaperNeighbors(grid, 0, 0))
	assert.Equal(t, 4, day04.CountPaperNeighbors(grid, 0, 1))
	assert.Equal(t, 3, day04.CountPaperNeighbors(grid, 0, 2))
	assert.Equal(t, 3, day04.CountPaperNeighbors(grid, 0, 3))
}

func TestCountAccessiblePaper(t *testing.T) {
	var grid = [][]byte{
		[]byte("..@@.@@@@."),
		[]byte("@@@.@.@.@@"),
		[]byte("@@@@@.@.@@"),
		[]byte("@.@@@@..@."),
		[]byte("@@.@@@@.@@"),
		[]byte(".@@@@@@@.@"),
		[]byte(".@.@.@.@@@"),
		[]byte("@.@@@.@@@@"),
		[]byte(".@@@@@@@@."),
		[]byte("@.@.@@@.@."),
	}

	assert.Equal(t, 13, day04.CountAccessiblePaper(grid))
}

func TestCountAndRemoveAccessiblePaper(t *testing.T) {
	var grid = [][]byte{
		[]byte("..@@.@@@@."),
		[]byte("@@@.@.@.@@"),
		[]byte("@@@@@.@.@@"),
		[]byte("@.@@@@..@."),
		[]byte("@@.@@@@.@@"),
		[]byte(".@@@@@@@.@"),
		[]byte(".@.@.@.@@@"),
		[]byte("@.@@@.@@@@"),
		[]byte(".@@@@@@@@."),
		[]byte("@.@.@@@.@."),
	}

	assert.Equal(t, 43, day04.CountAndRemoveAccessiblePaper(grid))
}
