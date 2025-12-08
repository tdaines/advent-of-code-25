package day03_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day03"
)

func TestFindLargest2DigitJoltage(t *testing.T) {
	assert.Equal(t, 98, day03.FindLargestNDigitJoltage("987654321111111", 2))
	assert.Equal(t, 89, day03.FindLargestNDigitJoltage("811111111111119", 2))
	assert.Equal(t, 78, day03.FindLargestNDigitJoltage("234234234234278", 2))
	assert.Equal(t, 92, day03.FindLargestNDigitJoltage("818181911112111", 2))
}

func TestFindLargest12DigitJoltage(t *testing.T) {
	assert.Equal(t, 987654321111, day03.FindLargestNDigitJoltage("987654321111111", 12))
	assert.Equal(t, 811111111119, day03.FindLargestNDigitJoltage("811111111111119", 12))
	assert.Equal(t, 434234234278, day03.FindLargestNDigitJoltage("234234234234278", 12))
	assert.Equal(t, 888911112111, day03.FindLargestNDigitJoltage("818181911112111", 12))
}
