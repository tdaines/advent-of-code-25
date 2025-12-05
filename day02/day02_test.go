package day02_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdaines/advent-of-code-25/day02"
)

func TestParseProductIDRange(t *testing.T) {
	start, end := day02.ParseProductIDRange("11-22")
	assert.Equal(t, 11, start)
	assert.Equal(t, 22, end)

	start, end = day02.ParseProductIDRange("95-115")
	assert.Equal(t, 95, start)
	assert.Equal(t, 115, end)

	start, end = day02.ParseProductIDRange("1188511880-1188511890")
	assert.Equal(t, 1188511880, start)
	assert.Equal(t, 1188511890, end)
}

func TestFindInvalidProductIDs(t *testing.T) {
	var invalid = day02.FindInvalidProductIDs(11, 22)
	assert.Equal(t, []int{11, 22}, invalid)

	invalid = day02.FindInvalidProductIDs(95, 115)
	assert.Equal(t, []int{99}, invalid)

	invalid = day02.FindInvalidProductIDs(998, 1012)
	assert.Equal(t, []int{1010}, invalid)

	invalid = day02.FindInvalidProductIDs(1188511880, 1188511890)
	assert.Equal(t, []int{1188511885}, invalid)

	invalid = day02.FindInvalidProductIDs(222220, 222224)
	assert.Equal(t, []int{222222}, invalid)

	invalid = day02.FindInvalidProductIDs(1698522, 1698528)
	assert.Equal(t, []int{}, invalid)

	invalid = day02.FindInvalidProductIDs(446443, 446449)
	assert.Equal(t, []int{446446}, invalid)

	invalid = day02.FindInvalidProductIDs(38593856, 38593862)
	assert.Equal(t, []int{38593859}, invalid)
}

func TestSumInvalidProductIDs(t *testing.T) {
	var line = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	var sum = day02.SumInvalidProductIDs(line)
	assert.Equal(t, 1227775554, sum)
}
